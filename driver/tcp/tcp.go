package tcp

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/juliotorresmoreno/proxy/config"
	uuid "github.com/satori/go.uuid"
)

var conf config.Config

type Channel struct {
	origen  net.Conn
	destino net.Conn
	uuid    string
}

func Start(config config.Config) error {
	conf = config
	server, err := net.Listen("tcp", conf.Address)
	if err != nil {
		return err
	}
	fmt.Printf("Listening on %v\n", conf.Address)
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		dest, err := net.Dial("tcp", conf.Remote)
		if err != nil {
			log.Println(err)
			if _, err = conn.Write([]byte(err.Error())); err != nil {
				log.Println(err)
			}
			continue
		}
		u1 := uuid.Must(uuid.NewV4())
		go handle(Channel{
			origen:  conn,
			destino: dest,
			uuid:    fmt.Sprintf("%s", u1),
		})
	}
	return nil
}

func handle(channel Channel) {
	if conf.Debug {
		fmt.Print("Conexion establecida\n\n")
	}
	go listen(channel)
	buffer := make([]byte, conf.Tunneling.BufferSize)
	for {
		n, err := channel.origen.Read(buffer)
		if err != nil {
			break
		}
		temp := buffer[:n]
		if conf.Debug {
			fmt.Printf("Request: %v\n%v\n\n", channel.uuid, string(temp))
		}
		if conf.Logging {
			d := time.Now()
			s := fmt.Sprintf("%v-%v-%v", d.Year(), d.Month().String(), d.Day())
			appendFile(path.Join(conf.Logs, s+"_"+channel.uuid+".http"), temp)
		}
		channel.destino.Write(temp)
	}
	channel.destino.Close()
	if conf.Debug {
		fmt.Print("Conexion cerrada\n")
	}
}

func listen(channel Channel) {
	buffer := make([]byte, conf.Tunneling.BufferSize)
	for {
		n, err := channel.destino.Read(buffer)
		if err != nil {
			break
		}
		temp := buffer[:n]
		if conf.Debug {
			fmt.Printf("Response: %v\n%v\n\n", channel.uuid, string(temp))
		}
		if conf.Logging {
			d := time.Now()
			s := fmt.Sprintf("%v-%v-%v", d.Year(), d.Month().String(), d.Day())
			appendFile(path.Join(conf.Logs, s+"_"+channel.uuid+".http"), temp)
		}
		channel.origen.Write(temp)
	}
	channel.origen.Close()
}

func appendFile(path string, text []byte) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(text); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
