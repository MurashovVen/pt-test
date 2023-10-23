package main

import (
	"pt-test/pkg/app/configuration"
)

type config struct {
	configuration.Default

	HTTPServerAddress string `desc:"host:port" default:":8080" split_words:"true"`

	DB string `desc:"database connection uri" required:"true" split_words:"true"`
}
