package storage

import "errors"

var m = make(map[string]string)

func AddToMap(key string, value string) {
	m[key] = value
}

func RemoveFromMap(key string) {
	delete(m, key)
}

func GetValueByKey(key string) (string, error) {
	value, found := m[key]
	if !found {
		return "", errors.New("key not found")
	}
	return value, nil
}
