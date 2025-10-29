package domain

import (
	"errors"
	"strings"
)

type Equipment struct {
	Weapon map[string]string
	Armor  string
	Shield string
}

type ArmorInfo struct {
	BaseAC   int
	DexBonus string
	Category string
}

var ArmorData = map[string]ArmorInfo{
	"padded":          {BaseAC: 11, DexBonus: "full", Category: "light"},
	"leather":         {BaseAC: 11, DexBonus: "full", Category: "light"},
	"studded-leather": {BaseAC: 12, DexBonus: "full", Category: "light"},

	"hide":         {BaseAC: 12, DexBonus: "max2", Category: "medium"},
	"chain-shirt":  {BaseAC: 13, DexBonus: "max2", Category: "medium"},
	"scale-mail":   {BaseAC: 14, DexBonus: "max2", Category: "medium"},
	"breast-plate": {BaseAC: 14, DexBonus: "max2", Category: "medium"},
	"half-plate":   {BaseAC: 15, DexBonus: "max2", Category: "medium"},

	"ring-mail":  {BaseAC: 14, DexBonus: "none", Category: "heavy"},
	"chain-mail": {BaseAC: 16, DexBonus: "none", Category: "heavy"},
	"splint":     {BaseAC: 17, DexBonus: "none", Category: "heavy"},
	"plate":      {BaseAC: 18, DexBonus: "none", Category: "heavy"},
}

var (
	ErrSlotOccupied = errors.New("slot already occupied")
	ErrInvalidSlot  = errors.New("invalid slot")
	ErrNoWeaponName = errors.New("weapon name cannot be empty")
	ErrNoArmorName  = errors.New("armor name cannot be empty")
	ErrNoShieldName = errors.New("shield name cannot be empty")
)

func (c *Character) EquipWeapon(slot, weapon string) error {
	slot = strings.ToLower(strings.TrimSpace(slot))
	if weapon == "" {
		return ErrNoWeaponName
	}

	if slot != "main hand" && slot != "off hand" {
		return ErrInvalidSlot
	}

	if c.Equipment.Weapon == nil {
		c.Equipment.Weapon = make(map[string]string)
	}

	if existing, ok := c.Equipment.Weapon[slot]; ok && existing != "" {
		return ErrSlotOccupied
	}

	c.Equipment.Weapon[slot] = weapon
	return nil
}

func (c *Character) EquipArmor(armor string) error {
	if armor == "" {
		return ErrNoArmorName
	}

	c.Equipment.Armor = armor
	return nil
}

func (c *Character) EquipShield(shield string) error {
	if shield == "" {
		return ErrNoShieldName
	}

	c.Equipment.Shield = shield
	return nil
}

func (c *Character) CalculateInitiative() int {
	return c.Stats.DexMod
}

func (c *Character) CalculatePassivePerception() int {
	base := 10 + c.Stats.WisMod

	for _, skill := range c.Skills.Skills {
		if skill == "perception" {
			base += c.Proficiency
			break
		}
	}
	return base
}

func (c *Character) CalculateArmorClass() int {
	s := &c.Stats
	e := &c.Equipment
	class := strings.ToLower(c.Class)

	if e.Armor != "" {
		armorKey := SanitizeLocalKey(e.Armor)
		armorInfo, ok := ArmorData[armorKey]

		ac := 10 + s.DexMod
		if ok {
			ac = armorInfo.BaseAC
			switch armorInfo.DexBonus {
			case "full":
				ac += s.DexMod
			case "max2":
				dex := s.DexMod
				if dex > 2 {
					dex = 2
				}
				ac += dex
			case "none":
				// no bolus
			}
		}

		if e.Shield != "" {
			ac += 2
		}
		return ac
	}

	if class == "barbarian" {
		ac := 10 + s.DexMod + s.ConMod
		if e.Shield != "" {
			ac += 2
		}
		return ac
	}

	if class == "monk" {
		ac := 10 + s.DexMod + s.WisMod
		if e.Shield != "" {
			ac += 2
		}
		return ac
	}

	ac := 10 + s.DexMod
	if e.Shield != "" {
		ac += 2
	}
	return ac
}
