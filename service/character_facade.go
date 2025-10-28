package service

import (
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/repository"
)

type CharacterFacade struct {
	repo repository.CharacterRepository
	svc  *CharacterService
}

func NewCharacterFacade(repo repository.CharacterRepository) *CharacterFacade {
	return &CharacterFacade{
		repo: repo,
		svc:  NewCharacterService(),
	}
}

func (f *CharacterFacade) CreateCharacter(c *domain.Character) error {
	c.Skills = f.svc.GetClassSkills(c)
	f.svc.ApplyRacialBonuses(c)
	f.svc.UpdateProficiency(c)
	f.svc.InitSpellcasting(c)
	ComputeDerivedStats(c)
	return f.repo.Save(c)
}

func (f *CharacterFacade) ViewCharacter(name string) (*domain.Character, error) {
	char, err := f.repo.Load(name)
	if err != nil {
		return nil, err
	}

	ComputeDerivedStats(char)
	return char, nil
}

// hook up for racial bonuses assignment for later
// if fn := f.svc.ApplyRacialBonusesSkillProficiencies; fn != nil {
// 	fn(c)
// }
