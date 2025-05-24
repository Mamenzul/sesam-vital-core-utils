package facture

import (
	"time"

	"github.com/Mamenzul/sesam-vital-core-utils/domain/beneficiaire"
	"github.com/Mamenzul/sesam-vital-core-utils/domain/contexte"
	"github.com/Mamenzul/sesam-vital-core-utils/domain/prestation"
	"github.com/Mamenzul/sesam-vital-core-utils/domain/ps"
	"github.com/Mamenzul/sesam-vital-core-utils/domain/valorisation"
)

type Facture struct {
	ID           string                       `json:"id"`
	DateEmission time.Time                    `json:"date_emission,omitempty"`
	PS           ps.ProfessionnelDeSante      `json:"ps,omitempty"`
	Beneficiaire beneficiaire.Beneficiaire    `json:"beneficiaire,omitempty"`
	Prestations  []prestation.Prestation      `json:"prestations,omitempty"`
	Contexte     contexte.ContexteFacturation `json:"contexte,omitempty"`
	Valorisation valorisation.Valorisation    `json:"valorisation,omitempty"`
	Statut       StatutFacture                `json:"statut,omitempty"`
}

type StatutFacture string

const (
	StatutInitial StatutFacture = "initial"
	StatutValide  StatutFacture = "valide"
	StatutErreur  StatutFacture = "erreur"
	StatutSoumise StatutFacture = "soumise"
	StatutRejetee StatutFacture = "rejetee"
)
