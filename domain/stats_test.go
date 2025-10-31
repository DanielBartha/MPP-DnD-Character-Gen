package domain_test

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestAbilityModifier(t *testing.T) {
	tests := []struct {
		score int
		want  int
	}{
		{10, 0},
		{8, -1},
		{14, 2},
		{17, 3},
	}

	for _, tt := range tests {
		stats := domain.Stats{
			Str: tt.score,
		}
		c := &domain.Character{Stats: stats}
		c.UpdateProficiency()

		if c.Stats.StrMod != tt.want {
			t.Errorf("Str=%d → want mod %d, got %d", tt.score, tt.want, c.Stats.StrMod)
		}
	}
}
