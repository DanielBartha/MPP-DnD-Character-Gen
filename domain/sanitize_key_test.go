package domain

import "testing"

func TestSanitizeLocalKey(t *testing.T) {
	cases := map[string]string{
		"Leather Armor":   "leather",
		"Chain Shirt":     "chain-shirt",
		"Scale Mail":      "scale-mail",
		"Padded Armor":    "padded",
		"Half Plate":      "half-plate",
		"Studded Leather": "studded-leather",
		"Ring Mail":       "ring-mail",
		"chain-mail":      "chain-mail",
		"Plate Armor":     "plate",
		"Hide":            "hide",
		"Shield":          "shield",
		"Mace (Silvered)": "mace-silvered",
		"breast plate":    "breast-plate",
	}

	for input, want := range cases {
		got := SanitizeLocalKey(input)
		if got != want {
			t.Errorf("SanitizeLocalKey(%q) = %q, want %q", input, got, want)
		}
	}
}
