// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/characters/repository/repository.go

// Package characters_repo is a generated GoMock package.
package characters_repo

import (
	reflect "reflect"

	model "github.com/esabril/paimoncookies/internal/service/characters/model"
	gomock "github.com/golang/mock/gomock"
)

// MockICharactersRepo is a mock of ICharactersRepo interface.
type MockICharactersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockICharactersRepoMockRecorder
}

// MockICharactersRepoMockRecorder is the mock recorder for MockICharactersRepo.
type MockICharactersRepoMockRecorder struct {
	mock *MockICharactersRepo
}

// NewMockICharactersRepo creates a new mock instance.
func NewMockICharactersRepo(ctrl *gomock.Controller) *MockICharactersRepo {
	mock := &MockICharactersRepo{ctrl: ctrl}
	mock.recorder = &MockICharactersRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICharactersRepo) EXPECT() *MockICharactersRepoMockRecorder {
	return m.recorder
}

// GetCharactersList mocks base method.
func (m *MockICharactersRepo) GetCharactersList() ([]model.Character, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharactersList")
	ret0, _ := ret[0].([]model.Character)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCharactersList indicates an expected call of GetCharactersList.
func (mr *MockICharactersRepoMockRecorder) GetCharactersList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharactersList", reflect.TypeOf((*MockICharactersRepo)(nil).GetCharactersList))
}
