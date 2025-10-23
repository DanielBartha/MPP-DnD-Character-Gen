package service

import "github.com/DanielBartha/MPP-DnD-Character-Gen/domain"

func ComputeDerivedStats(c *domain.Character) {
	c.ArmorClass = CalculateArmorClass(c)
	c.InitiativeBonus = CalculateInitiative(&c.Stats)
	c.PassivePerception = CalculatePassivePerception(c)
}
