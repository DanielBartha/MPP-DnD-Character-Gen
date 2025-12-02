package class

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

type ClassRepository struct {
	all     []domain.WeaponInfo
	simple  []domain.WeaponInfo
	martial []domain.WeaponInfo

	classes map[string]ClassSkills
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
		all:     append([]domain.WeaponInfo{}, all...),
		simple:  append([]domain.WeaponInfo{}, simple...),
		martial: append([]domain.WeaponInfo{}, martial...),
	}
	cr.classes = cr.buildClasses()
	return cr
}

func (r *ClassRepository) GetCS(className string) (ClassSkills, bool) {
	key := strings.ToLower(strings.TrimSpace(className))
	cs, ok := r.classes[key]
	return cs, ok
}

func (r *ClassRepository) findByName(name string) (domain.WeaponInfo, bool) {
	lo := strings.ToLower(strings.TrimSpace(name))
	for _, w := range r.all {
		if strings.EqualFold(strings.TrimSpace(w.Name), lo) {
			return w, true
		}
	}
	return domain.WeaponInfo{Name: name}, false
}

func (r *ClassRepository) CombineWeaponSets(sets ...[]domain.WeaponInfo) []domain.WeaponInfo {
	out := make([]domain.WeaponInfo, 0)
	for _, s := range sets {
		out = append(out, s...)
	}
	return out
}

func (r *ClassRepository) withExtra(base []domain.WeaponInfo, names ...string) []domain.WeaponInfo {
	out := append([]domain.WeaponInfo{}, base...)
	for _, n := range names {
		if w, ok := r.findByName(n); ok {
			out = append(out, w)
		} else {
			out = append(out, domain.WeaponInfo{Name: n})
		}
	}
	return out
}

func (r *ClassRepository) WeaponsByName(names ...string) []domain.WeaponInfo {
	out := make([]domain.WeaponInfo, 0, len(names))
	for _, n := range names {
		if w, ok := r.findByName(n); ok {
			out = append(out, w)
		} else {
			out = append(out, domain.WeaponInfo{Name: n})
		}
	}
	return out
}

func (r *ClassRepository) weaponOrFallback(name string) domain.WeaponInfo {
	if w, ok := r.findByName(name); ok {
		return w
	}
	return domain.WeaponInfo{Name: name}
}

func (r *ClassRepository) buildClasses() map[string]ClassSkills {
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
			Weapons:  r.all,
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
			Weapons: r.withExtra(
				r.simple,
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
			Weapons:  r.simple,
			MainHand: "main hand",
			OffHand:  "off hand",
		},
		"druid": {
			MaxAllowed: 2,
			Skills: []string{
				"arcana", "animal handling", "insight", "medicine", "nature", "perception", "religion", "survival",
			},
			Armor: []string{
				"padded", "leather", "studded leather", "hide", "chain shirt", "scale mail", "breast plate", "half plate",
			},
			Shields: "shield",
			Weapons: append(r.simple,
				r.weaponOrFallback("club"),
				r.weaponOrFallback("greatclub"),
				r.weaponOrFallback("dagger"),
				r.weaponOrFallback("dart"),
				r.weaponOrFallback("javelins"),
				r.weaponOrFallback("maces"),
				r.weaponOrFallback("quarterstaff"),
				r.weaponOrFallback("scimitar"),
				r.weaponOrFallback("sickle"),
				r.weaponOrFallback("sling"),
				r.weaponOrFallback("spear"),
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
			Weapons:  r.simple,
			MainHand: "main hand",
			OffHand:  "off hand",
		},
		"monk": {
			MaxAllowed: 2,
			Skills: []string{
				"acrobatics", "athletics", "history", "insight", "religion", "stealth",
			},
			// NO ARMOR
			Weapons: append(r.simple,
				r.weaponOrFallback("shortsword"),
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
			Weapons:  r.all,
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
			Weapons:  r.all,
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
			Weapons: r.withExtra(
				r.simple,
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
			Weapons: r.WeaponsByName(
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
			Weapons:  r.simple,
			MainHand: "main hand",
			OffHand:  "off hand",
		},
		"wizard": {
			MaxAllowed: 2,
			Skills: []string{
				"arcana", "history", "insight", "investigation", "medicine", "religion",
			},
			// NO ARMOR
			Weapons: r.WeaponsByName(
				"dagger", "dart", "sling", "quarterstaff", "light crossbow",
			),
			MainHand: "main hand",
			OffHand:  "off hand",
		},
	}
}
