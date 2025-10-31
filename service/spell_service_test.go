package service

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func setupCaster() *domain.Character {
	return &domain.Character{
		Class: "wizard",
		Spellcasting: &domain.Spellcasting{
			CanCast:        true,
			Slots:          map[int]int{1: 2},
			MaxSlots:       map[int]int{1: 2},
			LearnedSpells:  []string{},
			PreparedSpells: []string{},
		},
	}
}

func TestLearnSpell_Success(t *testing.T) {
	SpellDB = map[string]SpellMetadata{
		"light": {Level: 0, Classes: []string{"wizard", "cleric"}},
	}
	c := setupCaster()
	svc := NewSpellService()

	msg, err := svc.LearnSpell(c, "light")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "Learned spell light" {
		t.Errorf("got: %s", msg)
	}
	if len(c.Spellcasting.LearnedSpells) != 1 {
		t.Errorf("spell not stored")
	}
}

func TestLearnSpell_AlreadyKnown(t *testing.T) {
	c := setupCaster()
	c.Spellcasting.LearnedSpells = []string{"light"}

	svc := NewSpellService()
	_, err := svc.LearnSpell(c, "light")
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}

func TestPrepareSpell_Success(t *testing.T) {
	c := setupCaster()
	svc := NewSpellService()

	msg, err := svc.PrepareSpell(c, "light")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if msg != "Prepared spell light" {
		t.Errorf("got: %s", msg)
	}
	if len(c.Spellcasting.PreparedSpells) != 1 {
		t.Errorf("spell not stored")
	}
}

func TestPrepareSpell_AlreadyPrepared(t *testing.T) {
	c := setupCaster()
	c.Spellcasting.PreparedSpells = []string{"light"}

	svc := NewSpellService()
	_, err := svc.PrepareSpell(c, "light")
	if err == nil {
		t.Fatal("expected error")
	}
}
