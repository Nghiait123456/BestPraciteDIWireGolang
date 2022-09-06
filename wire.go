//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"net"
	"wire_best_practice/b"
	"wire_best_practice/c"
	"wire_best_practice/repo"
	"wire_best_practice/rpci"
	"wire_best_practice/service"
)

func NewListener() (net.Listener, error) {
	return net.Listen("tcp4", "0.0.0.0:5000")
}

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func DBConn() (*sql.DB, error) {
	return sql.Open("mysql", "127.0.0.1:3306")
}

func initApp() (*App, error) {
	wire.Build(
		rpci.New,
		NewListener,
		NewGRPCServer,
		repo.Provider,
		DBConn,
		service.Provider,
		wire.Struct(new(App), "*"),
	)

	return &App{}, nil
}

func ProviderB() (b.BInterface, error) {
	wire.Build(
		b.ProviderB,
	)

	return &b.B{}, nil
}

func ProviderC() (c.CInterface, error) {
	wire.Build(
		c.ProviderC,
	)

	return &c.C{}, nil
}

//func initC() (*c.C, error) {
//	wire_best_practice.Build(
//		//a.ProviderA,
//		b.ProviderB,
//		wire_best_practice.Struct(new(c.C), "*"),
//	)
//
//	return &c.C{}, nil
//}
