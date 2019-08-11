// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bft "github.com/SmartBFT-Go/consensus/internal/bft"
import mock "github.com/stretchr/testify/mock"
import smartbftprotos "github.com/SmartBFT-Go/consensus/smartbftprotos"

// LeaderMonitor is an autogenerated mock type for the LeaderMonitor type
type LeaderMonitor struct {
	mock.Mock
}

// ChangeRole provides a mock function with given fields: role, view, leaderID
func (_m *LeaderMonitor) ChangeRole(role bft.Role, view uint64, leaderID uint64) {
	_m.Called(role, view, leaderID)
}

// Close provides a mock function with given fields:
func (_m *LeaderMonitor) Close() {
	_m.Called()
}

// ProcessMsg provides a mock function with given fields: sender, msg
func (_m *LeaderMonitor) ProcessMsg(sender uint64, msg *smartbftprotos.Message) {
	_m.Called(sender, msg)
}
