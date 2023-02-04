package account

import (
	// golang package
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewResource(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := NewMockdbRepoProvider(ctrl)

	want := &Resource{
		db: mockDB,
	}
	assert.Equal(t, want, NewResource(AccountResourceParam{DB: mockDB}))
}
