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

func SanitizeLocalKey(name string) string {
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
	key = strings.TrimSuffix(key, "-armor")
	key = strings.TrimSuffix(key, " armor")
	key = strings.TrimSpace(key)
	return key
}
