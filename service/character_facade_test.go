package service

import (
	"errors"
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain/class"
)

type fakeRepo struct {
	storage map[string]*domain.Character
	saveErr error
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{storage: make(map[string]*domain.Character)}
}

func (r *fakeRepo) Save(c *domain.Character) error {
	if r.saveErr != nil {
		return r.saveErr
	}
	r.storage[c.Name] = c
	return nil
}
func (r *fakeRepo) Load(name string) (*domain.Character, error) {
	if c, ok := r.storage[name]; ok {
		return c, nil
	}
	return nil, errors.New("not found")
}
func (r *fakeRepo) List() ([]*domain.Character, error) {
	result := []*domain.Character{}
	for _, c := range r.storage {
		result = append(result, c)
	}
	return result, nil
}
func (r *fakeRepo) Delete(name string) error {
	delete(r.storage, name)
	return nil
}

func Get(string) (class.ClassSkills, bool) {
	return class.ClassSkills{
		MaxAllowed: 2,
		Skills:     []string{"athletics", "perception"},
		Armor:      []string{},
	}, true
}

func TestFacade_CreateCharacter(t *testing.T) {
	repo := newFakeRepo()
	facade := NewCharacterFacade(repo, &class.ClassRepository{})

	char := &domain.Character{
		Name:  "Test",
		Race:  "dwarf",
		Class: "fighter",
		Level: 1,
	}

	err := facade.CreateCharacter(char)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	saved, _ := repo.Load("Test")
	if saved == nil {
		t.Fatalf("character not saved")
	}

	if saved.Proficiency == 0 {
		t.Error("expected proficiency to be initialized")
	}
	if saved.ArmorClass == 0 {
		t.Error("expected derived stats to be calculated")
	}
}

func TestFacade_ViewCharacter_RecomputesDerivedStats(t *testing.T) {
	repo := newFakeRepo()
	facade := NewCharacterFacade(repo, &class.ClassRepository{})

	char := &domain.Character{
		Name:  "Test",
		Class: "fighter",
		Level: 1,
	}
	repo.Save(char)

	out, err := facade.ViewCharacter("Test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out.ArmorClass == 0 {
		t.Error("expected derived stats to be recalculated on view")
	}
}

func TestFacade_DeleteCharacter(t *testing.T) {
	repo := newFakeRepo()
	facade := NewCharacterFacade(repo, &class.ClassRepository{})

	repo.Save(&domain.Character{Name: "Test"})

	err := facade.DeleteCharacter("Test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := repo.Load("Test"); err == nil {
		t.Error("expected character to be deleted")
	}
}
