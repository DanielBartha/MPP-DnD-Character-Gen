package service

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestGetSlotsForClassLevel_FullCaster(t *testing.T) {
	slots, casterType, err := GetSlotsForClassLevel("Wizard", 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if casterType != "full" {
		t.Errorf("expected full caster, got %s", casterType)
	}
	if slots[1] != 2 {
		t.Errorf("expected 2 level-1 slots, got %d", slots[1])
	}
}

func TestGetSlotsForClassLevel_HalfCaster(t *testing.T) {
	slots, casterType, err := GetSlotsForClassLevel("Paladin", 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if casterType != "half" {
		t.Errorf("expected half caster, got %s", casterType)
	}
	if slots[1] != 4 || slots[2] != 2 {
		t.Errorf("unexpected slot table for paladin lvl 5: %#v", slots)
	}
}

func TestGetSlotsForClassLevel_PactCaster(t *testing.T) {
	slots, casterType, err := GetSlotsForClassLevel("Warlock", 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if casterType != "pact" {
		t.Errorf("expected pact caster, got %s", casterType)
	}
	if slots[2] != 2 {
		t.Errorf("expected 2 lvl-2 slots, got %d", slots[2])
	}
}

func TestGetSlotsForClassLevel_NonCaster(t *testing.T) {
	_, _, err := GetSlotsForClassLevel("Fighter", 1)
	if err == nil {
		t.Fatal("expected error for non-caster fighter")
	}
}

func TestGetSlotsForClassLevel_InvalidLevel(t *testing.T) {
	_, _, err := GetSlotsForClassLevel("Wizard", 99)
	if err == nil {
		t.Fatal("expected error for invalid level")
	}
}

func TestGetSpellcastingAbility(t *testing.T) {
	tests := map[string]string{
		"wizard":   "intelligence",
		"cleric":   "wisdom",
		"sorcerer": "charisma",
		"fighter":  "",
	}

	for class, want := range tests {
		got := GetSpellcastingAbility(class)
		if got != want {
			t.Errorf("class %s: want %s, got %s", class, want, got)
		}
	}
}

func TestInitSpellcasting_FullCaster(t *testing.T) {
	svc := &CharacterService{}
	c := &domain.Character{
		Class:       "Wizard",
		Level:       1,
		Stats:       domain.Stats{IntelMod: 3},
		Proficiency: 2,
	}

	svc.InitSpellcasting(c)

	if !c.Spellcasting.CanCast {
		t.Fatal("wizard should be able to cast")
	}
	if c.Spellcasting.SpellSaveDC != 8+2+3 {
		t.Errorf("wrong DC: %d", c.Spellcasting.SpellSaveDC)
	}
	if c.Spellcasting.CantripsKnown != 3 {
		t.Errorf("expected 3 cantrips lvl1 wizard, got %d", c.Spellcasting.CantripsKnown)
	}
}

func TestInitSpellcasting_HalfCaster_Level1_NoCast(t *testing.T) {
	svc := &CharacterService{}
	c := &domain.Character{
		Class: "Paladin",
		Level: 1,
	}

	svc.InitSpellcasting(c)

	if c.Spellcasting.CanCast {
		t.Error("paladin lvl1 should NOT cast")
	}
}

func TestInitSpellcasting_PactCaster(t *testing.T) {
	svc := &CharacterService{}
	c := &domain.Character{
		Class:       "Warlock",
		Level:       2,
		Stats:       domain.Stats{ChaMod: 2},
		Proficiency: 2,
	}

	svc.InitSpellcasting(c)

	if !c.Spellcasting.CanCast {
		t.Fatal("warlock lvl2 should cast")
	}
	if c.Spellcasting.Slots[1] != 2 {
		t.Errorf("expected 2 lvl-1 pact slots, got %d", c.Spellcasting.Slots[1])
	}
	if c.Spellcasting.CantripsKnown != 2 {
		t.Errorf("warlock lvl2 should know 2 cantrips")
	}
}

func TestCopyIntMap(t *testing.T) {
	orig := map[int]int{1: 2}
	copy := copyIntMap(orig)

	orig[1] = 99

	if copy[1] == 99 {
		t.Fatal("copy changed when original changed (map not copied)")
	}
}
