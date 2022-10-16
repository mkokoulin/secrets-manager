// Package mock_handlers is a generated GoMock package.
package handlers

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "github.com/mkokoulin/secrets-manager.git/server/internal/auth"
	models "github.com/mkokoulin/secrets-manager.git/server/internal/models"
)

// MockUserServiceInterface is a mock of UserServiceInterface interface.
type MockUserServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceInterfaceMockRecorder
}

// MockUserServiceInterfaceMockRecorder is the mock recorder for MockUserServiceInterface.
type MockUserServiceInterfaceMockRecorder struct {
	mock *MockUserServiceInterface
}

// NewMockUserServiceInterface creates a new mock instance.
func NewMockUserServiceInterface(ctrl *gomock.Controller) *MockUserServiceInterface {
	mock := &MockUserServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceInterface) EXPECT() *MockUserServiceInterfaceMockRecorder {
	return m.recorder
}

// AuthUser mocks base method.
func (m *MockUserServiceInterface) AuthUser(ctx context.Context, user models.User) (*auth.TokenDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUser", ctx, user)
	ret0, _ := ret[0].(*auth.TokenDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthUser indicates an expected call of AuthUser.
func (mr *MockUserServiceInterfaceMockRecorder) AuthUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUser", reflect.TypeOf((*MockUserServiceInterface)(nil).AuthUser), ctx, user)
}

// CreateUser mocks base method.
func (m *MockUserServiceInterface) CreateUser(ctx context.Context, user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceInterfaceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserServiceInterface)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockUserServiceInterface) DeleteUser(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceInterfaceMockRecorder) DeleteUser(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserServiceInterface)(nil).DeleteUser), ctx, userID)
}

// RefreshToken mocks base method.
func (m *MockUserServiceInterface) RefreshToken(ctx context.Context, refreshToken string) (*auth.TokenDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", ctx, refreshToken)
	ret0, _ := ret[0].(*auth.TokenDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockUserServiceInterfaceMockRecorder) RefreshToken(ctx, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockUserServiceInterface)(nil).RefreshToken), ctx, refreshToken)
}

// MockSecretServiceInterface is a mock of SecretServiceInterface interface.
type MockSecretServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSecretServiceInterfaceMockRecorder
}

// MockSecretServiceInterfaceMockRecorder is the mock recorder for MockSecretServiceInterface.
type MockSecretServiceInterfaceMockRecorder struct {
	mock *MockSecretServiceInterface
}

// NewMockSecretServiceInterface creates a new mock instance.
func NewMockSecretServiceInterface(ctrl *gomock.Controller) *MockSecretServiceInterface {
	mock := &MockSecretServiceInterface{ctrl: ctrl}
	mock.recorder = &MockSecretServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretServiceInterface) EXPECT() *MockSecretServiceInterfaceMockRecorder {
	return m.recorder
}

// AddSecret mocks base method.
func (m *MockSecretServiceInterface) AddSecret(ctx context.Context, secret models.RawSecretData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecret", ctx, secret)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSecret indicates an expected call of AddSecret.
func (mr *MockSecretServiceInterfaceMockRecorder) AddSecret(ctx, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecret", reflect.TypeOf((*MockSecretServiceInterface)(nil).AddSecret), ctx, secret)
}

// DeleteSecret mocks base method.
func (m *MockSecretServiceInterface) DeleteSecret(ctx context.Context, secretID, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", ctx, secretID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockSecretServiceInterfaceMockRecorder) DeleteSecret(ctx, secretID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockSecretServiceInterface)(nil).DeleteSecret), ctx, secretID, userID)
}

// GetSecret mocks base method.
func (m *MockSecretServiceInterface) GetSecret(ctx context.Context, secretID, userID string) (models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", ctx, secretID, userID)
	ret0, _ := ret[0].(models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}
