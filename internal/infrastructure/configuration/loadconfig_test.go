package configuration

import (
	// golang package
	"sync"
	"testing"

	// external package
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration_GetConfig(t *testing.T) {
	viperReadInConfigOri := viperReadInConfig
	defer func() {
		viperReadInConfig = viperReadInConfigOri
	}()

	tests := []struct {
		name             string
		mockReadInConfig func() error
		want             AppConfig
	}{
		{
			name:             "when_no_error_occured_then_return_nil",
			mockReadInConfig: func() error { return nil },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viperReadInConfig = test.mockReadInConfig

			c := Configuration{
				doLoadConfigOnce: new(sync.Once),
			}

			got := c.GetConfig()
			assert.Equal(t, test.want, got)
		})
	}
}

func TestConfiguration_LoadConfig(t *testing.T) {
	viperReadInConfigOri := viperReadInConfig
	viperUnmarshalOri := viperUnmarshal
	defer func() {
		viperReadInConfig = viperReadInConfigOri
		viperUnmarshal = viperUnmarshalOri
	}()
	tests := []struct {
		name             string
		mockReadInConfig func() error
		mockUnmarshal    func(rawVal interface{}, opts ...viper.DecoderConfigOption) error
		want             AppConfig
		wantErr          error
	}{
		{
			name: "when_viperReadInConfig_error_then_return_error",
			mockReadInConfig: func() error {
				return assert.AnError
			},
			wantErr: errConfigNotFound,
		},
		{
			name: "when_viperUnmarshal_error_then_return_error",
			mockReadInConfig: func() error {
				return nil
			},
			mockUnmarshal: func(rawVal interface{}, opts ...viper.DecoderConfigOption) error {
				return assert.AnError
			},
			wantErr: errConfigFailedToUnmarshal,
		},
		{
			name:             "when_no_error_occured_then_return_nil_error",
			mockReadInConfig: func() error { return nil },
			mockUnmarshal:    func(rawVal interface{}, opts ...viper.DecoderConfigOption) error { return nil },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			viperReadInConfig = test.mockReadInConfig
			viperUnmarshal = test.mockUnmarshal

			c := Configuration{}
			got, err := c.LoadConfig()
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
