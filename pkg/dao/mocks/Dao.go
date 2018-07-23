// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import bundle "github.com/automationbroker/bundle-lib/bundle"

import mock "github.com/stretchr/testify/mock"

// Dao is an autogenerated mock type for the Dao type
type Dao struct {
	mock.Mock
}

// BatchDeleteSpecs provides a mock function with given fields: _a0
func (_m *Dao) BatchDeleteSpecs(_a0 []*bundle.Spec) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*bundle.Spec) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchGetBundleInstances provides a mock function with given fields:
func (_m *Dao) BatchGetBundleInstances() ([]*bundle.ServiceInstance, error) {
	ret := _m.Called()

	var r0 []*bundle.ServiceInstance
	if rf, ok := ret.Get(0).(func() []*bundle.ServiceInstance); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*bundle.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchGetSpecs provides a mock function with given fields: _a0
func (_m *Dao) BatchGetSpecs(_a0 string) ([]*bundle.Spec, error) {
	ret := _m.Called(_a0)

	var r0 []*bundle.Spec
	if rf, ok := ret.Get(0).(func(string) []*bundle.Spec); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*bundle.Spec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BatchSetSpecs provides a mock function with given fields: _a0
func (_m *Dao) BatchSetSpecs(_a0 bundle.SpecManifest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(bundle.SpecManifest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBindInstance provides a mock function with given fields: _a0
func (_m *Dao) DeleteBindInstance(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBinding provides a mock function with given fields: _a0, _a1
func (_m *Dao) DeleteBinding(_a0 bundle.BindInstance, _a1 bundle.ServiceInstance) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(bundle.BindInstance, bundle.ServiceInstance) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteServiceInstance provides a mock function with given fields: _a0
func (_m *Dao) DeleteServiceInstance(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSpec provides a mock function with given fields: _a0
func (_m *Dao) DeleteSpec(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindJobStateByState provides a mock function with given fields: _a0
func (_m *Dao) FindJobStateByState(_a0 bundle.State) ([]bundle.RecoverStatus, error) {
	ret := _m.Called(_a0)

	var r0 []bundle.RecoverStatus
	if rf, ok := ret.Get(0).(func(bundle.State) []bundle.RecoverStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bundle.RecoverStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bundle.State) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBindInstance provides a mock function with given fields: _a0
func (_m *Dao) GetBindInstance(_a0 string) (*bundle.BindInstance, error) {
	ret := _m.Called(_a0)

	var r0 *bundle.BindInstance
	if rf, ok := ret.Get(0).(func(string) *bundle.BindInstance); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bundle.BindInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceInstance provides a mock function with given fields: _a0
func (_m *Dao) GetServiceInstance(_a0 string) (*bundle.ServiceInstance, error) {
	ret := _m.Called(_a0)

	var r0 *bundle.ServiceInstance
	if rf, ok := ret.Get(0).(func(string) *bundle.ServiceInstance); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bundle.ServiceInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSpec provides a mock function with given fields: _a0
func (_m *Dao) GetSpec(_a0 string) (*bundle.Spec, error) {
	ret := _m.Called(_a0)

	var r0 *bundle.Spec
	if rf, ok := ret.Get(0).(func(string) *bundle.Spec); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bundle.Spec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: _a0, _a1
func (_m *Dao) GetState(_a0 string, _a1 string) (bundle.JobState, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bundle.JobState
	if rf, ok := ret.Get(0).(func(string, string) bundle.JobState); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bundle.JobState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStateByKey provides a mock function with given fields: key
func (_m *Dao) GetStateByKey(key string) (bundle.JobState, error) {
	ret := _m.Called(key)

	var r0 bundle.JobState
	if rf, ok := ret.Get(0).(func(string) bundle.JobState); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bundle.JobState)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSvcInstJobsByState provides a mock function with given fields: _a0, _a1
func (_m *Dao) GetSvcInstJobsByState(_a0 string, _a1 bundle.State) ([]bundle.JobState, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []bundle.JobState
	if rf, ok := ret.Get(0).(func(string, bundle.State) []bundle.JobState); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bundle.JobState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bundle.State) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsNotFoundError provides a mock function with given fields: err
func (_m *Dao) IsNotFoundError(err error) bool {
	ret := _m.Called(err)

	var r0 bool
	if rf, ok := ret.Get(0).(func(error) bool); ok {
		r0 = rf(err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetBindInstance provides a mock function with given fields: _a0, _a1
func (_m *Dao) SetBindInstance(_a0 string, _a1 *bundle.BindInstance) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *bundle.BindInstance) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetServiceInstance provides a mock function with given fields: _a0, _a1
func (_m *Dao) SetServiceInstance(_a0 string, _a1 *bundle.ServiceInstance) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *bundle.ServiceInstance) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSpec provides a mock function with given fields: _a0, _a1
func (_m *Dao) SetSpec(_a0 string, _a1 *bundle.Spec) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *bundle.Spec) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetState provides a mock function with given fields: _a0, _a1
func (_m *Dao) SetState(_a0 string, _a1 bundle.JobState) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, bundle.JobState) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bundle.JobState) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
