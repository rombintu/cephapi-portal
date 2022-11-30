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

// type SDS map[string]interface{}

// 	Zones map[string]string `json:"zones,omitempty"`
// 	Pools map[string]string `json:"pools,omitempty"`
// }

// type Oceanstores map[string]interface{}

// type Clusters struct {
// 	ID          int
// 	SDS1        SDS         `json:"sds1,omitempty"`
// 	SDS2        SDS         `json:"sds2,omitempty"`
// 	SDS3        SDS         `json:"sds3,omitempty"`
// 	SDS4        SDS         `json:"sds4,omitempty"`
// 	SDS5        SDS         `json:"sds5,omitempty"`
// 	SDS6        SDS         `json:"sds6,omitempty"`
// 	Oceanstores Oceanstores `json:"oceanstores,omitempty"`
// 	Timestamp   string      `json:"timestamp,omitempty"`
// }

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
		// "templates/footer.gohtml",
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
