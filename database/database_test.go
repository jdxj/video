package database

import (
	"fmt"
	"testing"

	"github.com/jdxj/video/config"
)

func init() {
	config.Init("../config/config.yaml")
	Init()
}

func TestLoginCheck(t *testing.T) {
	u, err := LoginCheck("jdxj", "")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%#v\n", u)
}

func TestUser_Roles(t *testing.T) {
	defer db.Close()

	u := &User{
		ID: 1,
	}
	roles, err := u.Roles()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, r := range roles {
		fmt.Printf("%s\n", r.Name)
	}
}
