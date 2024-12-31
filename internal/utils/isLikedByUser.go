package utils

func IsLikedByUser(likes []string, username string) bool {
	for _, user := range likes {
		if user == username {
			return true
		}
	}
	return false
}
