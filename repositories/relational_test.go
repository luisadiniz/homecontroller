package repositories

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

type databaseEngineMock struct {
	rows     *sql.Rows
	row      *sql.Row
	result   sql.Result
	queryErr error
	execErr  error
}

func (db *databaseEngineMock) Query(query string, args ...any) (*sql.Rows, error) {
	return db.rows, db.queryErr
}

func (db *databaseEngineMock) QueryRow(query string, args ...any) *sql.Row {
	return db.row
}

func (db *databaseEngineMock) Exec(query string, args ...any) (sql.Result, error) {
	return db.result, db.execErr
}

func TestRelationalRepository_Get(t *testing.T) {
	type fields struct {
		data DatabaseEngine
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "IfDatabaseReturnError",
			fields: fields{
				data: &databaseEngineMock{
					queryErr: errors.New("Data base fail"),
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    map[string]bool{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RelationalRepository{
				data: tt.fields.data,
			}
			got, err := l.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationalRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RelationalRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
