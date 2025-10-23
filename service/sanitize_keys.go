package service

import "strings"

func SanitizeKey(name string) string {
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
