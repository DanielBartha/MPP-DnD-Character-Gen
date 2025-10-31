package service

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestEnrichArmor_BasicFlow(t *testing.T) {
	input := "test_input.csv"
	output := "test_output.csv"
	defer os.Remove(input)
	defer os.Remove(output)

	// header + one armor row + one non-armor row
	if err := os.WriteFile(input, []byte("Name,Type\nLeather,Armor\nRope,Gear\n"), 0o644); err != nil {
		t.Fatalf("write input: %v", err)
	}

	orig := fetchArmorBatchFn
	fetchArmorBatchFn = func(indexes []string) map[string]*apiArmorResp {

		resp := &apiArmorResp{Name: "Leather"}
		resp.ArmorClass.Base = 11
		resp.ArmorClass.DexBonus = true
		resp.ArmorClass.MaxBonus = 0

		return map[string]*apiArmorResp{
			"leather": resp,
		}
	}
	defer func() { fetchArmorBatchFn = orig }()

	if err := EnrichArmor(input, output); err != nil {
		t.Fatalf("EnrichArmor error: %v", err)
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

	if got := len(rows[0]); got < 5 {
		t.Errorf("expected extended header, got %d columns", got)
	}

	// leather armor row: base_ac should be 11
	// columns: Name,Type,base_ac,dex_bonus,max_bonus
	if rows[1][2] != "11" {
		t.Errorf("expected leather base_ac=11, got %s", rows[1][2])
	}

	if rows[2][2] != "N/A" || rows[2][3] != "N/A" || rows[2][4] != "N/A" {
		t.Errorf("expected N/A in appended cols for non-armor row, got %v", rows[2])
	}
}
