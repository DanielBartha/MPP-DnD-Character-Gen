package service

import (
	"strings"
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestFormatCharacterView_BasicFields(t *testing.T) {

	char := &domain.Character{
		Name:       "TestHero",
		Race:       "Elf",
		Class:      "Wizard",
		Background: "Sage",
		Level:      3,
		Stats: domain.Stats{
			Str: 8, StrMod: -1,
			Dex: 14, DexMod: 2,
			Con: 12, ConMod: 1,
			Intel: 16, IntelMod: 3,
			Wis: 10, WisMod: 0,
			Cha: 13, ChaMod: 1,
		},
		Proficiency: 2,
		Skills: domain.ClassLoadout{
			Skills: []string{"arcana", "history"},
		},
		ArmorClass:        12,
		InitiativeBonus:   2,
		PassivePerception: 10,
	}

	out := FormatCharacterView(char)

	assertContains(t, out, "Name: TestHero")
	assertContains(t, out, "Class: wizard")
	assertContains(t, out, "Race: elf")
	assertContains(t, out, "Background: Sage")
	assertContains(t, out, "Level: 3")

	assertContains(t, out, "STR: 8 (-1)")
	assertContains(t, out, "DEX: 14 (+2)")
	assertContains(t, out, "INT: 16 (+3)")

	assertContains(t, out, "Skill proficiencies: arcana, history")
	assertContains(t, out, "Armor class: 12")
	assertContains(t, out, "Initiative bonus: 2")
	assertContains(t, out, "Passive perception: 10")
}

func TestFormatCharacterView_SpellcastingSection(t *testing.T) {
	char := &domain.Character{
		Name:  "Mage",
		Class: "Wizard",
		Stats: domain.Stats{},
		Spellcasting: &domain.Spellcasting{
			CanCast:          true,
			CantripsKnown:    3,
			MaxSlots:         map[int]int{1: 2, 2: 0},
			Ability:          "int",
			SpellSaveDC:      14,
			SpellAttackBonus: 6,
		},
	}

	out := FormatCharacterView(char)

	assertContains(t, out, "Spell slots:")
	assertContains(t, out, "Level 0: 3")
	assertContains(t, out, "Level 1: 2")
	assertContains(t, out, "Spellcasting ability: int")
	assertContains(t, out, "Spell save DC: 14")
	assertContains(t, out, "Spell attack bonus: +6")
}

func TestFormatCharacterView_Equipment(t *testing.T) {
	char := &domain.Character{
		Name:  "Fighter",
		Stats: domain.Stats{},
		Equipment: domain.Equipment{
			Weapon: map[string]string{
				"main hand": "longsword",
				"off hand":  "dagger",
			},
			Armor:  "chain mail",
			Shield: "shield",
		},
	}

	out := FormatCharacterView(char)

	assertContains(t, out, "Main hand: longsword")
	assertContains(t, out, "Off hand: dagger")
	assertContains(t, out, "Armor: chain mail")
	assertContains(t, out, "Shield: shield")
}

func assertContains(t *testing.T, out, sub string) {
	t.Helper()
	if !strings.Contains(out, sub) {
		t.Errorf("expected output to contain %q\nOutput:\n%s", sub, out)
	}
}
