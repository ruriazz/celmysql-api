package entity

import "time"

type ImageFile struct {
	Oid                 string    `json:"oid"`
	Members             *string   `json:"members"`
	Employees           *string   `json:"employees"`
	Houses              *string   `json:"houses"`
	FileName            *string   `json:"fileName"`
	FilePath            *string   `json:"filePath"`
	FileMime            *string   `json:"fileMime"`
	FileUrl             *string   `json:"fileUrl"`
	Description         *string   `json:"description"`
	OptimisticLockField int       `json:"optimisticLockField"`
	GCRecord            int       `json:"gCRecord"`
	Deleted             bool      `json:"deleted"`
	UserInserted        *string   `json:"userInserted"`
	InsertedDate        time.Time `json:"insertedDate"`
	LastUserId          *string   `json:"lastUserId"`
	LastUpdate          time.Time `json:"lastUpdate"`
}
