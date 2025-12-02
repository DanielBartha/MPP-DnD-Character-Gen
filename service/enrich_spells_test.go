package service

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestEnrichSpells_BasicFlow(t *testing.T) {
	input := "test_spells_input.csv"
	output := "test_spells_output.csv"
	defer os.Remove(input)
	defer os.Remove(output)

	// header + 2 roes
	content := "Name\nFire Bolt\nUnknown Spell\n"
	if err := os.WriteFile(input, []byte(content), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}

	orig := fetchSpellsBatchFn

	fetchSpellsBatchFn = func(indexes []string) map[string]*apiSpellResp {
		resp := &apiSpellResp{
			Name:  "Fire Bolt",
			Range: "120 feet",
		}
		resp.School.Name = "Evocation"

		return map[string]*apiSpellResp{
			"fire-bolt": resp,
		}
	}

	defer func() { fetchSpellsBatchFn = orig }()

	if err := EnrichSpells(input, output); err != nil {
		t.Fatalf("EnrichSpells error: %v", err)
	}

	f, err := os.Open(output)
	if err != nil {
		t.Fatalf("open output: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		t.Fatalf("read csv: %v", err)
	}

	if len(rows) != 3 {
		t.Fatalf("expected 3 rows, got %d", len(rows))
	}

	if rows[1][1] != "Evocation" || rows[1][2] != "120 feet" {
		t.Errorf("Fire Bolt enrichment failed: %v", rows[1])
	}

	if rows[2][1] != "N/A" || rows[2][2] != "N/A" {
		t.Errorf("Unknown spell should have N/A values: %v", rows[2])
	}
}
