// Copyright 2020 Roie R. Black

package main

import (
	"fmt"

	"github.com/rblack42/go-burgers/config"
	"github.com/rblack42/go-burgers/gui"
)

func Hello() string {
	return "Burgers Equation Explorer"
}

func main() {
	fmt.Println(Hello())
	cfg := config.GetConfig()
	gui.Run(cfg)
}

