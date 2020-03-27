package main

import (
	"os"
	"testing"
)

func TestGetServerAddrDefault(t *testing.T) {
	os.Setenv("HOST", "")
	os.Setenv("PORT", "")
	addr := getServerAddr()

	if addr != "127.0.0.1:8080" {
		t.Fail()
	}
}

func TestGetServerAddrEnvVar(t *testing.T) {
	os.Setenv("HOST", "0.0.0.0")
	os.Setenv("PORT", "80")
	addr := getServerAddr()

	if addr != "0.0.0.0:80" {
		t.Fail()
	}
}
