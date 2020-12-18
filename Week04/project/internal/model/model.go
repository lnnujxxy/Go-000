package model

type Test struct {
	ID   uint32 `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}
