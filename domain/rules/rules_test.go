package rules

import (
	"testing"
	"time"

	"github.com/Mamenzul/sesam-vital-core-utils/domain/facture"
	"github.com/Mamenzul/sesam-vital-core-utils/domain/prestation"
)

func TestCheckDateSoinNotFuture(t *testing.T) {
	today := time.Now()
	future := today.Add(24 * time.Hour)

	t.Run("date valide", func(t *testing.T) {
		f := &facture.Facture{
			Prestations: []prestation.Prestation{
				{DateSoin: today},
			},
		}

		if err := CheckDateSoinNotFuture(f); err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("date future", func(t *testing.T) {
		f := &facture.Facture{
			Prestations: []prestation.Prestation{
				{DateSoin: future},
			},
		}

		if err := CheckDateSoinNotFuture(f); err == nil {
			t.Errorf("Expected error for future date, got nil")
		}
	})
}
