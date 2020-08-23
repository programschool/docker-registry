package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type conf struct {
	Host     string
	Port     string
	CertFile string
	KeyFile  string
	Proxy    []struct {
		Match   string
		Request string
		Scheme  string
		Host    string
		Port    string
	}
	Api   string
	DeBug bool
}

func Load() conf {
	var config conf
	f, err := os.Open(filepath.ToSlash("./config.json"))
	if err != nil {
		log.Println(err.Error())
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Println(err.Error())
	}

	return config
}
