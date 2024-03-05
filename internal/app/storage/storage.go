package storage

import "errors"

var m map[string]string

func AddToMap(key string, value string) {
	m[key] = value
}

func RemoveFromMap(key string) {
	delete(m, key)
}

func GetValueByKey(key string) (string, error) {
	value, found := m[key]
	if !found {
		return "", errors.New("Key not found")
	}
	return value, nil
}
