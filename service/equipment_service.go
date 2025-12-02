package service

import (
	"fmt"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

type EquipmentService struct{}

func NewEquipmentService() *EquipmentService {
	return &EquipmentService{}
}

func (s *EquipmentService) Equip(character *domain.Character, weapon, slot, armor, shield string) (string, error) {
	if character.Equipment.Weapon == nil {
		character.Equipment.Weapon = make(map[string]string)
	}

	switch {
	case weapon != "" && slot != "":
		if existing, ok := character.Equipment.Weapon[slot]; ok && existing != "" {
			return "", fmt.Errorf("%s already occupied", slot)
		}
		character.Equipment.Weapon[slot] = weapon
		return fmt.Sprintf("Equipped weapon %s to %s", weapon, slot), nil

	case armor != "":
		character.Equipment.Armor = armor
		return fmt.Sprintf("Equipped armor %s", armor), nil

	case shield != "":
		character.Equipment.Shield = shield
		return fmt.Sprintf("Equipped shield %s", shield), nil

	default:
		return "", fmt.Errorf("no equipment provided")
	}
}
