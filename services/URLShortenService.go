package services

var urlMap = make(map[string]string)

func RegisterURL(key string, value string){
	urlMap[key] = value
}

func UnregisterURL(key string){
	delete(urlMap, key)
}

func GetURL(key string) string {
	return urlMap[key]
}

func IsURLRegistered(key string) bool {
	if _, ok := urlMap[key]; ok {
		return true
	}
	return false
}
