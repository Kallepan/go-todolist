/*
	Package to load in env file.
*/

package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbInfo struct {
	User     string
	Password string
	DbName   string
	Host     string
	Port     int
}

func GetDbInfo() DbInfo {
	envSrcUser := "POSTGRES_USER"
	envSrcPass := "POSTGRES_PASSWORD"
	envSrcDbName := "POSTGRES_DB"
	envSrcPort := "POSTGRES_PORT"
	envSrcHost := "POSTGRES_HOST"

	port, err := strconv.Atoi(GetValueFromEnv(envSrcPort))
	if err != nil {
		log.Fatal("Failed to convert port to int")
	}

	dbInfo := DbInfo{
		User:     GetValueFromEnv(envSrcUser),
		Password: GetValueFromEnv(envSrcPass),
		DbName:   GetValueFromEnv(envSrcDbName),
		Host:     GetValueFromEnv(envSrcHost),
		Port:     port,
	}

	return dbInfo
}

func GetValueFromEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		fmt.Println(key + " is not present in the .env file... Using empty string")
	}

	return value
}

func LoadEnv() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Failed to load .env file.")
	}

	fmt.Println("Env loaded")
}
