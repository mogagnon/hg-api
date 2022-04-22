package hgApi

import (
	"fmt"
	"os"
	"testing"
)

const HgToken = "HG_TOKEN"

func TestMain(m *testing.M) {
	if token := os.Getenv(HgToken); token == "" {
		fmt.Println("\033[31m", "HG_TOKEN must be set for acceptance tests", "\033[0m")
	}

	code := m.Run()
	os.Exit(code)
}
