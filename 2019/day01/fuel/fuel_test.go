package fuel

import (
	"testing"
)

func TestCalculateFuelRequired(t *testing.T) {
	tests := map[string]struct {
		mass     int64
		wantFuel int64
	}{
		"simple":        {mass: 12, wantFuel: 2},
		"rounding down": {mass: 14, wantFuel: 2},
		"other number":  {mass: 1969, wantFuel: 654},
		"large number":  {mass: 100756, wantFuel: 33583},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CalculateFuelRequired(tc.mass)
			if got != tc.wantFuel {
				t.Fatalf("got %d, but should have: %d", got, tc.wantFuel)
			}
		})
	}
}

func TestCalculateFuelRequiredWithFuelForFuel(t *testing.T) {
	tests := map[string]struct {
		moduleMass int64
		wantFuel   int64
	}{
		"simple without fuel for fuel":  {moduleMass: 12, wantFuel: 2},
		"fuel for fuel":                 {moduleMass: 1969, wantFuel: 966},
		"large number of fuel for fuel": {moduleMass: 100756, wantFuel: 50346},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CalculateFuelRequiredWithFuelForFuel(tc.moduleMass)
			if got != tc.wantFuel {
				t.Fatalf("got %d, but should have: %d", got, tc.wantFuel)
			}
		})
	}
}
