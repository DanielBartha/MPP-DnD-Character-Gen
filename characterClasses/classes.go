package characterClasses

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

var (
	AllWeapons     []domain.WeaponInfo
	SimpleWeapons  []domain.WeaponInfo
	MartialWeapons []domain.WeaponInfo
)

type ClassSkills struct {
	MaxAllowed int
	Skills     []string
	Armor      []string
	Shields    string
	Weapons    []domain.WeaponInfo
	MainHand   string
	OffHand    string
}

func InitWeapons(all, simple, martial []domain.WeaponInfo) {
	AllWeapons = append([]domain.WeaponInfo{}, all...)
	SimpleWeapons = append([]domain.WeaponInfo{}, simple...)
	MartialWeapons = append([]domain.WeaponInfo{}, martial...)
}

func findWeaponByName(name string) domain.WeaponInfo {
	for _, w := range AllWeapons {
		if strings.EqualFold(w.Name, name) {
			return w
		}
	}
	return domain.WeaponInfo{Name: name}
}

func CombineWeaponSets(sets ...[]domain.WeaponInfo) []domain.WeaponInfo {
	out := make([]domain.WeaponInfo, 0)
	for _, s := range sets {
		out = append(out, s...)
	}
	return out
}

func WithExtraWeapons(base []domain.WeaponInfo, names ...string) []domain.WeaponInfo {
	out := append([]domain.WeaponInfo{}, base...)
	for _, n := range names {
		out = append(out, findWeaponByName(n))
	}
	return out
}

func WeaponsByName(names ...string) []domain.WeaponInfo {
	out := make([]domain.WeaponInfo, 0, len(names))
	for _, n := range names {
		out = append(out, findWeaponByName(n))
	}
	return out
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
		Weapons:  AllWeapons,
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
		Weapons: WithExtraWeapons(
			SimpleWeapons,
			"hand crossbow", "longsword", "rapier", "shortsword",
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
		Weapons:  SimpleWeapons,
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
		Weapons: append(SimpleWeapons,
			findWeaponByName("club"),
			findWeaponByName("greatclub"),
			findWeaponByName("dagger"),
			findWeaponByName("dart"),
			findWeaponByName("javelins"),
			findWeaponByName("maces"),
			findWeaponByName("quarterstaff"),
			findWeaponByName("scimitar"),
			findWeaponByName("sickle"),
			findWeaponByName("sling"),
			findWeaponByName("spear"),
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
		Weapons:  AllWeapons,
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"monk": {
		MaxAllowed: 2,
		Skills: []string{
			"acrobatics", "athletics", "history", "insight", "religion", "stealth",
		},
		// NO ARMOR
		Weapons: append(SimpleWeapons,
			findWeaponByName("shortsword"),
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
		Weapons:  AllWeapons,
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
		Weapons:  AllWeapons,
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"rogue": {
		MaxAllowed: 4,
		Skills: []string{
			"acrobatics", "athletics", "deception", "insight", "intimidation", "investigation", "perception", "performance", "persuasion", "sleight of hand", "stealth",
		},
		Armor: []string{
			"padded", "leather", "studded leather",
		},
		Weapons: WithExtraWeapons(
			SimpleWeapons,
			"hand crossbow", "longsword", "rapier", "shortsword",
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
		Weapons: WeaponsByName(
			"dagger", "dart", "sling", "quarterstaff", "light crossbow",
		),
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
		Weapons:  SimpleWeapons,
		MainHand: "main hand",
		OffHand:  "off hand",
	},
	"wizard": {
		MaxAllowed: 2,
		Skills: []string{
			"arcana", "history", "insight", "investigation", "medicine", "religion",
		},
		// NO ARMOR
		Weapons: WeaponsByName(
			"dagger", "dart", "sling", "quarterstaff", "light crossbow",
		),
		MainHand: "main hand",
		OffHand:  "off hand",
	},
}
