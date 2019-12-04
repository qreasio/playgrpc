package test

import (
	"context"
	"github.com/qreasio/playgrpc/pkg/hrd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()
	hrdServer := &hrd.Server{}
	hrd.RegisterHumanResourceServer(grpcServer, hrdServer)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestEmployee(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	client := hrd.NewHumanResourceClient(conn)

	// set up test cases
	testData := []struct {
		firstName      string
		lastName       string
		educationLevel int32
		wantSalary     float64
	}{
		{
			firstName:      "Steve",
			lastName:       "Rogers",
			educationLevel: 1,
			wantSalary:     1000.0,
		},
		{
			firstName:      "Tony",
			lastName:       "Stark",
			educationLevel: 2,
			wantSalary:     1200.0,
		},
		{
			firstName:      "Miles",
			lastName:       "Morales",
			educationLevel: 3,
			wantSalary:     1500.0,
		},
		{
			firstName:      "Bruce",
			lastName:       "Banner",
			educationLevel: 4,
			wantSalary:     2000.0,
		},
	}

	for _, tt := range testData {
		employee := &hrd.Employee{Firstname: tt.firstName, Lastname: tt.lastName, EducationLevel: tt.educationLevel}

		employeeResponse, err := client.CreateEmployee(ctx, employee)
		if err != nil {
			t.Errorf("could not create employee: %v", err)
		}
		if employeeResponse.Id == 0 {
			t.Error("invalid employee id")
		}
		if employeeResponse.Employee == nil {
			t.Error("failed to create employee")
		}
		if employeeResponse.Employee.Firstname != tt.firstName {
			t.Error("invalid employee firstname")
		}
		if employeeResponse.Employee.Lastname != tt.lastName {
			t.Error("invalid employee lastname")
		}
		if employeeResponse.Employee.Lastname != tt.lastName {
			t.Error("invalid employee lastname")
		}
		if employeeResponse.Department == nil {
			t.Error("invalid employee department")
		}
		if employeeResponse.Department.Name == "" {
			t.Error("invalid employee department name")
		}
		s, err := client.GetSalary(ctx, employee)
		if err != nil {
			t.Errorf("could not get salary: %v", err)
		}
		if s.GetCurrency() != "USD" {
			t.Error("invalid currency")
		}
		if s.GetAmount() != tt.wantSalary {
			t.Errorf("Salary :%v, wanted %v", s.Amount, tt.wantSalary)
		}

	}

}
