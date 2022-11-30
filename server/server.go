package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/rombintu/cephapi-portal/utils"
)

type Config struct {
	Urls map[string]string `toml:"Urls"`
	Port string            `toml:"Port"`
}

type Server struct {
	Router *mux.Router
	Config Config
}

// Default parsing config
func NewConfig(confPath string) Config {
	confFile, err := os.Open(confPath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer confFile.Close()

	content, err := ioutil.ReadAll(confFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var conf Config
	if _, err := toml.Decode(string(content), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	// Try get PORT from env
	if conf.Port == "" {
		fmt.Printf("Config.Port: not found.\nTry parse from env or set default 8081\n")
		conf.Port = utils.GetEnvOfDefault("PORT", "8081")
	}
	return conf
}

func NewServer(confPath string) *Server {
	return &Server{
		Router: NewRouter(),
		Config: NewConfig(confPath),
	}
}

func (s *Server) Listen() {
	fmt.Printf("Go to: http://localhost:%s\n", s.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", s.Config.Port), s.Router)
}
