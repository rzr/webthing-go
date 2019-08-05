// -*- Mode: Go; indent-tabs-mode: t -*-
// Copyright: 2019-present Samsung Electronics Co., Ltd. and other contributors
// SPDX-License-Identifier: MPL-2.0

package webthing

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// Value is a Multitypes place holder, only one is used
type Value struct {
	boolean bool
	number  float64
	string  string
}

// Property is made of value(s)
type Property struct {
	name      string
	valuetype string
	value     *Value
}

func (property Property) setValue(value interface{}) {
	switch property.valuetype {
	case "boolean":
		property.value.boolean = (value).(bool)
		break
	case "number":
		property.value.number = (value).(float64)
		break
	case "string":
		property.value.string = (value).(string)
		break
	}
}

func (property Property) getValue() interface{} {
	switch property.valuetype {
	case "boolean":
		return property.value.boolean
		break
	case "number":
		return property.value.number
		break
	case "string":
		return property.value.string
		break
	}
	return nil
}

// NewProperty contruct and assigned value
func NewProperty(name string, valuetype string, value interface{}) *Property {
	var property = Property{name: name, valuetype: valuetype}
	property.value = &Value{}
	property.setValue(value)
	return &property
}

// Thing is made of Property(ies)
type Thing struct {
	id          string
	title       string
	context     string
	types       []string
	description string
	properties  map[string]*Property
	href        string
}

// AddProperty add properties with values
func (thing Thing) AddProperty(property *Property) {
	thing.properties[property.name] = property
}

func (thing Thing) getPropertiesDescriptions() map[string]interface{} {
	result := make(map[string]interface{})

	for name, property := range thing.properties {
		description := make(map[string]interface{})
		description["type"] = property.valuetype
		description["href"] = thing.href + "properties/" + name
		result[name] = description
	}
	return result
}

// NewThing construct without properties
func NewThing(id string, title string, types []string, description string) *Thing {
	thing := Thing{
		id:          id,
		title:       title,
		href:        "/",
		properties:  make(map[string]*Property),
		types:       types,
		description: description,
		context:     "https://iot.mozilla.org/schemas"}
	if thing.types == nil {
		thing.types = make([]string, 0)
	}
	return &thing
}

// Server is an HTTP server made for webthings model
type Server struct {
	thing *Thing
	port  int
}

func (server Server) thingHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	description := make(map[string]interface{})
	description["id"] = server.thing.id
	description["title"] = server.thing.title
	description["@context"] = server.thing.context
	description["@type"] = server.thing.types
	description["properties"] = server.thing.getPropertiesDescriptions()
	description["events"] = make(map[string]interface{})
	description["links"] = [...]map[string]interface{}{
		{"rel": "properties", "href": "/properties"}}
	description["description"] = server.thing.description
	description["actions"] = make(map[string]interface{})

	description["href"] = server.thing.href

	// https://github.com/golang/go/issues/28940
	scheme := "http"
	if req.TLS != nil {
		scheme = "https"
	}
	description["base"] = scheme + "://" + req.Host + req.URL.RequestURI()

	description["security"] = "nosec_sc"
	description["securityDefinitions"] = make(map[string]interface{})
	description["securityDefinitions"].(map[string]interface{})["nosec_sc"] =
		map[string]interface{}{"scheme": "nosec"}

	var body, _ = json.Marshal(description)
	io.WriteString(w, string(body))
}

func (server Server) propertiesHandler(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	description := make(map[string]interface{})
	for name, property := range server.thing.properties {
		description[name] = property.getValue()
	}
	var body, _ = json.Marshal(description)
	io.WriteString(w, string(body))

}

func (server Server) propertyPutHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	object := make(map[string]interface{})
	err := decoder.Decode(&object)
	if err != nil {
		fmt.Println(err)
	}
	name := params.ByName("propertyName")
	property := server.thing.properties[name]
	property.setValue(object[name])
	var body, _ = json.Marshal(object)
	io.WriteString(w, string(body))
}

func (server Server) propertyGetHandler(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	name := params.ByName("propertyName")
	property := server.thing.properties[name]
	object := make(map[string]interface{})
	object[name] = property.getValue()
	var body, _ = json.Marshal(object)
	io.WriteString(w, string(body))
}

// NewServer construct an http server to be started
func NewServer(thing *Thing, port int) *Server {
	return &Server{thing: thing, port: port}
}

// Start listen and process to incoming requests
func (server Server) Start() {
	router := httprouter.New()
	router.GET("/", server.thingHandler)
	router.GET("/properties", server.propertiesHandler)
	router.GET("/properties/:propertyName", server.propertyGetHandler)
	router.PUT("/properties/:propertyName", server.propertyPutHandler)

	address := fmt.Sprintf(":%d", server.port)
	fmt.Print("Listening: ")
	fmt.Println(address)

	err := http.ListenAndServe(address, router)
	fmt.Println(err)
}
