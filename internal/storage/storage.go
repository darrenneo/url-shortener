package storage

import "fmt"

var urlMap map[string]string

func InitStorage() {
	urlMap = make(map[string]string)
}

func AddURL(orginal, new string) error {
	exist := checkExisting(new)

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

func checkExisting(key string) bool {
	_, ok := urlMap[key]
	return ok
}
