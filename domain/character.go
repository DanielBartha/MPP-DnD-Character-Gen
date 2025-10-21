package domain

import (
	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"
)

var SpellcastingType = map[string]string{
	"bard":     "full",
	"cleric":   "full",
	"druid":    "full",
	"sorcerer": "full",
	"wizard":   "full",

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
	CantripsKnown  int
	SpellsKnown    int
	CanCast        bool
	CasterType     string
	LearnedSpells  []string    `json:"learned_spells"`
	PreparedSpells []string    `json:"prepared_spells"`
	Slots          map[int]int `json:"slots"`
	MaxSlots       map[int]int `json:"maxslots"`
	PreparedMode   bool
	LearnedMode    bool
}

type SpellInfo struct {
	Name   string `json:"name"`
	Level  int    `json:"level"`
	School string `json:"school"`
	Range  string `json:"range"`
}

type WeaponInfo struct {
	Name      string
	Category  string
	Range     int
	TwoHanded bool
}

type ArmorInfo struct {
	Name     string
	BaseAC   int
	DexBonus bool
	MaxBonus int
}
