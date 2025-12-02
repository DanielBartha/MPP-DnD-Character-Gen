package service

import "github.com/DanielBartha/MPP-DnD-Character-Gen/domain"

type apiArmorDomain struct {
	Name     string
	BaseAC   int
	DexBonus bool
	MaxBonus int
}

func ToDomainSpell(r *apiSpellResp) *domain.SpellInfo {
	return &domain.SpellInfo{
		Name:   r.Name,
		School: r.School.Name,
		Range:  r.Range,
	}
}

func ToDomainWeapon(r *apiWeaponResp) *domain.WeaponInfo {
	return &domain.WeaponInfo{
		Name:      r.Name,
		Category:  r.WeaponCategory,
		Range:     r.Range.Normal,
		TwoHanded: r.TwoHanded,
	}
}

func ToDomainArmor(r *apiArmorResp) *apiArmorDomain {
	return &apiArmorDomain{
		Name:     r.Name,
		BaseAC:   r.ArmorClass.Base,
		DexBonus: r.ArmorClass.DexBonus,
		MaxBonus: r.ArmorClass.MaxBonus,
	}
}
