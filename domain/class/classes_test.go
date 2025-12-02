package class

import (
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func fakeWeapons() (all, simple, martial []domain.WeaponInfo) {
	all = []domain.WeaponInfo{
		{Name: "dagger"},
		{Name: "longsword"},
		{Name: "hand crossbow"},
	}

	simple = []domain.WeaponInfo{
		{Name: "dagger"},
	}

	martial = []domain.WeaponInfo{
		{Name: "longsword"},
	}

	return
}

func TestClassRepository_Get(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := NewClassRepository(all, simple, martial)

	cs, ok := repo.GetCS(" Bard ")
	if !ok {
		t.Fatalf("expected bard to exist")
	}

	if cs.MaxAllowed != 3 {
		t.Errorf("expected bard MaxAllowed=3, got %d", cs.MaxAllowed)
	}
	if len(cs.Skills) == 0 {
		t.Errorf("expected bard to have skills")
	}
}

func TestClassRepository_findByName(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := NewClassRepository(all, simple, martial)

	w, ok := repo.FindByName("dagger")
	if !ok || w.Name != "dagger" {
		t.Errorf("expected to find dagger")
	}

	w2, ok2 := repo.FindByName("nonexistent")
	if ok2 || w2.Name != "nonexistent" {
		t.Errorf("expected fallback weapon")
	}
}

func (r *ClassRepository) FindByName(name string) (domain.WeaponInfo, bool) {
	return r.findByName(name)
}

func TestClassRepository_WithExtra(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := NewClassRepository(all, simple, martial)

	result := repo.withExtra(simple, "longsword", "unknown")

	if len(result) != 3 {
		t.Fatalf("expected 3 weapons, got %d", len(result))
	}
	if result[1].Name != "longsword" {
		t.Errorf("expected longsword to be added")
	}
	if result[2].Name != "unknown" {
		t.Errorf("expected fallback for unknown weapon")
	}
}

func TestClassRepository_WeaponsByName(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := NewClassRepository(all, simple, martial)

	result := repo.WeaponsByName("dagger", "missing")

	if result[0].Name != "dagger" {
		t.Errorf("expected dagger")
	}
	if result[1].Name != "missing" {
		t.Errorf("expected missing fallback")
	}
}

func TestClassRepository_CombineWeaponSets(t *testing.T) {
	all, simple, martial := fakeWeapons()
	repo := NewClassRepository(all, simple, martial)

	result := repo.CombineWeaponSets(simple, martial)
	if len(result) != 2 {
		t.Errorf("expected 2 weapons, got %d", len(result))
	}
}
