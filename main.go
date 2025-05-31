package main

import (
	"calcs/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DefaultDBHost, cfg.DefaultDBUser, cfg.DefaultDBPassword, cfg.DefaultDBName, cfg.DefaultDBPort,
	)
	defaultDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
