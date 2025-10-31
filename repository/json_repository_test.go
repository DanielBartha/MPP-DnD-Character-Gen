package repository_test

import (
	"path/filepath"
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
	"github.com/DanielBartha/MPP-DnD-Character-Gen/repository"
)

func tempRepo(t *testing.T) *repository.JsonRepository {
	t.Helper()
	tmpDir := t.TempDir()
	file := filepath.Join(tmpDir, "characters.json")
	return repository.NewJsonRepository(file)
}

func TestJsonRepository_SaveAndLoad(t *testing.T) {
	repo := tempRepo(t)

	char := &domain.Character{Name: "TestHero", Race: "elf", Class: "wizard", Level: 3}
	err := repo.Save(char)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := repo.Load("TestHero")
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.Name != "TestHero" {
		t.Errorf("expected name %q, got %q", "TestHero", loaded.Name)
	}
	if loaded.Class != "wizard" {
		t.Errorf("expected class %q, got %q", "wizard", loaded.Class)
	}
	if loaded.Level != 3 {
		t.Errorf("expected level 3, got %d", loaded.Level)
	}
}

func TestJsonRepository_List(t *testing.T) {
	repo := tempRepo(t)

	repo.Save(&domain.Character{Name: "A"})
	repo.Save(&domain.Character{Name: "B"})

	list, err := repo.List()
	if err != nil {
		t.Fatalf("List error: %v", err)
	}
	if len(list) != 2 {
		t.Errorf("expected 2 characters, got %d", len(list))
	}
}

func TestJsonRepository_Delete(t *testing.T) {
	repo := tempRepo(t)

	repo.Save(&domain.Character{Name: "DeleteMe"})
	repo.Save(&domain.Character{Name: "KeepMe"})

	err := repo.Delete("DeleteMe")
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	list, _ := repo.List()
	if len(list) != 1 {
		t.Fatalf("expected 1 character after delete, got %d", len(list))
	}

	if list[0].Name != "KeepMe" {
		t.Errorf("expected remaining character to be KeepMe, got %s", list[0].Name)
	}
}

func TestJsonRepository_Load_NotFound(t *testing.T) {
	repo := tempRepo(t)

	_, err := repo.Load("missing")
	if err == nil {
		t.Fatalf("expected error when loading missing character")
	}
}
