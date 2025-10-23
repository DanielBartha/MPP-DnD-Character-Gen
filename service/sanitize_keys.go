package service

import "strings"

func SanitizeKey(name string) string {
	index := strings.ToLower(name)
	index = strings.ReplaceAll(index, "'", "")
	index = strings.ReplaceAll(index, "’", "")
	index = strings.ReplaceAll(index, "(", "")
	index = strings.ReplaceAll(index, ")", "")
	index = strings.ReplaceAll(index, ",", "")
	index = strings.ReplaceAll(index, "/", "-")
	index = strings.ReplaceAll(index, ":", "")
	index = strings.ReplaceAll(index, ".", "")
	index = strings.ReplaceAll(index, " ", "-")
	index = strings.TrimSpace(index)
	return index
}
