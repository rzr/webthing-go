// -*- Mode: Go; indent-tabs-mode: t -*-
// Copyright: 2019-present Samsung Electronics Co., Ltd. and other contributors
// SPDX-License-Identifier: MPL-2.0

package main

import (
	webthing "github.com/rzr/webthing-go"
)

func main() {
	thing := webthing.NewThing(
		"urn:dev:ops:my-lamp-1234",
		"My Lamp",
		[]string{"OnOffSwitch", "Light"},
		"A web connected lamp")

	onProperty := webthing.NewProperty("on", "boolean", false, nil)
	thing.AddProperty(onProperty)

	brightnessProperty := webthing.NewProperty("brightness", "integer", 0, nil)
	thing.AddProperty(brightnessProperty)

	server := webthing.NewServer(thing, 8888)
	server.Start()
}
