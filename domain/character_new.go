package domain

import "fmt"

func NewCharacter(name, race, background, class string, level int, stats Stats) (*Character, error) {
	if name == "" {
		return nil, fmt.Errorf("character name is required")
	}
	if race == "" {
		return nil, fmt.Errorf("character race is required")
	}
	if class == "" {
		return nil, fmt.Errorf("character class is required")
	}
	if level <= 0 {
		return nil, fmt.Errorf("character level must be > 0")
	}

	c := &Character{
		Name:       name,
		Race:       race,
		Background: background,
		Class:      class,
		Level:      level,
		Stats:      stats,
		Equipment: Equipment{
			Weapon: map[string]string{"main hand": "", "off hand": ""},
			Armor:  "",
			Shield: "",
		},
	}

	c.ApplyRacialBonuses()

	return c, nil
}
