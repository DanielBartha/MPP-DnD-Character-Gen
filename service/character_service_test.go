package service

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain/class"
)

func fakeWeapons() ([]domain.WeaponInfo, []domain.WeaponInfo, []domain.WeaponInfo) {
	all := []domain.WeaponInfo{{Name: "club"}, {Name: "dagger"}}
	return all, all, all
}

func TestGetClassSkills_ValidClass(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := class.NewClassRepository(all, simple, martial)
	svc := NewCharacterService(repo)

	char := &domain.Character{Class: "barbarian"}

	loadout := svc.GetClassSkills(char)

	if len(loadout.Skills) == 0 {
		t.Fatalf("expected skills for barbarian, got none")
	}

	// barbarian MaxAllowed = 2 - first 2 skills, plus "insight","religion"
	if got := loadout.Skills[0]; got != "animal handling" {
		t.Errorf("expected first skill animal handling, got %s", got)
	}
	if got := loadout.Skills[1]; got != "athletics" {
		t.Errorf("expected second skill athletics, got %s", got)
	}

	if !contains(loadout.Skills, "insight") {
		t.Errorf("expected racial/class insight to be included")
	}
	if !contains(loadout.Skills, "religion") {
		t.Errorf("expected racial/class religion to be included")
	}
}

func TestGetClassSkills_UnknownClass(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := class.NewClassRepository(all, simple, martial)
	svc := NewCharacterService(repo)

	char := &domain.Character{Class: "nonsense"}

	loadout := svc.GetClassSkills(char)

	if loadout.MaxAllowed != 0 {
		t.Errorf("expected MaxAllowed 0, got %d", loadout.MaxAllowed)
	}

	expected := []string{"insight", "religion"}
	if len(loadout.Skills) != len(expected) {
		t.Fatalf("expected only default skills, got %v", loadout.Skills)
	}

	for _, skill := range expected {
		if !contains(loadout.Skills, skill) {
			t.Errorf("expected %s in skills, got %v", skill, loadout.Skills)
		}
	}
}

func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}
