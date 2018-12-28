package main

import (
	"fmt"
	"log"
	"plugin"

	"github.com/kechako/plugin-sample/common"
)

func main() {
	p, err := plugin.Open("../plugins/hello/hello.so")
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("NewPlugin")
	if err != nil {
		log.Fatal(err)
	}

	plugin := f.(func() common.Plugin)()

	fmt.Println(plugin.Name())
	err = plugin.Do()
	if err != nil {
		log.Fatal(err)
	}
}
