package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	r "github.com/moemoe89/go-unit-test-sql/repository"
	"github.com/stretchr/testify/assert"
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
					rows:     nil,
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

var u = &r.UserModel{
	Name: "sala",
	IsOn: true,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestGetByID_IfHasRowReturnValue(t *testing.T) {
	db, mock := NewMock()
	repo := &RelationalRepository{db}

	query := "SELECT name, isOn FROM lighbulbs WHERE name=$1"

	rows := sqlmock.NewRows([]string{"name", "isOn"}).
		AddRow(u.Name, u.IsOn)

	mock.ExpectQuery(query).WithArgs(u.Name).WillReturnRows(rows)

	lightbulb, err := repo.GetById(context.Background(), u.Name)
	assert.NotNil(t, lightbulb)
	assert.NoError(t, err)

	want := true

	if !reflect.DeepEqual(lightbulb, want) {
		println("RelationalRepository.GetById() = %v, want %v", lightbulb, want)
	}
}
func TestCreate_IfItWorksReturnNoError(t *testing.T) {
	db, mock := NewMock()
	repo := &RelationalRepository{db}

	query := "INSERT INTO lightbulbs (name, isOn) VALUES ($1, $2)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("cozinha", true).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Create(context.Background(), "cozinha", true)
	if err != nil {
		println("RelationalRepository.Create() error = %v", err)
		return
	}

	want := map[string]bool{
		"cozinha": true,
	}

	if !reflect.DeepEqual(mock, want) {
		println("RelationalRepository.Create() = %v, want %v", mock, want)
	}
}
func TestUpdate_IfItWorksReturnNoError(t *testing.T) {
	db, mock := NewMock()
	repo := &RelationalRepository{db}

	mock.NewRows([]string{"name", "isOn"}).
		AddRow("cozinha", true)

	query := "UPDATE lightbulbs SET isOn=$2 WHERE name=$1;"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("cozinha", false).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(context.Background(), "cozinha", false)
	println("RelationalRepository.Update() error = %v", err)

	want := map[string]bool{
		"cozinha": false,
	}

	if !reflect.DeepEqual(mock, want) {
		println("RelationalRepository.Update() = %v, want %v", mock, want)
	}
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	repo := &RelationalRepository{db}

	mock.NewRows([]string{"name", "isOn"}).
		AddRow("sala", true).
		AddRow("cozinha", false)

	query := "DELETE FROM lightbulbs WHERE name=sala"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(context.Background(), "quarto")
	println("RelationalRepository.Detele() error = %v", err)

	want := map[string]bool{}

	if !reflect.DeepEqual(mock, want) {
		println("RelationalRepository.Delete() = %v, want %v", mock, want)
	}
}
