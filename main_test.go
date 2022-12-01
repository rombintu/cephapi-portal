package main

import (
	"fmt"
	"testing"

	"github.com/rombintu/cephapi-portal/server"
	"github.com/rombintu/cephapi-portal/utils"
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

func TestMapsToCSV(t *testing.T) {
	s := server.NewServer("config.toml")
	sites := s.LoadSitesData()
	for _, s := range sites {
		if s.Title == "DC1" {
			utils.MapsToDoubleList(s.ID, s.Title, s.Clusters)
		}
	}
}
