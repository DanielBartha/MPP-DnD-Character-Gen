package domain

import (
	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"
)

var SpellcastingType = map[string]string{
	"wizard":   "full",
	"sorcerer": "full",
	"cleric":   "full",
	"druid":    "full",
	"bard":     "full",

	"paladin": "half",
	"ranger":  "half",

	"warlock": "pact",
}

type Character struct {
	Name         string
	Race         string
	Background   string
	Class        string
	Level        int
	Proficiency  int
	Stats        Stats
	Skills       characterClasses.ClassSkills
	Equipment    Equipment
	Spellcasting *Spellcasting
}

type Stats struct {
	Str    int
	StrMod int

	Dex    int
	DexMod int

	Con    int
	ConMod int

	Intel    int
	IntelMod int

	Wis    int
	WisMod int

	Cha    int
	ChaMod int
}

type Equipment struct {
	Weapon map[string]string
	Armor  string
	Shield string
}

type Spellcasting struct {
	LearnedSpells  []string    `json:"learned_spells"`
	PreparedSpells []string    `json:"prepared_spells"`
	Slots          map[int]int `json:"slots"`
	MaxSlots       map[int]int `json:"maxslots"`
}
