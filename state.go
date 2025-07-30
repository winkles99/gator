package main

import "github.com/winkles99/gator/internal/config"

type State struct {
	ConfigPath string
	Config     *config.Config
}
