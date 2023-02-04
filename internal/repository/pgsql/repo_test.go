package pgsql

import (
	// golang package
	"context"
	"database/sql"
	"errors"
	"testing"

	// external package
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNewDBRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := NewMockpsqlProvider(ctrl)
	mockInfra := NewMockinfraProvider(ctrl)

	want := &DBRepository{
		db:    mockDB,
		infra: mockInfra,
	}

	assert.Equal(t, want, NewDBRepository(DBRepoParam{
		DB:    mockDB,
		Infra: mockInfra,
	}))
}

func TestDBRepository_BeginTX(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := NewMockpsqlProvider(ctrl)
	mockDB.EXPECT().BeginTx(context.Background(), nil).Return(&sql.Tx{}, nil)

	r := DBRepository{
		db: mockDB,
	}

	want := &sql.Tx{}

	got, err := r.BeginTX(context.Background(), nil)
	assert.Equal(t, want, got)
	assert.Equal(t, nil, err)
}

func TestDBRepository_Commit(t *testing.T) {
	expectedErr := errors.New("error commit")

	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		assert.Nil(t, err)
	}
	defer func() {
		mockDB.Close()
	}()

	tests := []struct {
		name    string
		wantErr error
		mock    func(mock sqlmock.Sqlmock)
	}{
		{
			name: "case_error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(nil)
				mock.ExpectCommit().WillReturnError(expectedErr)
			},
			wantErr: expectedErr,
		},
		{
			name: "success",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(nil)
				mock.ExpectCommit().WillReturnError(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(mock)
			r := DBRepository{
				db: sqlx.NewDb(mockDB, "postgres"),
			}

			tx, err := r.db.BeginTx(context.Background(), nil)
			if err != nil {
				assert.Nil(t, err)
				return
			}

			err = r.Commit(tx)
			assert.Equal(t, test.wantErr, err)
		})
	}
}

func TestDBRepository_Rollback(t *testing.T) {
	expectedErr := errors.New("error rollback")

	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		assert.Nil(t, err)
	}
	defer func() {
		mockDB.Close()
	}()

	tests := []struct {
		name    string
		wantErr error
		mock    func(mock sqlmock.Sqlmock)
	}{
		{
			name: "case_error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(nil)
				mock.ExpectRollback().WillReturnError(expectedErr)
			},
			wantErr: expectedErr,
		},
		{
			name: "success",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(nil)
				mock.ExpectRollback().WillReturnError(nil)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock(mock)
			r := DBRepository{
				db: sqlx.NewDb(mockDB, "postgres"),
			}

			tx, err := r.db.BeginTx(context.Background(), nil)
			if err != nil {
				assert.Nil(t, err)
				return
			}

			err = r.Rollback(tx)
			assert.Equal(t, test.wantErr, err)
		})
	}
}
