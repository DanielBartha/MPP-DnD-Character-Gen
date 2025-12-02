package domain

type Repository interface {
	Save(character *Character) error
	Load(name string) (*Character, error)
	List() ([]*Character, error)
	Delete(name string) error
}
