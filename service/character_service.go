package service

import (
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain/class"
)

type CharacterService struct {
	classRepo *class.ClassRepository
}

func NewCharacterService(classRepo *class.ClassRepository) *CharacterService {
	return &CharacterService{
		classRepo: classRepo,
	}
}

func (s *CharacterService) GetClassSkills(c *domain.Character) domain.ClassLoadout {
	cs, ok := s.classRepo.GetCS(c.Class)

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

	src := append([]string{}, cs.Skills...)
	selected := []string{}
	if cs.MaxAllowed > 0 && len(src) > 0 {
		limit := cs.MaxAllowed
		if limit > len(src) {
			limit = len(src)
		}
		selected = src[:limit]
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
