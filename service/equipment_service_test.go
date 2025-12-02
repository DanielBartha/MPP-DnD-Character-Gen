package service

import (
	"strings"
	"testing"

	"github.com/DanielBartha/MPP-DnD-Character-Gen/domain"
)

func TestEquipmentService_EquipWeaponSuccess(t *testing.T) {
	svc := NewEquipmentService()
	char := &domain.Character{
		Equipment: domain.Equipment{
			Weapon: map[string]string{},
		},
	}

	msg, err := svc.Equip(char, "shortsword", "main hand", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "Equipped weapon shortsword to main hand"
	if msg != expected {
		t.Errorf("expected %q, got %q", expected, msg)
	}

	if char.Equipment.Weapon["main hand"] != "shortsword" {
		t.Errorf("weapon not actually equipped")
	}
}

func TestEquipmentService_EquipWeaponSlotOccupied(t *testing.T) {
	svc := NewEquipmentService()
	char := &domain.Character{
		Equipment: domain.Equipment{
			Weapon: map[string]string{"off hand": "dagger"},
		},
	}

	_, err := svc.Equip(char, "shortsword", "off hand", "", "")
	if err == nil {
		t.Fatal("expected error but got none")
	}

	expected := "off hand already occupied"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}
}

func TestEquipmentService_EquipArmor(t *testing.T) {
	svc := NewEquipmentService()
	char := &domain.Character{}

	msg, err := svc.Equip(char, "", "", "chain shirt", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "Equipped armor chain shirt"
	if msg != expected {
		t.Errorf("expected %q, got %q", expected, msg)
	}

	if char.Equipment.Armor != "chain shirt" {
		t.Errorf("armor not equipped")
	}
}

func TestEquipmentService_EquipShield(t *testing.T) {
	svc := NewEquipmentService()
	char := &domain.Character{}

	msg, err := svc.Equip(char, "", "", "", "shield")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "Equipped shield shield"
	if msg != expected {
		t.Errorf("expected %q, got %q", expected, msg)
	}

	if char.Equipment.Shield != "shield" {
		t.Errorf("shield not equipped")
	}
}

func TestEquipmentService_NoInput(t *testing.T) {
	svc := NewEquipmentService()
	char := &domain.Character{}

	_, err := svc.Equip(char, "", "", "", "")
	if err == nil {
		t.Fatal("expected error for no input")
	}

	expected := "no equipment provided"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("expected error to contain %q, got %q", expected, err.Error())
	}
}
