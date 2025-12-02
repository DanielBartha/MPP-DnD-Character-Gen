package service

import "strings"

func SanitizeApiKey(name string) string {
	key := strings.ToLower(name)
	key = strings.ReplaceAll(key, "'", "")
	key = strings.ReplaceAll(key, "’", "")
	key = strings.ReplaceAll(key, "(", "")
	key = strings.ReplaceAll(key, ")", "")
	key = strings.ReplaceAll(key, ",", "")
	key = strings.ReplaceAll(key, "/", "-")
	key = strings.ReplaceAll(key, ":", "")
	key = strings.ReplaceAll(key, ".", "")
	key = strings.ReplaceAll(key, " ", "-")
	key = strings.ReplaceAll(key, " ", "-")
	key = strings.TrimSpace(key)
	return key
}
