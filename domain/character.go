package domain

import "strings"

type Character struct {
	Name              string
	Race              string
	Background        string
	Class             string
	Level             int
	Proficiency       int
	Stats             Stats
	Skills            ClassLoadout
	Equipment         Equipment
	Spellcasting      *Spellcasting
	ArmorClass        int
	InitiativeBonus   int
	PassivePerception int
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

func (c *Character) ApplyRacialBonuses() {
	race := strings.ToLower(strings.ReplaceAll(c.Race, " ", "-"))
	switch race {
	case "dwarf":
		c.Stats.Con += 2

	case "hill-dwarf":
		c.Stats.Con += 2
		c.Stats.Wis += 1

	case "elf":
		c.Stats.Dex += 2

	case "halfling", "stout-halfling":
		c.Stats.Dex += 2

	case "lightfoot-halfling":
		c.Stats.Dex += 2
		c.Stats.Cha++

	case "human":
		c.Stats.Str++
		c.Stats.Dex++
		c.Stats.Con++
		c.Stats.Intel++
		c.Stats.Wis++
		c.Stats.Cha++

	case "dragonborn":
		c.Stats.Str += 2
		c.Stats.Cha++

	case "gnome":
		c.Stats.Intel += 2

	// TODO: half-eelves get to choose which stats to increase besides the rizz, for now dex and wis as defaults
	case "half-elf":
		c.Stats.Cha += 2
		c.Stats.Dex++
		c.Stats.Wis++

	case "half-orc":
		c.Stats.Str += 2
		c.Stats.Con++

	case "tiefling":
		c.Stats.Cha += 2
		c.Stats.Intel++
	}

	c.UpdateProficiency()
}
