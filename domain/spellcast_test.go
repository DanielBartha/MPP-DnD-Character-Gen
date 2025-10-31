package domain_test

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestGetSpellcastingType(t *testing.T) {
	tests := []struct {
		class string
		want  string
	}{
		{"Bard", "full"},
		{"cleric", "full"},
		{"Druid", "full"},
		{"sorcerer", "full"},
		{"wizard", "full"},

		{"paladin", "half"},
		{"Ranger", "half"},

		{"Warlock", "pact"},

		{"fighter", ""},
		{"barbarian", ""},
		{"rogue", ""},
		{"", ""},
	}

	for _, tt := range tests {
		got := domain.GetSpellcastingType(tt.class)
		if got != tt.want {
			t.Errorf("GetSpellcastingType(%q) = %q, want %q",
				tt.class, got, tt.want)
		}
	}
}
