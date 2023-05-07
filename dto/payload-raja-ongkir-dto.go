package dto

type FilterRajaOngkirDto struct {
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

type PageRajaOngkirDto struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

type CreateRajaOngkirDto struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type UpdateRajaOngkirDto struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
