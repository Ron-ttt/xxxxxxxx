package config

import "flag"

func Flags() (string, string) {
	// Определение флагов
	address := flag.String("a", "localhost:8080", "адрес запуска HTTP-сервера")
	baseURL := flag.String("b", " ", "базовый адрес результирующего сокращённого URL") // порты должны совпадать иначе кабзда
	// Парсинг флагов
	flag.Parse()
	return *address, *baseURL
}
