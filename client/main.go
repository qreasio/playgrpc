package main

import (
	"context"
	"log"
	"time"

	"github.com/qreasio/playgrpc/pkg/hrd"
	"google.golang.org/grpc"
)

const (
	address = "localhost:4040"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hrd.NewHumanResourceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	employee := &hrd.Employee{Firstname: "Bruce", Lastname: "Banner", EducationLevel: 2}
	r, err := c.CreateEmployee(ctx, employee)
	if err != nil {
		log.Fatalf("could not create employee: %v", err)
	}
	log.Printf("Employee ID: %d", r.Id)
	log.Printf("Employee Name: %s %s", r.Employee.Firstname, r.Employee.Lastname)
	log.Printf("Employee Department Name: %s", r.Department.Name)

	salary, err := c.GetSalary(ctx, employee)
	if err != nil {
		log.Fatalf("could not get salary: %v", err)
	}
	log.Printf("Employee Salary: %.2f %s", salary.Amount, salary.Currency)
}
