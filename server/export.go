package server

import (
	"fmt"
	"net/http"

	"github.com/rombintu/cephapi-portal/utils"
)

func (s *Server) ExportToCsv(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	title := req.Form.Get("title")
	if title == "default" {
		http.Redirect(res, req, "/", 200)
		return
	}
	site := s.LoadSiteDataByTitle(title)

	csvfile := utils.MapsToCSV(site.ID, site.Title, site.Clusters)
	// csvfileName := fmt.Sprintf("%s.csv", title)
	// csvfile, err := ioutil.ReadFile(csvfileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	res.Header().Set("Accept-ranges", "bytes")
	res.Header().Set("Content-Type", "application/octet-stream")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.csv\"", site.Title))
	res.WriteHeader(http.StatusOK)
	res.Write(csvfile.Bytes())
	// http.Redirect(res, req, "/", 200)
}
