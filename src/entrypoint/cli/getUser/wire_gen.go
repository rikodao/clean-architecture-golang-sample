// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/repository/userRepository"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/application/usecase"
)

// Injectors from wire.go:

func InitializeUserController() (*userController.UserJsonController, error) {
	userInMemoryRepository, err := userRepository.New()
	if err != nil {
		return nil, err
	}
	userJsonPresentator, err := userPresentator.New()
	if err != nil {
		return nil, err
	}
	getUserInteractor, err := usecase.New(userInMemoryRepository, userJsonPresentator)
	if err != nil {
		return nil, err
	}
	userJsonController, err := userController.New(getUserInteractor)
	if err != nil {
		return nil, err
	}
	return userJsonController, nil
}
