package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
type OpeningResponse struct {
	ID        uint64 `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	Remote    bool   `json:"remote"`
	Link      string `json:"link"`
	Salary    int64  `json:"salary"`
}
