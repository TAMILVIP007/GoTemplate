package src

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func loadenv() *Envs {
	if os.Getenv("ENV") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	}
	log.Println("Loaded environment variables from .env file")
	env := Envs{
		Token:   os.Getenv("TOKEN"),
		DbUrl:   os.Getenv("DATABASE_URL"),
		Devs:    strings.Split(os.Getenv("DEVS"), " "),
		Webhook: os.Getenv("WEBHOOK"),
		Port:    os.Getenv("PORT"),
	}

	return &env
}

var Envars = loadenv()

func Converttoin32(s string) int32 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int32(i)
}

func Converttoin64(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int64(i)
}

func CheckIsDev(id int64) bool {
	for _, i := range Envars.Devs {
		if strconv.FormatInt(id, 10) == i {
			return true
		}
	}
	return false
}
