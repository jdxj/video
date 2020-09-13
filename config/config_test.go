package config

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init("config.yaml")

	fmt.Printf("%#v\n", DB)
	fmt.Printf("%s\n", Mod)
}
