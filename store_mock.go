package main

import (
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (mock *MockStore) CreateBird(bird *Bird) error {
	rets := mock.Called(bird)
	return rets.Error(0)
}

func (mock *MockStore) GetBirds() ([]*Bird, error) {
	rets := mock.Called()
	return rets.Get(0).([]*Bird), rets.Error(1)
}

func InitMockstore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
