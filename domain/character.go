package domain

import "github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"

type Character struct {
	Name        string
	Race        string
	Background  string
	Class       string
	Level       int
	Proficiency int
	Stats       Stats
	Skills      characterClasses.ClassSkills
	// Equipment Equipment
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
