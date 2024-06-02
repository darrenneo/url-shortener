package storage

var urlMap map[string]string

func InitStorage() {
	urlMap = make(map[string]string)
}

func AddURL(orginal, new string) {
	urlMap[new] = orginal
}
