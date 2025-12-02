package service

import (
	"fmt"
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func FormatCharacterView(c *domain.Character) string {
	var sb strings.Builder

	fmt.Fprintf(&sb,
		"Name: %s\n"+
			"Class: %s\n"+
			"Race: %s\n"+
			"Background: %s\n"+
			"Level: %d\n"+
			"Ability scores:\n"+
			"  STR: %d (%+d)\n"+
			"  DEX: %d (%+d)\n"+
			"  CON: %d (%+d)\n"+
			"  INT: %d (%+d)\n"+
			"  WIS: %d (%+d)\n"+
			"  CHA: %d (%+d)\n"+
			"Proficiency bonus: +%d\n"+
			"Skill proficiencies: %s\n",
		c.Name,
		strings.ToLower(c.Class),
		strings.ToLower(c.Race),
		c.Background,
		c.Level,
		c.Stats.Str, c.Stats.StrMod,
		c.Stats.Dex, c.Stats.DexMod,
		c.Stats.Con, c.Stats.ConMod,
		c.Stats.Intel, c.Stats.IntelMod,
		c.Stats.Wis, c.Stats.WisMod,
		c.Stats.Cha, c.Stats.ChaMod,
		c.Proficiency,
		strings.Join(c.Skills.Skills, ", "),
	)

	if c.Spellcasting != nil && c.Spellcasting.CanCast {
		fmt.Fprintln(&sb, "Spell slots:")

		if c.Spellcasting.CantripsKnown > 0 {
			fmt.Fprintf(&sb, "  Level 0: %d\n", c.Spellcasting.CantripsKnown)
		}

		for lvl := 1; lvl <= 9; lvl++ {
			if count, ok := c.Spellcasting.MaxSlots[lvl]; ok && count > 0 {
				fmt.Fprintf(&sb, "  Level %d: %d\n", lvl, count)
			}
		}

		if c.Spellcasting.Ability != "" {
			fmt.Fprintf(&sb, "Spellcasting ability: %s\n", strings.ToLower(c.Spellcasting.Ability))
			fmt.Fprintf(&sb, "Spell save DC: %d\n", c.Spellcasting.SpellSaveDC)
			fmt.Fprintf(&sb, "Spell attack bonus: +%d\n", c.Spellcasting.SpellAttackBonus)
		}
	}

	if weapon, ok := c.Equipment.Weapon["main hand"]; ok && weapon != "" {
		fmt.Fprintf(&sb, "Main hand: %s\n", weapon)
	}

	if weapon, ok := c.Equipment.Weapon["off hand"]; ok && weapon != "" {
		fmt.Fprintf(&sb, "Off hand: %s\n", weapon)
	}
	if c.Equipment.Armor != "" {
		fmt.Fprintf(&sb, "Armor: %s\n", c.Equipment.Armor)
	}

	if c.Equipment.Shield != "" {
		fmt.Fprintf(&sb, "Shield: %s\n", c.Equipment.Shield)
	}

	fmt.Fprintf(&sb,
		"Armor class: %d\n"+
			"Initiative bonus: %d\n"+
			"Passive perception: %d\n",
		c.ArmorClass,
		c.InitiativeBonus,
		c.PassivePerception,
	)

	return sb.String()
}
