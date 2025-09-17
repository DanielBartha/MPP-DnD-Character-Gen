package characterBase

import characterClasses "github.com/DanielBartha/MPP-DnD-Character-Gen/characteerClasses"

func abilityModifier(score int) int {
	return (score - 10) / 2
}

func (c *Character) UpdateProficiency() {
	switch {
	case c.Level >= 17:
		c.Proficiency = 6
	case c.Level >= 13:
		c.Proficiency = 5
	case c.Level >= 9:
		c.Proficiency = 4
	case c.Level >= 5:
		c.Proficiency = 3
	default:
		c.Proficiency = 2
	}

	// ability modifiers
	c.Stats.StrMod = abilityModifier(c.Stats.Str)
	c.Stats.DexMod = abilityModifier(c.Stats.Dex)
	c.Stats.ConMod = abilityModifier(c.Stats.Con)
	c.Stats.IntelMod = abilityModifier(c.Stats.Intel)
	c.Stats.WisMod = abilityModifier(c.Stats.Wis)
	c.Stats.ChaMod = abilityModifier(c.Stats.Cha)
}

func (c *Character) checkClass() {

}

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

// type Equipment struct {
// 	Armaments string
// 	Armor     string
// 	Gear      string
// 	Tools     string
// 	Mounts    string
// }
