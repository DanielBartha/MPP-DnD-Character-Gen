package domain_test

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestApplyRacialSkillProficiencies_Dwarf(t *testing.T) {
	c := &domain.Character{
		Race: "dwarf",
		Skills: domain.ClassLoadout{
			Skills: []string{"insight", "religion"},
		},
	}

	c.ApplyRacialSkillProficiencies()

	want := []string{"insight", "religion", "history"}

	if len(c.Skills.Skills) != len(want) {
		t.Fatalf("expected %v skills, got %v", want, c.Skills.Skills)
	}

	for i, v := range want {
		if c.Skills.Skills[i] != v {
			t.Errorf("expected %s, got %s", v, c.Skills.Skills[i])
		}
	}
}
