package storage

import "fmt"

var urlMap map[string]string

func InitStorage() {
	urlMap = make(map[string]string)
}

func AddURL(orginal, new string) error {
	_, exist := GetURL(new)

	if exist {
		return fmt.Errorf("new URL name exist")
	}

	urlMap[new] = orginal

	return nil
}

func ListAll() {
	for key, value := range urlMap {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

func GetURL(key string) (string, bool) {
	value, ok := urlMap[key]
	return value, ok
}
