package class

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

type ClassRepository struct {
	AllWeapons     []domain.WeaponInfo
	SimpleWeapons  []domain.WeaponInfo
	MartialWeapons []domain.WeaponInfo
	Classes        map[string]ClassSkills
}

type ClassSkills struct {
	MaxAllowed int
	Skills     []string
	Armor      []string
	Shields    string
	Weapons    []domain.WeaponInfo
	MainHand   string
	OffHand    string
}

func NewClassRepository(all, simple, martial []domain.WeaponInfo) *ClassRepository {
	cr := &ClassRepository{
		AllWeapons:     append([]domain.WeaponInfo{}, all...),
		SimpleWeapons:  append([]domain.WeaponInfo{}, simple...),
		MartialWeapons: append([]domain.WeaponInfo{}, martial...),
	}
	cr.Classes = cr.buildClasses()
	return cr
}

func (cr *ClassRepository) findWeaponByName(name string) domain.WeaponInfo {
	for _, w := range cr.AllWeapons {
		if strings.EqualFold(w.Name, name) {
			return w
		}
	}
	return domain.WeaponInfo{Name: name}
}

func (cr *ClassRepository) CombineWeaponSets(sets ...[]domain.WeaponInfo) []domain.WeaponInfo {
	out := make([]domain.WeaponInfo, 0)
	for _, s := range sets {
		out = append(out, s...)
	}
	return out
}

func (cr *ClassRepository) WithExtraWeapons(base []domain.WeaponInfo, names ...string) []domain.WeaponInfo {
	out := append([]domain.WeaponInfo{}, base...)
	for _, n := range names {
		out = append(out, cr.findWeaponByName(n))
	}
	return out
}

func (cr *ClassRepository) WeaponsByName(names ...string) []domain.WeaponInfo {
	out := make([]domain.WeaponInfo, 0, len(names))
	for _, n := range names {
		out = append(out, cr.findWeaponByName(n))
	}
	return out
}

func (cr *ClassRepository) buildClasses() map[string]ClassSkills {
	return map[string]ClassSkills{
		"barbarian": {
			MaxAllowed: 2,
			Skills: []string{
				"animal handling", "athletics", "intimidation", "nature", "perception", "survival",
			},
			Armor: []string{
				"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
			},
			Shields:  "shield",
			Weapons:  cr.AllWeapons,
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
			Weapons: cr.WithExtraWeapons(
				cr.SimpleWeapons,
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
			Weapons:  cr.SimpleWeapons,
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
			Weapons: append(cr.SimpleWeapons,
				cr.findWeaponByName("club"),
				cr.findWeaponByName("greatclub"),
				cr.findWeaponByName("dagger"),
				cr.findWeaponByName("dart"),
				cr.findWeaponByName("javelins"),
				cr.findWeaponByName("maces"),
				cr.findWeaponByName("quarterstaff"),
				cr.findWeaponByName("scimitar"),
				cr.findWeaponByName("sickle"),
				cr.findWeaponByName("sling"),
				cr.findWeaponByName("spear"),
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
			Weapons:  cr.AllWeapons,
			MainHand: "main hand",
			OffHand:  "off hand",
		},
		"monk": {
			MaxAllowed: 2,
			Skills: []string{
				"acrobatics", "athletics", "history", "insight", "religion", "stealth",
			},
			// NO ARMOR
			Weapons: append(cr.SimpleWeapons,
				cr.findWeaponByName("shortsword"),
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
			Weapons:  cr.AllWeapons,
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
			Weapons:  cr.AllWeapons,
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
			Weapons: cr.WithExtraWeapons(
				cr.SimpleWeapons,
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
			Weapons: cr.WeaponsByName(
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
			Weapons:  cr.SimpleWeapons,
			MainHand: "main hand",
			OffHand:  "off hand",
		},
		"wizard": {
			MaxAllowed: 2,
			Skills: []string{
				"arcana", "history", "insight", "investigation", "medicine", "religion",
			},
			// NO ARMOR
			Weapons: cr.WeaponsByName(
				"dagger", "dart", "sling", "quarterstaff", "light crossbow",
			),
			MainHand: "main hand",
			OffHand:  "off hand",
		},
	}
}
