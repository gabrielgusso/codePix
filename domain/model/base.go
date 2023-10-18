package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init(){
	govalidator.SetFieldsRequoredByDefault(true)
}

type Base struct {
	ID 		  string 	`json:"id" valid:"id"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}


