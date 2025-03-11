package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/SyntinelNyx/syntinel-agent/internal/kopia"
	"github.com/SyntinelNyx/syntinel-agent/internal/sysinfo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	sysinfo.SysInfo()
	sysinfo.ConnectToServer(nil)
	kopia.OpenRepository()
	
}
