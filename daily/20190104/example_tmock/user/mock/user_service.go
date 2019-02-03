// Code generated by mocker. DO NOT EDIT.
// github.com/travisjeffery/mocker
// Source: user_service.go

package mock

import (
	sync "sync"

	x "."
)

// MockUserService is a mock of UserService interface
type MockUserService struct {
	lockGetUser sync.Mutex
	GetUserFunc func(id string) (*x.User, error)

	calls struct {
		GetUser []struct {
			Id string
		}
	}
}

// GetUser mocks base method by wrapping the associated func.
func (m *MockUserService) GetUser(id string) (*x.User, error) {
	m.lockGetUser.Lock()
	defer m.lockGetUser.Unlock()

	if m.GetUserFunc == nil {
		panic("mocker: MockUserService.GetUserFunc is nil but MockUserService.GetUser was called.")
	}

	call := struct {
		Id string
	}{
		Id: id,
	}

	m.calls.GetUser = append(m.calls.GetUser, call)

	return m.GetUserFunc(id)
}

// GetUserCalled returns true if GetUser was called at least once.
func (m *MockUserService) GetUserCalled() bool {
	m.lockGetUser.Lock()
	defer m.lockGetUser.Unlock()

	return len(m.calls.GetUser) > 0
}

// GetUserCalls returns the calls made to GetUser.
func (m *MockUserService) GetUserCalls() []struct {
	Id string
} {
	m.lockGetUser.Lock()
	defer m.lockGetUser.Unlock()

	return m.calls.GetUser
}

// Reset resets the calls made to the mocked methods.
func (m *MockUserService) Reset() {
	m.lockGetUser.Lock()
	m.calls.GetUser = nil
	m.lockGetUser.Unlock()
}