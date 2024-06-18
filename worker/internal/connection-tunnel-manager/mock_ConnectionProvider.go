// Code generated by mockery. DO NOT EDIT.

package connectiontunnelmanager

import (
	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	mock "github.com/stretchr/testify/mock"

	slog "log/slog"
)

// MockConnectionProvider is an autogenerated mock type for the ConnectionProvider type
type MockConnectionProvider[T interface{}, TConfig interface{}] struct {
	mock.Mock
}

type MockConnectionProvider_Expecter[T interface{}, TConfig interface{}] struct {
	mock *mock.Mock
}

func (_m *MockConnectionProvider[T, TConfig]) EXPECT() *MockConnectionProvider_Expecter[T, TConfig] {
	return &MockConnectionProvider_Expecter[T, TConfig]{mock: &_m.Mock}
}

// CloseClientConnection provides a mock function with given fields: client
func (_m *MockConnectionProvider[T, TConfig]) CloseClientConnection(client T) error {
	ret := _m.Called(client)

	if len(ret) == 0 {
		panic("no return value specified for CloseClientConnection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(T) error); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionProvider_CloseClientConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseClientConnection'
type MockConnectionProvider_CloseClientConnection_Call[T interface{}, TConfig interface{}] struct {
	*mock.Call
}

// CloseClientConnection is a helper method to define mock.On call
//   - client T
func (_e *MockConnectionProvider_Expecter[T, TConfig]) CloseClientConnection(client interface{}) *MockConnectionProvider_CloseClientConnection_Call[T, TConfig] {
	return &MockConnectionProvider_CloseClientConnection_Call[T, TConfig]{Call: _e.mock.On("CloseClientConnection", client)}
}

func (_c *MockConnectionProvider_CloseClientConnection_Call[T, TConfig]) Run(run func(client T)) *MockConnectionProvider_CloseClientConnection_Call[T, TConfig] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(T))
	})
	return _c
}

func (_c *MockConnectionProvider_CloseClientConnection_Call[T, TConfig]) Return(_a0 error) *MockConnectionProvider_CloseClientConnection_Call[T, TConfig] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockConnectionProvider_CloseClientConnection_Call[T, TConfig]) RunAndReturn(run func(T) error) *MockConnectionProvider_CloseClientConnection_Call[T, TConfig] {
	_c.Call.Return(run)
	return _c
}

// GetConnectionClient provides a mock function with given fields: driver, connectionString, opts
func (_m *MockConnectionProvider[T, TConfig]) GetConnectionClient(driver string, connectionString string, opts TConfig) (T, error) {
	ret := _m.Called(driver, connectionString, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionClient")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, TConfig) (T, error)); ok {
		return rf(driver, connectionString, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, TConfig) T); ok {
		r0 = rf(driver, connectionString, opts)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(string, string, TConfig) error); ok {
		r1 = rf(driver, connectionString, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionProvider_GetConnectionClient_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionClient'
type MockConnectionProvider_GetConnectionClient_Call[T interface{}, TConfig interface{}] struct {
	*mock.Call
}

// GetConnectionClient is a helper method to define mock.On call
//   - driver string
//   - connectionString string
//   - opts TConfig
func (_e *MockConnectionProvider_Expecter[T, TConfig]) GetConnectionClient(driver interface{}, connectionString interface{}, opts interface{}) *MockConnectionProvider_GetConnectionClient_Call[T, TConfig] {
	return &MockConnectionProvider_GetConnectionClient_Call[T, TConfig]{Call: _e.mock.On("GetConnectionClient", driver, connectionString, opts)}
}

func (_c *MockConnectionProvider_GetConnectionClient_Call[T, TConfig]) Run(run func(driver string, connectionString string, opts TConfig)) *MockConnectionProvider_GetConnectionClient_Call[T, TConfig] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(TConfig))
	})
	return _c
}

func (_c *MockConnectionProvider_GetConnectionClient_Call[T, TConfig]) Return(_a0 T, _a1 error) *MockConnectionProvider_GetConnectionClient_Call[T, TConfig] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionProvider_GetConnectionClient_Call[T, TConfig]) RunAndReturn(run func(string, string, TConfig) (T, error)) *MockConnectionProvider_GetConnectionClient_Call[T, TConfig] {
	_c.Call.Return(run)
	return _c
}

// GetConnectionClientConfig provides a mock function with given fields: connectionConfig
func (_m *MockConnectionProvider[T, TConfig]) GetConnectionClientConfig(connectionConfig *mgmtv1alpha1.ConnectionConfig) (TConfig, error) {
	ret := _m.Called(connectionConfig)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionClientConfig")
	}

	var r0 TConfig
	var r1 error
	if rf, ok := ret.Get(0).(func(*mgmtv1alpha1.ConnectionConfig) (TConfig, error)); ok {
		return rf(connectionConfig)
	}
	if rf, ok := ret.Get(0).(func(*mgmtv1alpha1.ConnectionConfig) TConfig); ok {
		r0 = rf(connectionConfig)
	} else {
		r0 = ret.Get(0).(TConfig)
	}

	if rf, ok := ret.Get(1).(func(*mgmtv1alpha1.ConnectionConfig) error); ok {
		r1 = rf(connectionConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionProvider_GetConnectionClientConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionClientConfig'
type MockConnectionProvider_GetConnectionClientConfig_Call[T interface{}, TConfig interface{}] struct {
	*mock.Call
}

// GetConnectionClientConfig is a helper method to define mock.On call
//   - connectionConfig *mgmtv1alpha1.ConnectionConfig
func (_e *MockConnectionProvider_Expecter[T, TConfig]) GetConnectionClientConfig(connectionConfig interface{}) *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig] {
	return &MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig]{Call: _e.mock.On("GetConnectionClientConfig", connectionConfig)}
}

func (_c *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig]) Run(run func(connectionConfig *mgmtv1alpha1.ConnectionConfig)) *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*mgmtv1alpha1.ConnectionConfig))
	})
	return _c
}

func (_c *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig]) Return(_a0 TConfig, _a1 error) *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig]) RunAndReturn(run func(*mgmtv1alpha1.ConnectionConfig) (TConfig, error)) *MockConnectionProvider_GetConnectionClientConfig_Call[T, TConfig] {
	_c.Call.Return(run)
	return _c
}

// GetConnectionDetails provides a mock function with given fields: connectionConfig, connectionTimeout, logger
func (_m *MockConnectionProvider[T, TConfig]) GetConnectionDetails(connectionConfig *mgmtv1alpha1.ConnectionConfig, connectionTimeout *uint32, logger *slog.Logger) (ConnectionDetails, error) {
	ret := _m.Called(connectionConfig, connectionTimeout, logger)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionDetails")
	}

	var r0 ConnectionDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(*mgmtv1alpha1.ConnectionConfig, *uint32, *slog.Logger) (ConnectionDetails, error)); ok {
		return rf(connectionConfig, connectionTimeout, logger)
	}
	if rf, ok := ret.Get(0).(func(*mgmtv1alpha1.ConnectionConfig, *uint32, *slog.Logger) ConnectionDetails); ok {
		r0 = rf(connectionConfig, connectionTimeout, logger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ConnectionDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(*mgmtv1alpha1.ConnectionConfig, *uint32, *slog.Logger) error); ok {
		r1 = rf(connectionConfig, connectionTimeout, logger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionProvider_GetConnectionDetails_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionDetails'
type MockConnectionProvider_GetConnectionDetails_Call[T interface{}, TConfig interface{}] struct {
	*mock.Call
}

// GetConnectionDetails is a helper method to define mock.On call
//   - connectionConfig *mgmtv1alpha1.ConnectionConfig
//   - connectionTimeout *uint32
//   - logger *slog.Logger
func (_e *MockConnectionProvider_Expecter[T, TConfig]) GetConnectionDetails(connectionConfig interface{}, connectionTimeout interface{}, logger interface{}) *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig] {
	return &MockConnectionProvider_GetConnectionDetails_Call[T, TConfig]{Call: _e.mock.On("GetConnectionDetails", connectionConfig, connectionTimeout, logger)}
}

func (_c *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig]) Run(run func(connectionConfig *mgmtv1alpha1.ConnectionConfig, connectionTimeout *uint32, logger *slog.Logger)) *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*mgmtv1alpha1.ConnectionConfig), args[1].(*uint32), args[2].(*slog.Logger))
	})
	return _c
}

func (_c *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig]) Return(_a0 ConnectionDetails, _a1 error) *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig]) RunAndReturn(run func(*mgmtv1alpha1.ConnectionConfig, *uint32, *slog.Logger) (ConnectionDetails, error)) *MockConnectionProvider_GetConnectionDetails_Call[T, TConfig] {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectionProvider creates a new instance of MockConnectionProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectionProvider[T interface{}, TConfig interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectionProvider[T, TConfig] {
	mock := &MockConnectionProvider[T, TConfig]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}