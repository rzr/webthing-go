// -*- Mode: Go; indent-tabs-mode: t -*-
// Copyright: 2019-present Samsung Electronics Co., Ltd. and other contributors
// SPDX-License-Identifier: MPL-2.0

package main

import (
	webthing "github.com/rzr/webthing-go"
	"fmt"
)

func valueHandler(update interface{}) {
	fmt.Println("change:", (update).(bool))
}

func main() {
	thing := webthing.NewThing(
		"urn:dev:ops:my-actuator-1234",
		"ActuatorExample",
		[]string{"OnOffSwitch"},
		"An actuator example")

	onProperty := webthing.NewProperty("on", "boolean", false, valueHandler)
	thing.AddProperty(onProperty)

	server := webthing.NewServer(thing, 8888)
	server.Start()
}
