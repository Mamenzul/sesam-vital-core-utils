package beneficiaire

import "time"

type Beneficiaire struct {
	NumeroVitale            string
	Nom                     string
	Prenom                  string
	DateNaissance           time.Time
	CodeOrganismeAMO        string
	CodeOrganismeAMC        string
	EstCMUC                 bool
	SituationsParticulieres []string // ex: ["SP03", "SP10"]
}
