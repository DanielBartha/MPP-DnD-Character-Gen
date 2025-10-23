package service

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

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

func CalculateInitiative(s *domain.Stats) int {
	return s.DexMod
}

func CalculateArmorClass(c *domain.Character) int {
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
