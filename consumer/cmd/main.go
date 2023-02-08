package main

import (
	"dheeraj-326/plugin-consumer/domain"
	"fmt"
	"plugin"
)

// type aliasing NOT definition
type pluginfactory = func(pluginName string) interface{}

func main() {
	fmt.Println("Consumer")
	plugin, err := plugin.Open("../plugins/bin/plugins.so")
	if err != nil {
		panic(err)
	}
	symbol, err := plugin.Lookup("PluginFactory")
	if err != nil {
		panic(err)
	}
	pluginfactory := symbol.(pluginfactory)
	adder := pluginfactory("Adder").(domain.Adder)
	fmt.Println(adder.Add(1, 2))
}
