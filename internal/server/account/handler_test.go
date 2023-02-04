package account

import (
	// golang package
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAccountUC := NewMockaccountUCManager(ctrl)
	mockInfra := NewMockinfraProvider(ctrl)

	want := &Handler{
		account: mockAccountUC,
		infra:   mockInfra,
	}

	assert.Equal(t, want, NewHandler(AccountHandlerParam{
		Account: mockAccountUC,
		Infra:   mockInfra,
	}))
}
