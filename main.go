package main

import (
	"log"
	"os"

	"github.com/juliotorresmoreno/proxy/config"
	"github.com/juliotorresmoreno/proxy/driver/http"
	"github.com/juliotorresmoreno/proxy/driver/tunneling"
)

func main() {
	conf := config.GetConfig()
	os.MkdirAll(conf.Logs, 644)
	switch conf.Mode {
	case "tunneling":
		log.Fatal(tunneling.Start(conf))
	case "http":
		log.Fatal(http.Start(conf))
	}
}
