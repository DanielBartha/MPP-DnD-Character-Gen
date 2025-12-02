package service

import (
	"os"
	"strings"
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestEnrichWeapons(t *testing.T) {
	// fake csv
	inputCSV := `name,type
Longsword,Weapon
Stick,Weapon
Rock,Other`

	in := createTempFile(t, inputCSV)
	out := createTempFile(t, "")

	// mock API response
	original := fetchWeaponsBatchFn
	fetchWeaponsBatchFn = func(indexes []string) map[string]*apiWeaponResp {
		resp := &apiWeaponResp{
			Name:           "Longsword",
			WeaponCategory: "Martial",
		}
		resp.Range.Normal = 5
		resp.TwoHanded = false

		return map[string]*apiWeaponResp{
			"longsword": resp,
		}
	}
	defer func() { fetchWeaponsBatchFn = original }()

	err := EnrichWeapons(in, out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readFile(t, out)
	expectedContains := []string{
		"Longsword,Martial,5,false",
		"Stick,N/A,N/A,N/A",
	}
	for _, want := range expectedContains {
		if !strings.Contains(got, want) {
			t.Errorf("expected output to contain %q\nGot:\n%s", want, got)
		}
	}
}

func TestSimpleAndMartialWeapons(t *testing.T) {
	weapons := []domain.WeaponInfo{
		{Name: "Stick", Category: "Simple"},
		{Name: "Longsword", Category: "Martial"},
		{Name: "Bow", Category: "Simple"},
	}

	simple := SimpleWeapons(weapons)
	if len(simple) != 2 {
		t.Errorf("expected 2 simple weapons, got %d", len(simple))
	}

	martial := MartialWeapons(weapons)
	if len(martial) != 1 {
		t.Errorf("expected 1 martial weapon, got %d", len(martial))
	}
}

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "wep_in")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(content)
	f.Close()
	return f.Name()
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
