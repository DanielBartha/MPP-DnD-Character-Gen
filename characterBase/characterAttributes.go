package characterBase

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"
)

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

	c.Stats.StrMod = abilityModifier(c.Stats.Str)
	c.Stats.DexMod = abilityModifier(c.Stats.Dex)
	c.Stats.ConMod = abilityModifier(c.Stats.Con)
	c.Stats.IntelMod = abilityModifier(c.Stats.Intel)
	c.Stats.WisMod = abilityModifier(c.Stats.Wis)
	c.Stats.ChaMod = abilityModifier(c.Stats.Cha)
}

func (c *Character) AssignClassSkills() {
	// match uppercase and lowercase inputs
	classKey := strings.ToLower(strings.TrimSpace(c.Class))

	// "ok" is some go voodoo that checks if value exists in the map
	classSkills, ok := characterClasses.Classes[classKey]
	if !ok {
		// fallback just in case no value is found
		c.Skills = characterClasses.ClassSkills{
			MaxAllowed: 0,
			Skills:     []string{},
		}
		c.Skills.Skills = append(c.Skills.Skills, "Insight", "Religion")
		return
	}

	// copying value to not change the global map accidentally
	src := classSkills.Skills
	localSlice := make([]string, len(src))
	copy(localSlice, src)

	local := characterClasses.ClassSkills{
		MaxAllowed: classSkills.MaxAllowed,
		Skills:     localSlice,
	}
	local.Skills = append(local.Skills, "insight", "religion")

	c.Skills = local
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
