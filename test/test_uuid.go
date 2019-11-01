package test

import uuid "github.com/satori/go.uuid"

type TestUUIDPort struct{}

func NewTestUUIDPort() *TestUUIDPort {
	return &TestUUIDPort{}
}

func (u *TestUUIDPort) Generate() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
