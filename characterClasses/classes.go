package characterClasses

import "github.com/DanielBartha/MPP-DnD-Character-Gen/weapons"

type ClassSkills struct {
	MaxAllowed int
	Skills     []string
	Armor      []string
	Shields    string
	Weapons    []weapons.Weapon
	MainHand   string
	OffHand    string
}

var Classes = map[string]ClassSkills{
	"barbarian": {
		MaxAllowed: 2,
		Skills: []string{
			"animal handling", "athletics", "intimidation", "nature", "perception", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		Weapons:  weapons.GetAllWeapons(),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"bard": {
		MaxAllowed: 3,
		Skills: []string{
			"acrobatics", "animal handling", "arcana", "athletics", "deception",
			"history", "insight", "intimidation", "investigation", "medicine", "nature",
			"perception", "performance", "persuasion", "religion", "sleight of hand", "stealth", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		Weapons: append(
			append([]weapons.Weapon{}, weapons.SimpleMelee...),
			append(
				weapons.SimpleRanged,
				weapons.Weapon{Name: "hand crossbow", TwoHanded: false},
				weapons.Weapon{Name: "longsword", TwoHanded: false},
				weapons.Weapon{Name: "rapier", TwoHanded: false},
				weapons.Weapon{Name: "shortsword", TwoHanded: false},
			)...,
		),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"cleric": {
		MaxAllowed: 2,
		Skills: []string{
			"history", "insight", "medicine", "persuasion", "religion",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		Weapons:  weapons.GetSimpleWeapons(),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"druid": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "animal handling", "insight", "medicine", "nature", "perception", "religion", "survival",
		},
		// druids don't wear armor or shields made of metal, fucking animals
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields: "shield",
		Weapons: append(
			weapons.SimpleRanged,
			weapons.Weapon{Name: "club", TwoHanded: false},
			weapons.Weapon{Name: "greatclub", TwoHanded: true},
			weapons.Weapon{Name: "dagger", TwoHanded: false},
			weapons.Weapon{Name: "dart", TwoHanded: false},
			weapons.Weapon{Name: "javelins", TwoHanded: false},
			weapons.Weapon{Name: "maces", TwoHanded: false},
			weapons.Weapon{Name: "quarterstaff", TwoHanded: false},
			weapons.Weapon{Name: "scimitar", TwoHanded: false},
			weapons.Weapon{Name: "sickle", TwoHanded: false},
			weapons.Weapon{Name: "sling", TwoHanded: true},
			weapons.Weapon{Name: "spear", TwoHanded: false},
		),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"fighter": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "animal handling", "athletics", "history", "insight", "intimidation", "perception", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate", "ring mail", "chain mail", "splint", "plate",
		},
		Shields:  "shield",
		Weapons:  weapons.GetAllWeapons(),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"monk": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "athletics", "history", "insight", "religion", "stealth",
		},
		// NO ARMOR
		Weapons: append(weapons.GetSimpleWeapons(),
			weapons.Weapon{Name: "shortsword", TwoHanded: false},
		),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"paladin": {
		MaxAllowed: 2,
		Skills: []string{
			"athletics", "insight", "intimidation", "medicine", "persuasion", "religion",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate", "ring mail", "chain mail", "splint", "plate",
		},
		Shields:  "shield",
		Weapons:  weapons.GetAllWeapons(),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"ranger": {
		MaxAllowed: 3,
		Skills: []string{
			"animal handling", "athletics", "insight", "investigation", "nature", "perception", "stealth", "survival",
		},
		Armor: []string{
			"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
		},
		Shields:  "shield",
		Weapons:  weapons.GetAllWeapons(),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"rogue": {
		MaxAllowed: 4,
		Skills: []string{
			"acrobatics", "athletcis", "deception", "insight", "intimidation", "investigation", "perception", "performance", "persuasion", "sleight of hand", "stealth",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		Weapons: append(weapons.GetSimpleWeapons(),
			weapons.Weapon{Name: "hand crossbow", TwoHanded: false},
			weapons.Weapon{Name: "longsword", TwoHanded: false},
			weapons.Weapon{Name: "rapier", TwoHanded: false},
			weapons.Weapon{Name: "shortsword", TwoHanded: false},
		),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"sorcerer": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "deception", "insight", "intimidation", "persuasion", "religion",
		},
		// NO ARMOR
		Weapons: []weapons.Weapon{
			{Name: "dagger", TwoHanded: false},
			{Name: "dart", TwoHanded: false},
			{Name: "sling", TwoHanded: true},
			{Name: "quarterstaff", TwoHanded: false},
			{Name: "light crossbow", TwoHanded: true},
		},
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"warlock": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "deception", "history", "intimidation", "investigation", "nature", "religion",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		Weapons:  weapons.GetSimpleWeapons(),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"wizard": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "history", "insight", "investigation", "medicine", "religion",
		},
		// NO ARMOR
		Weapons: []weapons.Weapon{
			{Name: "dagger", TwoHanded: false},
			{Name: "dart", TwoHanded: false},
			{Name: "sling", TwoHanded: true},
			{Name: "quarterstaff", TwoHanded: false},
			{Name: "light crossbow", TwoHanded: true},
		},
		MainHand: "main hand",
		OffHand:  "off hand",
	},
}
