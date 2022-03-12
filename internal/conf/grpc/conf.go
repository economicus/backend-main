package grpc

import (
	"context"
	"fmt"
	"log"
	"os"
)

type conf struct {
	host string
	port string
	ctx  context.Context
}

func newConf() *conf {
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.SetPrefix("[WARNING] ")

	host := os.Getenv("GRPC_HOST")
	port := os.Getenv("GRPC_INSECURE_PORT")

	if host == "" {
		log.Println("MISSING GRPC ENV: empty host")
		host = "172.17.0.1"
	}
	if port == "" {
		log.Println("MISSING GRPC ENV: empty port")
		port = "9000"
	}
	return &conf{
		host: host,
		port: port,
		ctx:  context.Background(),
	}
}

func (c *conf) getDSN() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}
