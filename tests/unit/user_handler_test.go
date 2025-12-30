package unit

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    
    "go-setup/internal/entity"
    "go-setup/internal/delivery/http"
    "go-setup/tests/mocks"
)

func TestCreateUserHandler(t *testing.T) {
    // Table-driven test approach
    tests := []struct {
        name           string
        request        entity.CreateUserRequest
        expectedStatus int
        expectError    bool
    }{
        {
            name:           "Success",
            request:        entity.CreateUserRequest{Name: "John", Email: "john@example.com"},
            expectedStatus: 201,
            expectError:    false,
        },
        {
            name:           "Empty Name",
            request:        entity.CreateUserRequest{Name: "", Email: "john@example.com"},
            expectedStatus: 400,
            expectError:    true,
        },
        {
            name:           "Invalid Email",
            request:        entity.CreateUserRequest{Name: "John", Email: "invalid"},
            expectedStatus: 400,
            expectError:    true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Setup
            mockRepo := mocks.NewMockUserRepository()
            mockUC := usecase.NewUserUsecase(mockRepo)
            handler := http.NewUserHandler(mockUC)
            
            // Create request
            body, _ := json.Marshal(tt.request)
            req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(body))
            req.Header.Set("Content-Type", "application/json")
            
            // Record response
            recorder := httptest.NewRecorder()
            gin.SetMode(gin.TestMode)
            router := gin.New()
            router.POST("/api/users", handler.CreateUser)
            
            router.ServeHTTP(recorder, req)
            
            // Assert
            assert.Equal(t, tt.expectedStatus, recorder.Code)
        })
    }
}