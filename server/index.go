package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Clusters map[string]interface{}

type Site struct {
	ID       int
	Title    string
	Clusters Clusters
}

type IndexPage struct {
	Sites []Site
}

func (s *Server) Index(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
	)
	if err != nil {
		fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
	}
	DataIndexPage := IndexPage{
		Sites: s.LoadSitesData(),
	}

	template.ExecuteTemplate(res, "index", DataIndexPage)
}

func ParseDataFromUrl(url string) Clusters {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return Clusters{}
	}
	defer resp.Body.Close()
	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return Clusters{}
	}
	var clusters Clusters
	if err := json.Unmarshal(jsonData, &clusters); err != nil {
		log.Println(err)
		return Clusters{}
	}
	return clusters
}

func (s *Server) LoadSitesData() []Site {
	var sites []Site
	var i int
	for title, url := range s.Config.Urls {
		i++
		sites = append(
			sites,
			Site{
				ID:       i,
				Title:    title,
				Clusters: ParseDataFromUrl(url),
			},
		)
	}
	return sites
}

func (s *Server) LoadSiteDataByTitle(title string) Site {
	return Site{
		Title:    title,
		Clusters: ParseDataFromUrl(s.Config.Urls[title]),
	}
}
