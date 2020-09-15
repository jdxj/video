package logger

import (
	"fmt"
	"testing"

	"github.com/jdxj/video/config"
)

func TestInit(t *testing.T) {
	config.Init("../config/config.yaml")
	fmt.Printf("%#v\n", config.Log)
	Init()

	Debug("test: %s", "abc")
	Info("test2: %s", "def")
	Sync()
}
