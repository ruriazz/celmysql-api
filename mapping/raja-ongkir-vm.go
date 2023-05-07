package mapping

import "github.com/celmysql-api/entity"

type RajaOngkirVm struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	Title     *string `json:"title"`
	Completed bool    `json:"completed"`
}

func ToRajaOngkirResponse(rayon entity.RajaOngkir) RajaOngkirVm {
	return RajaOngkirVm{
		Id:        rayon.Id,
		UserId:    rayon.UserId,
		Title:     rayon.Title,
		Completed: rayon.Completed,
	}
}

func ToRajaOngkirResponses(rayons []entity.RajaOngkir) []RajaOngkirVm {
	var rayonResponses []RajaOngkirVm
	for _, rayon := range rayons {
		rayonResponses = append(rayonResponses, ToRajaOngkirResponse(rayon))
	}
	return rayonResponses
}
