package repository_test

import (
	"database/sql"
	_ "errors"
	_ "github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/repository"
	"testing"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Get(dest interface{}, query string, args ...interface{}) error {
	args := m.Called(dest, query, args)
	return args.Error(0)
}

func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	args := m.Called(query, args)
	return nil, args.Error(0)
}

func TestPromocodeExists(t *testing.T) {
	mockDB := new(MockDB)
	repo := &repository.Repository{DB: mockDB}

	mockDB.On("Get", mock.Anything, "SELECT COUNT(*) FROM promocodes WHERE code=$1", "ABC123").Return(nil).Run(func(args mock.Arguments) {
		*args.Get(0).(*int) = 1
	})

	exists, err := repo.PromocodeExists("ABC123")
	assert.NoError(t, err)
	assert.True(t, exists)

	mockDB.AssertExpectations(t)
}

func TestCreatePromocode(t *testing.T) {
	mockDB := new(MockDB)
	repo := repository.NewRepository(mockDB)

	promocode := &entities.Promocode{
		Code:      "ABC123",
		MaxUses:   10,
		UsesCount: 0,
		RewardId:  "reward-uuid",
	}

	mockDB.On("Exec", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	err := repo.CreatePromocode(promocode)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

func TestApplyPromocode_Success(t *testing.T) {
	mockDB := new(MockDB)
	repo := repository.NewRepository(mockDB)

	promocode := &entities.Promocode{
		ID:        "promocode-id",
		Code:      "ABC123",
		MaxUses:   10,
		UsesCount: 0,
		RewardId:  "reward-id",
	}

	mockDB.On("Get", mock.Anything, "SELECT id, code, max_uses, uses_count, reward_id FROM promocodes WHERE code=$1", "ABC123").Return(nil).Run(func(args mock.Arguments) {
		*args.Get(0).(*entities.Promocode) = *promocode
	})

	mockDB.On("Exec", "UPDATE promocodes SET uses_count = uses_count + 1 WHERE code = $1", "ABC123").Return(nil)

	result, err := repo.ApplyPromocode("ABC123")
	assert.NoError(t, err)
	assert.Equal(t, promocode.Code, result.Code)

	mockDB.AssertExpectations(t)
}

func TestApplyPromocode_MaxUsesReached(t *testing.T) {
	mockDB := new(MockDB)
	repo := repository.NewRepository(mockDB)

	promocode := &entities.Promocode{
		ID:        "promocode-id",
		Code:      "ABC123",
		MaxUses:   1,
		UsesCount: 1,
		RewardId:  "reward-id",
	}

	mockDB.On("Get", mock.Anything, "SELECT id, code, max_uses, uses_count, reward_id FROM promocodes WHERE code=$1", "ABC123").Return(nil).Run(func(args mock.Arguments) {
		*args.Get(0).(*entities.Promocode) = *promocode
	})

	result, err := repo.ApplyPromocode("ABC123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, entities.ErrPromocodeMaxUsesReached, err)

	mockDB.AssertExpectations(t)
}
