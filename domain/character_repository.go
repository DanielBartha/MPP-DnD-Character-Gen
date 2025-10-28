package domain

type CharacterRepository interface {
	Save(character *Character) error
	Load(name string) (*Character, error)
	List() ([]*Character, error)
	Delete(name string) error
}
