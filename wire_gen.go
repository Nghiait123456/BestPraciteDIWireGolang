// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"google.golang.org/grpc"
	"net"
	"wire_best_practice/a"
	"wire_best_practice/b"
	"wire_best_practice/c"
	"wire_best_practice/repo"
	"wire_best_practice/rpci"
	"wire_best_practice/service"
)

// Injectors from wire.go:

func initApp() (*App, error) {
	listener, err := NewListener()
	if err != nil {
		return nil, err
	}
	server := NewGRPCServer()
	db, err := DBConn()
	if err != nil {
		return nil, err
	}
	repoDB := repo.New(db)
	serviceService := service.New(repoDB)
	rpciServer := rpci.New(serviceService, server)
	app := &App{
		listener: listener,
		gsrv:     server,
		rpcImpl:  rpciServer,
	}
	return app, nil
}

func ProviderB() (b.BInterface, error) {
	a2 := a.NewA2()
	bB := b.NewB(a2)
	return bB, nil
}

func ProviderC() (c.CInterface, error) {
	a2 := a.NewA2()
	bB := b.NewB(a2)
	cC := c.NewC(bB, a2)
	return cC, nil
}

// wire.go:

func NewListener() (net.Listener, error) {
	return net.Listen("tcp4", "0.0.0.0:5000")
}

func NewGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func DBConn() (*sql.DB, error) {
	return sql.Open("mysql", "127.0.0.1:3306")
}
