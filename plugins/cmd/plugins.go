package main

import (
	"dheeraj-326/go-plugin-sample/modules"
	"errors"
)

func PluginFactory(pluginName string) interface{} {
	switch pluginName {
	case "Adder":
		return modules.NewAdder()
	case "Divider":
		return modules.NewDivider()
	default:
		panic(errors.New("invalid plugin name"))
	}
}
