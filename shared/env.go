package shared

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// This function gets .env file in service root folder and init it.
func LoadEnv() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	newDir := filepath.Join(dir, "..")
	log.Println(newDir)
	envFile := filepath.Join(newDir, ".env")
	err = godotenv.Load(envFile)

	if err != nil {
		panic(err)
	}
}