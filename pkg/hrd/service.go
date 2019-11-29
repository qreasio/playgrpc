package hrd

import (
	"context"
)

type Server struct {
	UnimplementedHumanResourceServer
}

func (s *Server) CreateEmployee(ctx context.Context, employee *Employee) (*EmployeeResponse, error) {
	dept := &Department{Name: "Backend Engineering"}
	resp := &EmployeeResponse{Id: int64(1), Department: dept, Employee: employee}
	return resp, nil
}

func (s *Server) GetSalary(ctx context.Context, e *Employee) (*Salary, error) {
	baseSalaryAmount := 100000.0
	salaryAmount := baseSalaryAmount

	if e.EducationLevel == 2 {
		salaryAmount = baseSalaryAmount * 1.2
	} else if e.EducationLevel == 3 {
		salaryAmount = baseSalaryAmount * 1.5
	} else if e.EducationLevel > 3 {
		salaryAmount = baseSalaryAmount * 2
	}

	return &Salary{Amount: salaryAmount, Currency: "USD"}, nil
}
