package rules

import (
	"errors"
	"time"

	"github.com/Mamenzul/sesam-vital-core-utils/domain/facture"
)

// RG_DAT_001 : La date de soin ne doit pas être dans le futur
func CheckDateSoinNotFuture(f *facture.Facture) error {
	now := time.Now()
	for _, p := range f.Prestations {
		if p.DateSoin.After(now) {
			return errors.New("RG_DAT_001: La date de soin ne peut pas être dans le futur")
		}
	}
	return nil
}
