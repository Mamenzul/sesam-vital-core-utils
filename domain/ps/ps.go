package ps

type ProfessionnelDeSante struct {
	ID             string
	Nom            string
	Prenom         string
	SpecialiteCode string
	IdentifiantCPS string
	ModeExercice   string // ex : "libéral", "salarié"
}
