package main

import (
	"os"
	"testing"
)

func TestStartApp(t *testing.T) {
	initEnv()
	go main()
}

func initEnv() {
	_ = os.Setenv("DB_HOST", "host.docker.internal")
	_ = os.Setenv("DB_USER", "spuser")
	_ = os.Setenv("DB_PASSWORD", "SPuser96")
	_ = os.Setenv("DB_NAME", "challenge_appdb")
	_ = os.Setenv("DB_PORT", "1234")
	_ = os.Setenv("DB_SCHEMA", "challenge")
	_ = os.Setenv("APP_KEY", "123")
	_ = os.Setenv("PORT", "8080")
}
