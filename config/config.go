package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	DB  db
	Mod string
	Log log

	Service service
	Server  server
)

type all struct {
	DB      db      `yaml:"db"`
	Log     log     `yaml:"log"`
	Mod     string  `yaml:"mod"`
	Service service `yaml:"service"`
	Server  server  `yaml:"server"`
}

type db struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Base string `yaml:"base"`
}

type log struct {
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

type service struct {
}

type server struct {
	Port string `yaml:"port"`
}

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	a := &all{}
	err = decoder.Decode(a)
	if err != nil {
		return err
	}

	DB = a.DB
	Log = a.Log
	Mod = a.Mod

	Service = a.Service
	Server = a.Server

	return nil
}
