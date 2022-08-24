package main

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

func main() {
	// Base directory
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	if err = godotenv.Load(exPath + "/.env"); err != nil {
		panic("error loading .env file: " + exPath + "/.env")
	}

	// Init config
	//conf, err := config.New(exPath, os.Getenv("ENVIRONMENT"))
	//if err != nil {
	//	log.Fatal("Can't init config: ", err)
	//
	//	return
	//}

	//rbacService, err := rbac.NewSimple(conf.Users, conf.Roles)
	//if err != nil {
	//	log.Fatal("Can't init rbac system: ", err)
	//
	//	return
	//}
}
