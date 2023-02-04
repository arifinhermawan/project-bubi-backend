package server

import (
	// golang package
	"io"
	"testing"
	"time"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

func TestNewInfra(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConfig := NewMockconfigProvider(ctrl)
	mockGolang := NewMockgolangProvider(ctrl)
	mockReader := NewMockreaderProvider(ctrl)

	want := &Infra{
		Config: mockConfig,
		Golang: mockGolang,
		Reader: mockReader,
	}

	got := NewInfra(InfraParam{
		Config: mockConfig,
		Golang: mockGolang,
		Reader: mockReader,
	})

	assert.Equal(t, want, got)
}

func TestInfra_GetConfig(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConfig := NewMockconfigProvider(ctrl)
	mockConfig.EXPECT().GetConfig().Return(&configuration.AppConfig{})

	want := &configuration.AppConfig{}

	i := &Infra{
		Config: mockConfig,
	}

	got := i.GetConfig()
	assert.Equal(t, want, got)
}

func TestInfra_GetTimeGMT7(t *testing.T) {
	mockDate := time.Date(1993, 5, 16, 0, 0, 0, 0, time.UTC)

	ctrl := gomock.NewController(t)
	mockGolang := NewMockgolangProvider(ctrl)
	mockGolang.EXPECT().GetTimeGMT7().Return(mockDate)

	want := mockDate

	i := &Infra{
		Golang: mockGolang,
	}

	got := i.GetTimeGMT7()
	assert.Equal(t, want, got)
}

func TestInfra_JsonMarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGolang := NewMockgolangProvider(ctrl)
	mockGolang.EXPECT().JsonMarshal("abc").Return([]byte("abc"), nil)

	want := []byte("abc")

	i := &Infra{
		Golang: mockGolang,
	}

	got, _ := i.JsonMarshal("abc")
	assert.Equal(t, want, got)
}

func TestInfra_JsonUnmarshal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockGolang := NewMockgolangProvider(ctrl)

	var got string
	mockGolang.EXPECT().JsonUnmarshal([]byte("abc"), &got).DoAndReturn(
		func(input []byte, dest interface{}) error {
			got = "abc"
			return nil
		})

	want := "abc"

	i := &Infra{
		Golang: mockGolang,
	}

	_ = i.JsonUnmarshal([]byte("abc"), &got)
	assert.Equal(t, want, got)
}

func TestInfra_ReadAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockReader := NewMockreaderProvider(ctrl)
	mockReader.EXPECT().ReadAll(&io.LimitedReader{}).Return([]byte(nil), nil)

	i := &Infra{
		Reader: mockReader,
	}

	want := []byte(nil)

	got, _ := i.ReadAll(&io.LimitedReader{})
	assert.Equal(t, want, got)
}
