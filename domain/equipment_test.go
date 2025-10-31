package domain

import "testing"

func newChar() *Character {
	return &Character{
		Equipment: Equipment{Weapon: map[string]string{}},
		Stats:     Stats{DexMod: 2, WisMod: 1, ConMod: 3},
		Skills:    ClassLoadout{Skills: []string{}},
	}
}

func TestEquipWeapon(t *testing.T) {
	c := newChar()

	if err := c.EquipWeapon("main hand", "sword"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := c.EquipWeapon("main hand", "axe"); err != ErrSlotOccupied {
		t.Errorf("expected ErrSlotOccupied, got %v", err)
	}

	if err := c.EquipWeapon("foot", "dagger"); err != ErrInvalidSlot {
		t.Errorf("expected ErrInvalidSlot, got %v", err)
	}

	if err := c.EquipWeapon("off hand", ""); err != ErrNoWeaponName {
		t.Errorf("expected ErrNoWeaponName, got %v", err)
	}
}

func TestEquipArmorAndShield(t *testing.T) {
	c := newChar()

	if err := c.EquipArmor(""); err != ErrNoArmorName {
		t.Errorf("expected ErrNoArmorName, got %v", err)
	}
	if err := c.EquipArmor("chain mail"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if c.Equipment.Armor != "chain mail" {
		t.Errorf("armor not set correctly")
	}

	if err := c.EquipShield(""); err != ErrNoShieldName {
		t.Errorf("expected ErrNoShieldName, got %v", err)
	}
	if err := c.EquipShield("shield"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if c.Equipment.Shield != "shield" {
		t.Errorf("shield not set correctly")
	}
}

func TestCalculateInitiative(t *testing.T) {
	c := newChar()
	if got := c.CalculateInitiative(); got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

func TestCalculatePassivePerception(t *testing.T) {
	// without perception proficiency
	c := newChar()
	if got := c.CalculatePassivePerception(); got != 11 {
		t.Errorf("expected 11, got %d", got)
	}

	// with perception proficiency
	c.Skills.Skills = []string{"perception"}
	c.Proficiency = 2
	if got := c.CalculatePassivePerception(); got != 13 {
		t.Errorf("expected 13, got %d", got)
	}
}

func TestCalculateArmorClass(t *testing.T) {
	// No armor, no shield
	c := newChar()
	c.Class = "fighter"                          // normal class
	if ac := c.CalculateArmorClass(); ac != 12 { // 10 + DexMod 2
		t.Errorf("expected 12, got %d", ac)
	}

	// Armor + full Dex bonus
	c.EquipArmor("leather") // base 11 + dexMod 2
	if ac := c.CalculateArmorClass(); ac != 13 {
		t.Errorf("expected 13, got %d", ac)
	}

	// Shield adds +2
	c.EquipShield("shield")
	if ac := c.CalculateArmorClass(); ac != 15 {
		t.Errorf("expected 15, got %d", ac)
	}

	// Barbarian unarmored defense = 10 + Dex + Con
	b := newChar()
	b.Class = "barbarian"
	b.Stats.ConMod = 3
	if ac := b.CalculateArmorClass(); ac != 15 { // 10 + 2 + 3
		t.Errorf("barbarian expected 15, got %d", ac)
	}

	// Monk unarmored defense = 10 + Dex + Wis
	m := newChar()
	m.Class = "monk"
	m.Stats.WisMod = 2
	if ac := m.CalculateArmorClass(); ac != 14 { // 10 + 2 + 2
		t.Errorf("monk expected 14, got %d", ac)
	}
}
