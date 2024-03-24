package config

import (
	"flag"
	"os"
)

func Flags() (string, string) {
	// Определение флагов
	address := flag.String("a", "localhost:8080", "адрес запуска HTTP-сервера")
	baseURL := flag.String("b", "http://localhost:8080", "базовый адрес результирующего сокращённого URL") // порты должны совпадать иначе кабзда
	// Парсинг флагов
	flag.Parse()
	if envAddress := os.Getenv("SERVER_ADDRESS"); envAddress != "" {
		*address = envAddress
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		*baseURL = envBaseURL
	}
	return *address, *baseURL
}
