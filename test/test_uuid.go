package test

import uuid "github.com/satori/go.uuid"

type UUID struct{}

func (u UUID) Generate() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
