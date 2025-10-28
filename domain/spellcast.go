package domain

import "strings"

type Spellcasting struct {
	CantripsKnown    int
	SpellsKnown      int
	CanCast          bool
	CasterType       string
	LearnedSpells    []string    `json:"learned_spells"`
	PreparedSpells   []string    `json:"prepared_spells"`
	Slots            map[int]int `json:"slots"`
	MaxSlots         map[int]int `json:"maxslots"`
	PreparedMode     bool
	LearnedMode      bool
	Ability          string
	SpellSaveDC      int
	SpellAttackBonus int
}

type SpellInfo struct {
	Name   string `json:"name"`
	Level  int    `json:"level"`
	School string `json:"school"`
	Range  string `json:"range"`
}

func GetSpellcastingType(class string) string {
	switch strings.ToLower(class) {
	case "bard", "cleric", "druid", "sorcerer", "wizard":
		return "full"
	case "paladin", "ranger":
		return "half"
	case "warlock":
		return "pact"
	default:
		return ""
	}
}
