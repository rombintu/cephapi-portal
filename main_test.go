package main

import (
	"fmt"
	"testing"

	"github.com/rombintu/cephapi-portal/server"
)

func TestCreateServer(t *testing.T) {
	s := server.NewServer("config.toml")
	fmt.Println(s.Config)
}

func TestParseData(t *testing.T) {
	s := server.NewServer("config.toml")
	sites := s.LoadSitesData()
	fmt.Println(sites)
}
