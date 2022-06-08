// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/service/world/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/esabril/paimoncookies/internal/service/world/model"
	gomock "github.com/golang/mock/gomock"
)

// MockIRepo is a mock of IRepo interface.
type MockIRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRepoMockRecorder
}

// MockIRepoMockRecorder is the mock recorder for MockIRepo.
type MockIRepoMockRecorder struct {
	mock *MockIRepo
}

// NewMockIRepo creates a new mock instance.
func NewMockIRepo(ctrl *gomock.Controller) *MockIRepo {
	mock := &MockIRepo{ctrl: ctrl}
	mock.recorder = &MockIRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepo) EXPECT() *MockIRepoMockRecorder {
	return m.recorder
}

// GetWeekdayTalentBooksWithLocation mocks base method.
func (m *MockIRepo) GetWeekdayTalentBooksWithLocation(weekday string) ([]model.TalentBook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeekdayTalentBooksWithLocation", weekday)
	ret0, _ := ret[0].([]model.TalentBook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeekdayTalentBooksWithLocation indicates an expected call of GetWeekdayTalentBooksWithLocation.
func (mr *MockIRepoMockRecorder) GetWeekdayTalentBooksWithLocation(weekday interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeekdayTalentBooksWithLocation", reflect.TypeOf((*MockIRepo)(nil).GetWeekdayTalentBooksWithLocation), weekday)
}

// GetWeekdayWeaponMaterialsWithLocation mocks base method.
func (m *MockIRepo) GetWeekdayWeaponMaterialsWithLocation(weekday string) ([]model.WeaponMaterial, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeekdayWeaponMaterialsWithLocation", weekday)
	ret0, _ := ret[0].([]model.WeaponMaterial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWeekdayWeaponMaterialsWithLocation indicates an expected call of GetWeekdayWeaponMaterialsWithLocation.
func (mr *MockIRepoMockRecorder) GetWeekdayWeaponMaterialsWithLocation(weekday interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeekdayWeaponMaterialsWithLocation", reflect.TypeOf((*MockIRepo)(nil).GetWeekdayWeaponMaterialsWithLocation), weekday)
}