package http

import (
	"github.com/juliotorresmoreno/proxy/config"
)

var conf config.Config

func Start(config config.Config) error {
	conf = config

	return nil
}
