package main

import (
	"flag"

	"github.com/rombintu/cephapi-portal/server"
)

func main() {
	confPath := flag.String("conf", "config.toml", "Config toml-file for parsing")
	flag.Parse()
	s := server.NewServer(*confPath)
	s.ConfigureRouter()
	s.Listen()
}
