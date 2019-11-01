package adapter

import uuid "github.com/satori/go.uuid"

type goid struct{}

func NewUUID() *goid {
	return &goid{}
}

func (u *goid) Generate() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
