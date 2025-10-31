package service

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestComputeDerivedStats(t *testing.T) {
	char := &domain.Character{
		Class: "fighter",
		Stats: domain.Stats{Dex: 14, Wis: 12, Con: 10},
	}
	char.UpdateProficiency()

	// no armor, no shield - AC = 10 + dex mod (2) = 12
	ComputeDerivedStats(char)

	if char.ArmorClass == 0 || char.InitiativeBonus == 0 || char.PassivePerception == 0 {
		t.Fatalf("derived stats were not computed: %+v", char)
	}

	if char.ArmorClass != 12 {
		t.Errorf("expected AC 12, got %d", char.ArmorClass)
	}
	if char.InitiativeBonus != 2 {
		t.Errorf("expected InitiativeBonus 2, got %d", char.InitiativeBonus)
	}
	if char.PassivePerception != 11 { // 10 + wis mod 1
		t.Errorf("expected PassivePerception 11, got %d", char.PassivePerception)
	}
}
