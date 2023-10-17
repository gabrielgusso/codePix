import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base `valid:"required"`
	Code String `json:"code" valid:"notnull"`
	Name String `json:"name" valid:"notnull"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if err != nil {
		return null, err
	}
	return nil
}
 
func NewBank(code string, name string) (*Bank, error) {
	bank := Bank {
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return null, err
	}

	return &bank, nil
}