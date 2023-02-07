package account

import (
	// golang package
	"context"
	"testing"

	// external package
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_GenerateJWT(t *testing.T) {
	type mockFields struct {
		infra *MockinfraProvider
		rsc   *MockresourceProvider
	}
	type args struct {
		userID int64
		email  string
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		want       string
		wantErr    error
	}{
		{
			name: "when_GetJWTFromCache_error_then_return_error",
			args: args{userID: 123},
			mockFields: func(mf mockFields) {
				mf.rsc.EXPECT().GetJWTFromCache(context.Background(), int64(123)).Return("", assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_token_exist_then_return_existing_token",
			args: args{userID: 123},
			mockFields: func(mf mockFields) {
				mf.rsc.EXPECT().GetJWTFromCache(context.Background(), int64(123)).Return("token", nil)
			},
			want: "token",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				rsc:   NewMockresourceProvider(ctrl),
				infra: NewMockinfraProvider(ctrl),
			}
			test.mockFields(mockFields)

			svc := &Service{
				infra: mockFields.infra,
				rsc:   mockFields.rsc,
			}

			got, err := svc.GenerateJWT(context.Background(), test.args.userID, test.args.email)
			assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestService_InvalidateJWT(t *testing.T) {
	type mockFields struct {
		rsc *MockresourceProvider
	}
	type args struct {
		userID int64
	}
	tests := []struct {
		name       string
		args       args
		mockFields func(mockFields)
		wantErr    error
	}{
		{
			name: "when_DeleteJWTInCache_error_then_return_error",
			args: args{userID: 123},
			mockFields: func(mf mockFields) {
				mf.rsc.EXPECT().DeleteJWTInCache(context.Background(), int64(123)).Return(assert.AnError)
			},
			wantErr: assert.AnError,
		},
		{
			name: "when_no_error_occured_then_return_nil",
			args: args{userID: 123},
			mockFields: func(mf mockFields) {
				mf.rsc.EXPECT().DeleteJWTInCache(context.Background(), int64(123)).Return(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockFields := mockFields{
				rsc: NewMockresourceProvider(ctrl),
			}
			test.mockFields(mockFields)

			svc := &Service{
				rsc: mockFields.rsc,
			}

			err := svc.InvalidateJWT(context.Background(), test.args.userID)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
