package main

import (
	"fmt"

	"github.com/kechako/plugin-sample/common"
)

type Plugin struct{}

func NewPlugin() common.Plugin {
	return &Plugin{}
}

func (p *Plugin) Name() string {
	return "Plugin!!!"
}
func (p *Plugin) Do() error {
	fmt.Println("Do Plugin!!!")
	return nil
}
