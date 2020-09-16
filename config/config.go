package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfg *configuration
)

func DB() db {
	return *cfg.DB
}

func Log() log {
	return *cfg.Log
}

func Mode() string {
	return *cfg.Mode
}

func Server() server {
	return *cfg.Server
}

type configuration struct {
	DB     *db     `yaml:"db"`
	Log    *log    `yaml:"log"`
	Mode   *string `yaml:"mode"`
	Server *server `yaml:"server"`
}

type db struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type log struct {
	Path string `yaml:"path"`
}

type server struct {
	Port       string `yaml:"port"`
	AssetsPath string `yaml:"assets_path"`
	Secret     string `yaml:"secret"`
}

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	cfg = &configuration{}
	return decoder.Decode(cfg)
}
