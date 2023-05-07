package mapping

import "github.com/celmysql-api/entity"

type ImageFileVm struct {
	Oid      string  `json:"oid"`
	FileUrl  *string `json:"fileUrl"`
	FileName *string `json:"fileName"`
}

func ToImageFileResponse(imageFile entity.ImageFile) ImageFileVm {
	return ImageFileVm{
		Oid:      imageFile.Oid,
		FileUrl:  imageFile.FileUrl,
		FileName: imageFile.FileName,
	}
}

func ToImageFileResponses(imageFiles []entity.ImageFile) []ImageFileVm {
	var imageFileResponses []ImageFileVm
	for _, imageFile := range imageFiles {
		imageFileResponses = append(imageFileResponses, ToImageFileResponse(imageFile))
	}
	return imageFileResponses
}
