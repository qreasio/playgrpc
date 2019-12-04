#PlaygRPC

This repo is used for personal learning and playing around with gRPC using Go.
There are two parts, gRPC client and gRPC server.

To see this works, we need to run gRPC server first then run gRPC client.

## Run gRPC Server

```sh
go run server/main.go
```

## Run gRPC Client


```sh
go run client/main.go
```

Output:
```sh
2019/11/29 16:21:51 Employee ID: 1
2019/11/29 16:21:51 Employee Name: Bruce Banner
2019/11/29 16:21:51 Employee Department Name: Backend Engineering
2019/11/29 16:21:51 Employee Salary: 1200.00 USD
```