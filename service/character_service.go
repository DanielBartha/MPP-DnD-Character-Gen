package service

import (
	"strings"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/characterClasses"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

type CharacterService struct {
	// CharacterSkillProvider in the future for loading data from db
	// c'est possible
}

func NewCharacterService() *CharacterService {
	return &CharacterService{}
}

func (s *CharacterService) GetClassSkills(c *domain.Character) domain.ClassLoadout {
	classKey := strings.ToLower(strings.TrimSpace(c.Class))
	cs, ok := characterClasses.Classes[classKey]

	if !ok {
		return domain.ClassLoadout{
			MaxAllowed: 0,
			Skills:     []string{"insight", "religion"},
			Armor:      []string{},
			Shields:    "",
			Weapons:    []domain.WeaponInfo{},
			MainHand:   "",
			OffHand:    "",
		}
	}

	src := cs.Skills
	localSkills := make([]string, len(src))
	copy(localSkills, src)

	selected := []string{}
	if cs.MaxAllowed > 0 && len(localSkills) > 0 {
		limit := cs.MaxAllowed
		if limit > len(localSkills) {
			limit = len(localSkills)
		}
		selected = localSkills[:limit]
	}

	selected = append(selected, "insight", "religion")

	return domain.ClassLoadout{
		MaxAllowed: cs.MaxAllowed,
		Skills:     selected,
		Armor:      cs.Armor,
		Shields:    cs.Shields,
		Weapons:    cs.Weapons,
		MainHand:   cs.MainHand,
		OffHand:    cs.OffHand,
	}
}
