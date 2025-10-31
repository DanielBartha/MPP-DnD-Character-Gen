package domain_test

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestApplyRacialSkillProficiencies(t *testing.T) {
	tests := []struct {
		name     string
		race     string
		initial  []string
		expected []string
	}{
		{name: "Dwarf adds history",
			race:     "Dwarf",
			initial:  []string{"acrobatics", "athletics", "deception", "insight"},
			expected: []string{"acrobatics", "athletics", "deception", "insight", "history"},
		},
		{
			name:     "Half-orc add intimidation",
			race:     "Half-orc",
			initial:  []string{"animal handling", "athletics", "insight", "religion"},
			expected: []string{"animal handling", "athletics", "insight", "religion", "intimidation"},
		},
		{
			name:     "Unknown race does nothing",
			race:     "Unknown",
			initial:  []string{""},
			expected: []string{""},
		},
		{
			name:     "Duplicate skills allowed",
			race:     "Dwarf",
			initial:  []string{"history"},
			expected: []string{"history", "history"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			char := &domain.Character{
				Race: tt.race,
				Skills: domain.ClassLoadout{
					Skills: append([]string{}, tt.initial...),
				},
			}

			char.ApplyRacialSkillProficiencies()

			if len(char.Skills.Skills) != len(tt.expected) {
				t.Fatalf("expected %v skills, got %v", tt.expected, char.Skills.Skills)
			}

			for i, exp := range tt.expected {
				if char.Skills.Skills[i] != exp {
					t.Errorf("expected %v at %d, got %v", exp, i, char.Skills.Skills[i])
				}
			}
		})
	}
}
