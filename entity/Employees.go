package entity

import (
	"gorm.io/gorm"
	"github.com/asaskevich/govalidator"
)

type Empolyees struct {
	gorm.Model
	Name         string   `valid:"stringlength(2|80)~Name must be 2 and 80 characters"`
	Salary       float64  `valid:"float,range(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode string   `valid:"matches(^(A-Z)-(0-9){4}$)~EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(false)
}

func ValidateEmployees(e *Empolyees) (bool, error) {
	return govalidator.ValidateStruct(e)
}
