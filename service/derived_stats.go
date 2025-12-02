package service

import (
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func ComputeDerivedStats(c *domain.Character) {
	c.ArmorClass = c.CalculateArmorClass()
	c.InitiativeBonus = c.CalculateInitiative()
	c.PassivePerception = c.CalculatePassivePerception()
}
