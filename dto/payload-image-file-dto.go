package dto

type FilterImageFileDto struct {
	FileName string `json:"fileName"`
	FileUrl  string `json:"fileUrl"`
}

type PageImageFileDto struct {
	Oid      string `json:"oid"`
	FileName string `json:"fileName"`
	FileUrl  string `json:"fileUrl"`
}

type CreateImageFileDto struct {
	FileName     string `json:"fileName"`
	FileUrl      string `json:"fileUrl"`
	UserInserted string `json:"userInserted"`
}

type UpdateImageFileDto struct {
	Oid        string `json:"oid"`
	FileName   string `json:"fileName"`
	FileUrl    string `json:"fileUrl"`
	LastUserId string `json:"lastUserId"`
}
