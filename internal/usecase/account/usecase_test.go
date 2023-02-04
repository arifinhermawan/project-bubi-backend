package account

import (
	// golang package
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAccountSvc := NewMockaccountServiceProvider(ctrl)

	want := &UseCase{
		account: mockAccountSvc,
	}
	assert.Equal(t, want, NewUseCase(AccountUsecaseParam{Account: mockAccountSvc}))
}
