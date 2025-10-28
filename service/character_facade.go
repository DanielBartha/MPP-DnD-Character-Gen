package service

import (
	"fmt"

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

func (f *CharacterFacade) ListCharacters() ([]*domain.Character, error) {
	chars, err := f.repo.List()
	if err != nil {
		return nil, err
	}

	return chars, nil
}

func (f *CharacterFacade) DeleteCharacter(name string) error {
	return f.repo.Delete(name)
}

func (f *CharacterFacade) EquipItem(name, weapon, slot, armor, shield string) (string, error) {
	char, err := f.repo.Load(name)
	if err != nil {
		return "", fmt.Errorf("character %q not found", name)
	}

	equipService := NewEquipmentService()
	message, err := equipService.Equip(char, weapon, slot, armor, shield)
	if err != nil {
		return "", err
	}

	if err := f.repo.Save(char); err != nil {
		return "", fmt.Errorf("error saving character: %v", err)
	}

	return message, nil
}

func (f *CharacterFacade) LearnSpell(name, spell string) (string, error) {
	char, err := f.repo.Load(name)
	if err != nil {
		return "", fmt.Errorf("character %q not found", name)
	}

	spellService := NewSpellService()
	message, err := spellService.LearnSpell(char, spell)
	if err != nil {
		return "", err
	}

	if err := f.repo.Save(char); err != nil {
		return "", fmt.Errorf("error saving character: %v", err)
	}

	return message, nil
}

func (f *CharacterFacade) PrepareSpell(name, spell string) (string, error) {
	char, err := f.repo.Load(name)
	if err != nil {
		return "", fmt.Errorf("character %q not found", name)
	}

	spellService := NewSpellService()
	message, err := spellService.PrepareSpell(char, spell)
	if err != nil {
		return "", err
	}

	if err := f.repo.Save(char); err != nil {
		return "", fmt.Errorf("error saving character: %v", err)
	}

	return message, nil
}

// hook up for racial bonuses assignment for later
// if fn := f.svc.ApplyRacialBonusesSkillProficiencies; fn != nil {
// 	fn(c)
// }
