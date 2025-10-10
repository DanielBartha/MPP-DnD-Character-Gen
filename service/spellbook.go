package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SpellInfo struct {
	Level   int
	Classes []string
}

var SpellDB map[string]SpellInfo

func LoadSpellsCSV(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Read()

	SpellDB = make(map[string]SpellInfo)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		name := strings.TrimSpace(record[0])
		levelStr := record[1]
		classStr := record[2]

		level, _ := strconv.Atoi(levelStr)
		classes := splitAndTrim(classStr)

		SpellDB[strings.ToLower(name)] = SpellInfo{
			Level:   level,
			Classes: classes,
		}
	}

	return nil
}

func splitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))

	for _, p := range parts {
		result = append(result, strings.ToLower(strings.TrimSpace(p)))
	}

	return result
}

func GetSpellLevel(name string) (int, error) {
	spell, ok := SpellDB[strings.ToLower(name)]

	if !ok {
		return 0, fmt.Errorf("unknown spell: %s", name)
	}

	return spell.Level, nil
}

func IsSpellForClass(spellName, class string) bool {
	spell, ok := SpellDB[strings.ToLower(spellName)]
	if !ok {
		return false
	}

	class = strings.ToLower(class)
	for _, c := range spell.Classes {
		if c == class {
			return true
		}
	}

	return false
}
