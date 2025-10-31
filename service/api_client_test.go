package service

import (
	"os"
	"path/filepath"
	"testing"
)

func fakeSpell(_ string) (*apiSpellResp, error) {
	return &apiSpellResp{Name: "Fire Bolt", Level: 0}, nil
}
func fakeWeapon(_ string) (*apiWeaponResp, error) {
	return &apiWeaponResp{Name: "Longsword"}, nil
}
func fakeArmor(_ string) (*apiArmorResp, error) {
	return &apiArmorResp{Name: "Chain Mail", ArmorClass: struct {
		Base     int  `json:"base"`
		DexBonus bool `json:"dex_bonus"`
		MaxBonus int  `json:"max_bonus"`
	}{Base: 16}}, nil
}

func TestFetchSpellsBatch(t *testing.T) {
	fetchSpellFunc = fakeSpell
	defer func() { fetchSpellFunc = FetchSpell }()

	idx := []string{"spell1", "spell2"}

	results := FetchSpellsBatch(idx)

	if len(results) != 2 {
		t.Fatalf("expected 2 spells, got %d", len(results))
	}

	if results["spell1"].Name != "Fire Bolt" {
		t.Errorf("unexpected spell result: %v", results["spell1"])
	}

	if _, err := os.Stat(filepath.Join("data", "api_cache", "spells.json")); err != nil {
		t.Errorf("cache file not written")
	}
}

func TestFetchWeaponsBatch(t *testing.T) {
	fetchWeaponFunc = fakeWeapon
	defer func() { fetchWeaponFunc = FetchWeapon }()

	idx := []string{"w1"}

	results := FetchWeaponsBatch(idx)

	if results["w1"].Name != "Longsword" {
		t.Errorf("unexpected weapon result: %v", results["w1"])
	}
}

func TestFetchArmorBatch(t *testing.T) {
	fetchArmorFunc = fakeArmor
	defer func() { fetchArmorFunc = FetchArmor }()

	idx := []string{"a1"}

	results := FetchArmorBatch(idx)

	if results["a1"].Name != "Chain Mail" {
		t.Errorf("unexpected armor result: %v", results["a1"])
	}
}

func TestMain(m *testing.M) {
	os.RemoveAll("data/api_cache")
	os.MkdirAll("data/api_cache", 0755)
	code := m.Run()
	os.RemoveAll("data/api_cache")
	os.Exit(code)
}
