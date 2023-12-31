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
		if s.Title == "API1" {
			bts := utils.MapsToCSV(s.ID, s.Title, s.Clusters)
			fmt.Println(bts)
		}
	}
}

func TestStrToFloatStr(t *testing.T) {
	strs := []string{
		"-864.0 GB",
		"0B",
		"1.1 PB",
		"+1301.0 GB",
		"-100 KB",
	}
	for _, str := range strs {
		fmt.Println(utils.StrToFloatStr(str))
	}
}
