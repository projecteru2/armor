// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/projecteru2/barrel/types"
	mock "github.com/stretchr/testify/mock"
)

// FixedIPAllocator is an autogenerated mock type for the FixedIPAllocator type
type FixedIPAllocator struct {
	mock.Mock
}

// AllocFixedIP provides a mock function with given fields: _a0, _a1
func (_m *FixedIPAllocator) AllocFixedIP(_a0 context.Context, _a1 types.IP) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IP) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AllocFixedIPFromPools provides a mock function with given fields: ctx, pools
func (_m *FixedIPAllocator) AllocFixedIPFromPools(ctx context.Context, pools []types.Pool) (types.IPAddress, error) {
	ret := _m.Called(ctx, pools)

	var r0 types.IPAddress
	if rf, ok := ret.Get(0).(func(context.Context, []types.Pool) types.IPAddress); ok {
		r0 = rf(ctx, pools)
	} else {
		r0 = ret.Get(0).(types.IPAddress)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []types.Pool) error); ok {
		r1 = rf(ctx, pools)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocIP provides a mock function with given fields: ctx, ip
func (_m *FixedIPAllocator) AllocIP(ctx context.Context, ip types.IP) error {
	ret := _m.Called(ctx, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IP) error); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AllocIPFromPool provides a mock function with given fields: ctx, poolID
func (_m *FixedIPAllocator) AllocIPFromPool(ctx context.Context, poolID string) (types.IPAddress, error) {
	ret := _m.Called(ctx, poolID)

	var r0 types.IPAddress
	if rf, ok := ret.Get(0).(func(context.Context, string) types.IPAddress); ok {
		r0 = rf(ctx, poolID)
	} else {
		r0 = ret.Get(0).(types.IPAddress)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocIPFromPools provides a mock function with given fields: ctx, pools
func (_m *FixedIPAllocator) AllocIPFromPools(ctx context.Context, pools []types.Pool) (types.IPAddress, error) {
	ret := _m.Called(ctx, pools)

	var r0 types.IPAddress
	if rf, ok := ret.Get(0).(func(context.Context, []types.Pool) types.IPAddress); ok {
		r0 = rf(ctx, pools)
	} else {
		r0 = ret.Get(0).(types.IPAddress)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []types.Pool) error); ok {
		r1 = rf(ctx, pools)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssignFixedIP provides a mock function with given fields: _a0, _a1
func (_m *FixedIPAllocator) AssignFixedIP(_a0 context.Context, _a1 types.IP) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IP) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDefaultPool provides a mock function with given fields: ipv6
func (_m *FixedIPAllocator) GetDefaultPool(ipv6 bool) types.Pool {
	ret := _m.Called(ipv6)

	var r0 types.Pool
	if rf, ok := ret.Get(0).(func(bool) types.Pool); ok {
		r0 = rf(ipv6)
	} else {
		r0 = ret.Get(0).(types.Pool)
	}

	return r0
}

// GetPoolByCidr provides a mock function with given fields: ctx, cidr
func (_m *FixedIPAllocator) GetPoolByCidr(ctx context.Context, cidr string) (types.Pool, error) {
	ret := _m.Called(ctx, cidr)

	var r0 types.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) types.Pool); ok {
		r0 = rf(ctx, cidr)
	} else {
		r0 = ret.Get(0).(types.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, cidr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoolByID provides a mock function with given fields: ctx, poolID
func (_m *FixedIPAllocator) GetPoolByID(ctx context.Context, poolID string) (types.Pool, error) {
	ret := _m.Called(ctx, poolID)

	var r0 types.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) types.Pool); ok {
		r0 = rf(ctx, poolID)
	} else {
		r0 = ret.Get(0).(types.Pool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, poolID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoolsByCidrs provides a mock function with given fields: ctx, cidr
func (_m *FixedIPAllocator) GetPoolsByCidrs(ctx context.Context, cidr []string) ([]types.Pool, error) {
	ret := _m.Called(ctx, cidr)

	var r0 []types.Pool
	if rf, ok := ret.Get(0).(func(context.Context, []string) []types.Pool); ok {
		r0 = rf(ctx, cidr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, cidr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPoolsByNetworkName provides a mock function with given fields: ctx, name
func (_m *FixedIPAllocator) GetPoolsByNetworkName(ctx context.Context, name string) ([]types.Pool, error) {
	ret := _m.Called(ctx, name)

	var r0 []types.Pool
	if rf, ok := ret.Get(0).(func(context.Context, string) []types.Pool); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Pool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnallocFixedIP provides a mock function with given fields: _a0, _a1
func (_m *FixedIPAllocator) UnallocFixedIP(_a0 context.Context, _a1 types.IP) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IP) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnallocIP provides a mock function with given fields: ctx, ip
func (_m *FixedIPAllocator) UnallocIP(ctx context.Context, ip types.IP) error {
	ret := _m.Called(ctx, ip)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IP) error); ok {
		r0 = rf(ctx, ip)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnassignFixedIP provides a mock function with given fields: _a0, _a1
func (_m *FixedIPAllocator) UnassignFixedIP(_a0 context.Context, _a1 types.IP) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.IP) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
