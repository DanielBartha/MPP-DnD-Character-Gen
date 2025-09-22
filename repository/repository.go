package repository

import "github.com/DanielBartha/MPP-DnD-Character-Gen/domain"

type CharacterRepository interface {
	Save(character *domain.Character) error
	Load(name string) (*domain.Character, error)
	List() ([]*domain.Character, error)
	Delete(name string) error
}
