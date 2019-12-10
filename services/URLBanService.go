package services

var banList = make(map[string]int)

func RegisterBanURL(key string){
	banList[key] = -1
}

func UnregisterBanURL(key string){
	delete(banList, key)
}

func IsBanned(key string) bool {
	return banList[key] == -1
}
