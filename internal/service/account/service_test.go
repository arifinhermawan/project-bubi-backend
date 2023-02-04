package account

import (
	// golang package
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockResource := *NewMockresourceProvider(ctrl)

	want := &Service{
		rsc: &mockResource,
	}
	assert.Equal(t, want, NewService(AccountServiceParam{Rsc: &mockResource}))
}
