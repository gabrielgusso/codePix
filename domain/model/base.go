package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

fun init(){
	govalidator.SetFieldsRequoredByDefault(value:true)
}

type Base struct {
	ID 		  string 	`json:"id" valid:"id"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
package model

