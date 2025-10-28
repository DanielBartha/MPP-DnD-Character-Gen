package domain

import (
	"errors"
	"strings"
)

type Equipment struct {
	Weapon map[string]string
	Armor  string
	Shield string
}

var (
	ErrSlotOccupied = errors.New("slot already occupied")
	ErrInvalidSlot  = errors.New("invalid slot")
	ErrNoWeaponName = errors.New("weapon name cannot be empty")
	ErrNoArmorName  = errors.New("armor name cannot be empty")
	ErrNoShieldName = errors.New("shield name cannot be empty")
)

func (c *Character) EquipWeapon(slot, weapon string) error {
	slot = strings.ToLower(strings.TrimSpace(slot))
	if weapon == "" {
		return ErrNoWeaponName
	}

	if slot != "main hand" && slot != "off hand" {
		return ErrInvalidSlot
	}

	if c.Equipment.Weapon == nil {
		c.Equipment.Weapon = make(map[string]string)
	}

	if existing, ok := c.Equipment.Weapon[slot]; ok && existing != "" {
		return ErrSlotOccupied
	}

	c.Equipment.Weapon[slot] = weapon
	return nil
}

func (c *Character) EquipArmor(armor string) error {
	if armor == "" {
		return ErrNoArmorName
	}

	c.Equipment.Armor = armor
	return nil
}

func (c *Character) EquipShield(shield string) error {
	if shield == "" {
		return ErrNoShieldName
	}

	c.Equipment.Shield = shield
	return nil
}
