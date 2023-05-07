package entity

type RajaOngkir struct {
	Id        int     `json:"id"`
	UserId    int     `json:"userId"`
	Title     *string `json:"title"`
	Completed bool    `json:"completed"`
}
