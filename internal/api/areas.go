package api

import "projetoinfiel/internal/repositories"

type AreaAPI struct {
	areaRepository repositories.AreaRepository
}

func NewAreaAPI(areaRepository repositories.AreaRepository) *AreaAPI {
	return &AreaAPI{
		areaRepository: areaRepository,
	}
}
