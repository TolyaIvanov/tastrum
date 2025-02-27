package handlers_test

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
	handlers "t_astrum/internal/promo/handlers/http"
	"testing"
)

type MockUsecase struct {
	mock.Mock
}

func (m *MockUsecase) CreatePromocode(promocode *entities.Promocode) error {
	args := m.Called(promocode)
	return args.Error(0)
}

func (m *MockUsecase) ApplyPromocode(code string) (*DTOs.PromocodeResponse, error) {
	args := m.Called(code)
	return args.Get(0).(*DTOs.PromocodeResponse), args.Error(1)
}

func TestCreatePromocode(t *testing.T) {
	mockUsecase := new(MockUsecase)              // Создаем mock-объект
	handler := handlers.NewHandlers(mockUsecase) // Используем mockUsecase
	r := gin.Default()

	r.POST("/promocodes", handler.CreatePromocode)

	// Test case 1: Successful promocode creation
	promocodeRequest := `{"code": "ABC123", "max_uses": 10, "reward_id": "reward-uuid"}`
	mockUsecase.On("CreatePromocode", mock.Anything).Return(nil)

	req, _ := http.NewRequest("POST", "/promocodes", bytes.NewBufferString(promocodeRequest))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Test case 2: Promocode already exists
	mockUsecase.On("CreatePromocode", mock.Anything).Return(entities.ErrPromocodeExists)

	req, _ = http.NewRequest("POST", "/promocodes", bytes.NewBufferString(promocodeRequest))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	// Test case 3: Internal server error
	mockUsecase.On("CreatePromocode", mock.Anything).Return(errors.New("internal error"))

	req, _ = http.NewRequest("POST", "/promocodes", bytes.NewBufferString(promocodeRequest))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockUsecase.AssertExpectations(t)
}

func TestApplyPromocode(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := handlers.NewHandlers(mockUsecase) // Передаем mockUsecase
	r := gin.Default()

	r.GET("/promocodes/:code", handler.ApplyPromocode)

	// Test case 1: Successful promocode application
	promocodeResponse := &DTOs.PromocodeResponse{
		Code:        "ABC123",
		CurrentUses: 1,
		MaxUses:     10,
	}

	mockUsecase.On("ApplyPromocode", "ABC123").Return(promocodeResponse, nil)

	req, _ := http.NewRequest("GET", "/promocodes/ABC123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Test case 2: Promocode not found
	mockUsecase.On("ApplyPromocode", "XYZ987").Return(nil, entities.ErrPromocodeNotFound)

	req, _ = http.NewRequest("GET", "/promocodes/XYZ987", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Test case 3: Internal server error
	mockUsecase.On("ApplyPromocode", "ABC123").Return(nil, errors.New("internal error"))

	req, _ = http.NewRequest("GET", "/promocodes/ABC123", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockUsecase.AssertExpectations(t)
}
