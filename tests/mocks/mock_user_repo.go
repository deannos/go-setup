package mocks

import (
	"sync"
	"time"

	"go-setup/internal/entity"
	"go-setup/pkg/errors"
)

// MockUserRepository implements the UserRepository interface for testing
type MockUserRepository struct {
    users map[int64]*entity.User
    mutex sync.RWMutex
    nextID int64
}

func NewMockUserRepository() *MockUserRepository {
    return &MockUserRepository{
        users:  make(map[int64]*entity.User),
        nextID: 1,
    }
}

func (m *MockUserRepository) Create(user *entity.User) (int64, error) {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    user.ID = m.nextID
    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()
    
    m.users[user.ID] = user
    m.nextID++
    
    return user.ID, nil
}

func (m *MockUserRepository) GetByID(id int64) (*entity.User, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    user, exists := m.users[id]
    if !exists {
        return nil, errors.ErrNotFound
    }
    
    return user, nil
}

func (m *MockUserRepository) GetAll(limit, offset int) ([]*entity.User, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    result := make([]*entity.User, 0)
    for _, user := range m.users {
        result = append(result, user)
    }
    
    return result, nil
}

func (m *MockUserRepository) Update(id int64, user *entity.User) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    _, exists := m.users[id]
    if !exists {
        return errors.ErrNotFound
    }
    
    user.ID = id
    user.UpdatedAt = time.Now()
    m.users[id] = user
    
    return nil
}

func (m *MockUserRepository) Delete(id int64) error {
    m.mutex.Lock()
    defer m.mutex.Unlock()
    
    _, exists := m.users[id]
    if !exists {
        return errors.ErrNotFound
    }
    
    delete(m.users, id)
    return nil
}

func (m *MockUserRepository) GetByEmail(email string) (*entity.User, error) {
    m.mutex.RLock()
    defer m.mutex.RUnlock()
    
    for _, user := range m.users {
        if user.Email == email {
            return user, nil
        }
    }
    
    return nil, errors.ErrNotFound
}