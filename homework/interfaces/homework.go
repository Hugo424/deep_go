package main

import (
	"errors"
)

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	services map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	// зарегистрировать конструктор по созданию типа
	if constructor == nil {
		return
	}

	c.services[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, ok := c.services[name]
	if !ok {
		return nil, errors.New("constructor not found")
	}

	constructorFunc, ok := constructor.(func() interface{})
	if !ok {
		return nil, errors.New("invalid constructor")
	}

	return constructorFunc(), nil
}
