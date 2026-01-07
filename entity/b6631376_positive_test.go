package entity

import (
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Employees is valid`, func(t *testing.T) {
		user := Empolyees {
			Name:         "nichaphat",
			Salary:       30000,
			EmployeeCode: "HR-1024",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
