package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBname   string `json:"dbname"`
	} `json:"database"`
}

var Conf Config

func init() {
	file, fileErr := os.Open("conf/config.json")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Conf = Config{}
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Read config file error:", err)
	}
}

func GetConfig() Config {
	return Conf
}
