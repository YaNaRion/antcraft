package ressource

type RessourceMap struct {
	Ressources []RessourceMineral
}

func NewRessourceMap(ressources []RessourceMineral) *RessourceMap {
	return &RessourceMap{
		Ressources: ressources,
	}
}

func (rm *RessourceMap) ClearEmptyRessource() {
	newRessources := make([]RessourceMineral, 0)
	for _, ressource := range rm.Ressources {
		if ressource.GetQuantity() > 0 {
			newRessources = append(newRessources, ressource)
		}
	}
	rm.Ressources = newRessources
}
