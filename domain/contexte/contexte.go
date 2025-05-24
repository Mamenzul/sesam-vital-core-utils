package contexte

type ContexteFacturation struct {
	TypeFlux        string // "FSE", "DRE", etc.
	TiersPayant     bool
	Urgence         bool
	MedecinTraitant bool
	AccidentTravail *AccidentTravail
}

type AccidentTravail struct {
	NumAccident string
	Date        string
}
