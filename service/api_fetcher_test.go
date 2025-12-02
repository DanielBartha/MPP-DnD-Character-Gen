package service

import (
	"testing"
)

func TestToDomainSpell(t *testing.T) {
	input := &apiSpellResp{
		Name:  "Fire Bolt",
		Level: 0,
		School: struct {
			Name string `json:"name"`
		}{Name: "Evocation"},
		Range: "120 feet",
	}

	got := ToDomainSpell(input)

	if got.Name != "Fire Bolt" {
		t.Errorf("expected spell name Fire Bolt, got %s", got.Name)
	}
	if got.School != "Evocation" {
		t.Errorf("expected school Evocation, got %s", got.School)
	}
	if got.Range != "120 feet" {
		t.Errorf("expected range 120 feet, got %s", got.Range)
	}
}

func TestToDomainWeapon(t *testing.T) {
	input := &apiWeaponResp{
		Name:           "Longsword",
		WeaponCategory: "Martial",
		Range: struct {
			Normal int `json:"normal"`
			Long   int `json:"long"`
		}{Normal: 5, Long: 0},
		TwoHanded: false,
	}

	got := ToDomainWeapon(input)

	if got.Name != "Longsword" {
		t.Errorf("expected name Longsword, got %s", got.Name)
	}
	if got.Category != "Martial" {
		t.Errorf("expected category Martial, got %s", got.Category)
	}
	if got.Range != 5 {
		t.Errorf("expected range 5, got %d", got.Range)
	}
	if got.TwoHanded != false {
		t.Errorf("expected TwoHanded false, got %v", got.TwoHanded)
	}
}

func TestToDomainArmor(t *testing.T) {
	input := &apiArmorResp{
		Name: "Chain Mail",
		ArmorClass: struct {
			Base     int  `json:"base"`
			DexBonus bool `json:"dex_bonus"`
			MaxBonus int  `json:"max_bonus"`
		}{Base: 16, DexBonus: false, MaxBonus: 0},
	}

	got := ToDomainArmor(input)

	if got.Name != "Chain Mail" {
		t.Errorf("expected Chain Mail, got %s", got.Name)
	}
	if got.BaseAC != 16 {
		t.Errorf("expected BaseAC 16, got %d", got.BaseAC)
	}
	if got.DexBonus != false {
		t.Errorf("expected DexBonus false, got %v", got.DexBonus)
	}
	if got.MaxBonus != 0 {
		t.Errorf("expected MaxBonus 0, got %d", got.MaxBonus)
	}
}
