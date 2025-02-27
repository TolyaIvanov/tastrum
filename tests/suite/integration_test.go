package integration_test

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"t_astrum/internal/promo/DTOs"
	v1 "t_astrum/internal/promo/handlers/http"
	"t_astrum/internal/promo/repository"
	"t_astrum/internal/promo/usecase"
	"testing"
)

var testRouter *gin.Engine

func setupTestServer(db *sqlx.DB) {
	repo := repository.NewRepository(db)
	uc := usecase.NewPromocodeUsecase(repo)
	h := v1.NewHandlers(uc)
	r := gin.Default()

	r.POST("/admin/promocode", h.CreatePromocode)
	r.GET("/api/promocode/:code", h.ApplyPromocode)
	r.GET("/api/players", h.GetPlayers)
	r.GET("/api/rewards", h.GetRewards)

	testRouter = r
}

func TestCreateAndApplyPromocode(t *testing.T) {
	// Подключение к тестовой БД
	db, err := sqlx.Open("postgres", "host=localhost user=test password=test dbname=testdb sslmode=disable")
	require.NoError(t, err)
	setupTestServer(db)

	rewardID := uuid.New()
	newPromo := DTOs.CreatePromocodeRequest{
		Code:     "TEST123",
		MaxUses:  5,
		RewardId: rewardID,
	}

	// Тест создания промокода
	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(newPromo)
	req, _ := http.NewRequest("POST", "/admin/promocode", jsonBody(reqBody))
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var createdPromo DTOs.PromocodeResponse
	json.Unmarshal(w.Body.Bytes(), &createdPromo)
	assert.Equal(t, "TEST123", createdPromo.Code)
	assert.Equal(t, 5, createdPromo.MaxUses)

	// Тест применения промокода
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/promocode/TEST123", nil)
	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var appliedPromo DTOs.PromocodeResponse
	json.Unmarshal(w.Body.Bytes(), &appliedPromo)
	assert.Equal(t, "TEST123", appliedPromo.Code)
	assert.Equal(t, 1, appliedPromo.CurrentUses)
}

func jsonBody(data interface{}) io.Reader {
	body, _ := json.Marshal(data)
	return bytes.NewBuffer(body)
}
