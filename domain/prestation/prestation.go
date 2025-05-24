package prestation

import "time"

type Prestation struct {
	CodeActe        string
	Libelle         string
	DateSoin        time.Time
	Quantite        int
	TarifUnitaire   float64
	MontantTotal    float64
	CodeExoneration string
}
