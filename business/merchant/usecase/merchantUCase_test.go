package usecase

import (
	"errors"
	"himanshu/poc/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/* Merchant Repository Mock Object and corresponding functions */
type MerchantRepoMock struct {
	mock.Mock
}

func (m *MerchantRepoMock) AddMerchant(merchant entity.Merchant) (int, error) {
	ret := m.Called(merchant)
	if ret.Get(1) != nil {
		return 0, ret.Get(1).(error)
	}
	return ret.Get(0).(int), nil
}

func (m *MerchantRepoMock) DeleteMerchant(merchant entity.Merchant) error {
	return nil
}

func (m *MerchantRepoMock) GetMerchantById(mid int) (entity.Merchant, error) {
	ret := m.Called(mid)
	if ret.Get(1) != nil {
		return entity.Merchant{}, ret.Get(1).(error)
	}
	return entity.Merchant{}, nil
}

func (m *MerchantRepoMock) UpdateMerchant(merchant entity.Merchant) error {
	ret := m.Called(merchant)
	return ret.Get(0).(error)
}

func Test_merchantUCase_RegisterMerchantSuccess(t *testing.T) {
	m := new(MerchantRepoMock)
	t.Run("TestGetStoresForMerchant", func(t *testing.T) {
		muc := &merchantUCase{merchantRepo: m}
		m.On("AddMerchant", mock.Anything).Return(1, nil)
		merchant := entity.Merchant{Name: "barista", EmailId: "barista@gmail.com"}
		got, err := muc.RegisterMerchant(merchant)
		assert.Equal(t, got, 1)
		assert.Equal(t, err, nil)
	})
}

func Test_merchantUCase_RegisterMerchantFail(t *testing.T) {
	m := new(MerchantRepoMock)
	t.Run("TestGetStoresForMerchant", func(t *testing.T) {
		muc := &merchantUCase{merchantRepo: m}
		m.On("AddMerchant", mock.Anything).Return(0, errors.New("failed to create merchant"))
		merchant := entity.Merchant{Name: "barista", EmailId: "barista@gmail.com"}
		got, err := muc.RegisterMerchant(merchant)
		assert.Equal(t, got, 0)
		assert.NotEmpty(t, err)
	})
}
