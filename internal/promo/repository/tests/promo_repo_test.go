package repository_test

import (
	"bytes"
	_ "database/sql"
	"errors"
	_ "errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"t_astrum/internal/promo/entities"
	handlers "t_astrum/internal/promo/handlers/http"
	_ "t_astrum/internal/promo/repository"
	"testing"
)

type MockUsecase struct {
	mock.Mock
}

// Реализация интерфейса UsecaseInterface
func (m *MockUsecase) CreatePromocode(promocode *entities.Promocode) error {
	args := m.Called(promocode)
	return args.Error(0)
}

func (m *MockUsecase) ApplyPromocode(code string) (*entities.Promocode, error) {
	args := m.Called(code)
	if args.Get(0) != nil {
		return args.Get(0).(*entities.Promocode), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUsecase) GetPlayers() ([]entities.Player, error) {
	args := m.Called()
	return args.Get(0).([]entities.Player), args.Error(1)
}

func (m *MockUsecase) GetRewards() ([]entities.Reward, error) {
	args := m.Called()
	return args.Get(0).([]entities.Reward), args.Error(1)
}

func TestCreatePromocode(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := handlers.NewHandlers(mockUsecase, mockUsecase, mockUsecase)
	r := gin.Default()

	r.POST("/promocodes", handler.CreatePromocode)

	// Test case 1: Successful promocode creation
	rewardID := uuid.NewString()
	promocodeRequest := `{"code": "ABC123", "max_uses": 10, "reward_id": "` + rewardID + `"}`
	mockUsecase.On("CreatePromocode", mock.Anything).Return(nil).Once()

	req, _ := http.NewRequest("POST", "/promocodes", bytes.NewBufferString(promocodeRequest))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Test case 2: Promocode already exists
	mockUsecase.On("CreatePromocode", mock.Anything).Return(entities.ErrPromocodeExists).Once()

	req, _ = http.NewRequest("POST", "/promocodes", bytes.NewBufferString(promocodeRequest))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

	// Test case 3: Internal server error
	mockUsecase.On("CreatePromocode", mock.Anything).Return(errors.New("internal error")).Once()

	req, _ = http.NewRequest("POST", "/promocodes", bytes.NewBufferString(promocodeRequest))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockUsecase.AssertExpectations(t)
}

func TestApplyPromocode(t *testing.T) {
	mockUsecase := new(MockUsecase)
	handler := handlers.NewHandlers(mockUsecase, mockUsecase, mockUsecase)
	r := gin.Default()

	r.GET("/promocodes/:code", handler.ApplyPromocode)

	// 1: Successful
	promocodeResponse := &entities.Promocode{
		Code:      "ABC123",
		UsesCount: 1,
		MaxUses:   10,
	}

	mockUsecase.On("ApplyPromocode", "ABC123").Return(promocodeResponse, nil).Once()

	req, _ := http.NewRequest("GET", "/promocodes/ABC123", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// 2: Promocode not found
	mockUsecase.On("ApplyPromocode", "XYZ987").Return(nil, entities.ErrPromocodeNotFound).Once()

	req, _ = http.NewRequest("GET", "/promocodes/XYZ987", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// 3: Internal server error
	mockUsecase.On("ApplyPromocode", "ABC123").Return(nil, errors.New("internal error")).Once()

	req, _ = http.NewRequest("GET", "/promocodes/ABC123", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	mockUsecase.AssertExpectations(t)
}
