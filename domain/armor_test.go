package domain_test

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestArmorClass_NoArmor(t *testing.T) {
	c := &domain.Character{
		Class: "fighter",
		Stats: domain.Stats{DexMod: 2},
		Equipment: domain.Equipment{
			Armor:  "",
			Shield: "",
		},
	}

	got := c.CalculateArmorClass()
	want := 12 // 10 + Dex(2)

	if got != want {
		t.Errorf("AC: want %d got %d", want, got)
	}
}

func TestArmorClass_WithShield(t *testing.T) {
	c := &domain.Character{
		Class: "fighter",
		Stats: domain.Stats{DexMod: 2},
		Equipment: domain.Equipment{
			Armor:  "",
			Shield: "shield",
		},
	}

	got := c.CalculateArmorClass()
	want := 14 // 10 + Dex(2) + 2 shield

	if got != want {
		t.Errorf("AC: want %d got %d", want, got)
	}
}
