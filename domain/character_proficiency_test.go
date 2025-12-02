package domain

import "testing"

func TestUpdateProficiency(t *testing.T) {
	tests := []struct {
		level    int
		wantProf int
	}{
		{1, 2},
		{4, 2},
		{5, 3},
		{8, 3},
		{9, 4},
		{12, 4},
		{13, 5},
		{16, 5},
		{17, 6},
		{20, 6},
	}

	for _, tt := range tests {
		c := &Character{Level: tt.level}
		c.UpdateProficiency()

		if c.Proficiency != tt.wantProf {
			t.Errorf("Level %d: expected prof %d, got %d",
				tt.level, tt.wantProf, c.Proficiency)
		}
	}
}
