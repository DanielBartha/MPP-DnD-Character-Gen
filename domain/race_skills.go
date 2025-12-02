package domain

import "strings"

var racialSkillProficiencies = map[string][]string{
	"dwarf":              {"history"},
	"hill-dwarf":         {"history"},
	"elf":                {"perception"},
	"halfling":           {},
	"lightfoot-halfling": {},
	"half-elf":           {"insight", "perception"},
	"half-orc":           {"intimidation"},
	"gnome":              {},
	"rock gnome":         {"history"},
	"human":              {},
	"dragonborn":         {},
	"tiefling":           {},
}

func (c *Character) ApplyRacialSkillProficiencies() {
	race := strings.ToLower(strings.ReplaceAll(c.Race, " ", "-"))

	skills, ok := racialSkillProficiencies[race]
	if !ok {
		return
	}

	c.Skills.Skills = append(c.Skills.Skills, skills...)
}
