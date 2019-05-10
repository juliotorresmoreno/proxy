package main

import (
	"log"
	"os"

	"github.com/juliotorresmoreno/proxy/config"
	"github.com/juliotorresmoreno/proxy/driver/http"
	"github.com/juliotorresmoreno/proxy/driver/tcp"
)

func main() {
	conf := config.GetConfig()
	os.MkdirAll(conf.Logs, 644)
	switch conf.Mode {
	case "tunneling":
		log.Fatal(tcp.Start(conf))
	case "http":
		log.Fatal(http.Start(conf))
	}
}
