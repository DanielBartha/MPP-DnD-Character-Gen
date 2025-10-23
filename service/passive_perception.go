package service

import "github.com/DanielBartha/MPP-DnD-Character-Gen/domain"

func CalculatePassivePerception(c *domain.Character) int {
	base := 10 + c.Stats.WisMod

	for _, skill := range c.Skills.Skills {
		if skill == "perception" {
			base += c.Proficiency
			break
		}
	}
	return base
}
