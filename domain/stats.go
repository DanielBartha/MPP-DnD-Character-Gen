package domain

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

func abilityModifier(score int) int {
	result := (score - 10) / 2
	if (score-10)%2 < 0 {
		result--
	}
	return result
}
