package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSalary(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Salary must be between 15000 and 200000`, func(t *testing.T) {
		user := Empolyees {
			Name:         "nichaphat",
			Salary:       1000,
			EmployeeCode: "HR-1024",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
} 


