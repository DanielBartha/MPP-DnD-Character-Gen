package service

import (
	"os"
	"testing"
)

func TestSplitAndTrim(t *testing.T) {
	in := "Wizard, Sorcerer ,  Cleric"
	want := []string{"wizard", "sorcerer", "cleric"}

	got := splitAndTrim(in)

	if len(got) != len(want) {
		t.Fatalf("expected %d items, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("expected %s, got %s", want[i], got[i])
		}
	}
}

func TestGetSpellLevel(t *testing.T) {
	SpellDB = map[string]SpellMetadata{
		"fire bolt": {Level: 0},
	}

	level, err := GetSpellLevel("fire bolt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if level != 0 {
		t.Errorf("expected level 0, got %d", level)
	}

	_, err = GetSpellLevel("unknown spell")
	if err == nil {
		t.Fatal("expected error for unknown spell")
	}
}

func TestIsSpellForClass(t *testing.T) {
	SpellDB = map[string]SpellMetadata{
		"cure wounds": {Level: 1, Classes: []string{"cleric", "druid"}},
	}

	if !IsSpellForClass("cure wounds", "Cleric") {
		t.Error("expected cleric to be allowed")
	}

	if IsSpellForClass("cure wounds", "wizard") {
		t.Error("expected wizard NOT to be allowed")
	}
}

func TestLoadSpellsCSV(t *testing.T) {
	content := `name,level,class
Acid Splash,0,"Sorcerer,Wizard"
Cure Wounds,1,"Cleric,Druid"
`
	tmp := "test_spells.csv"
	os.WriteFile(tmp, []byte(content), 0644)
	defer os.Remove(tmp)

	err := LoadSpellsCSV(tmp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	spell, ok := SpellDB["acid splash"]
	if !ok {
		t.Fatal("expected acid splash entry")
	}
	if spell.Level != 0 {
		t.Errorf("expected level 0, got %d", spell.Level)
	}
	if len(spell.Classes) == 0 || spell.Classes[0] != "sorcerer" {
		t.Errorf("expected sorcerer in class list")
	}
}
