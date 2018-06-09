package config

import (
	"testing"
	"os"
)


type Config struct {
	DB string
	PORT int64
	Message string
}

func TestSetConfig(t *testing.T) {
	c := Config{}
	os.Setenv("TEST_MESSAGE", "HOLA")
	SetConfig("TEST", &c, "config_test")
	AddConfigPath("/etc/config")
	if c.DB != "test" {
		t.Fatal("DB should be test")
	}
	if c.PORT != 1234 {
		t.Fatal("POST should be 1234")
	}
	if c.Message != "HOLA" {
		t.Fatal("MESSAGE should be HOLA")
	}

}

func TestAddConfigPath(t *testing.T) {
	configPath := []string{}
	path := "/etc/config"
	configPath = append(configPath,path)
	AddConfigPath(path)
	if GetConfigPath()[0] != configPath[0] {
		t.Fatal("config maps should be equals")
	}
}