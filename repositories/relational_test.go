package repositories

import (
	"context"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

func TestRelationalRepository_Get(t *testing.T) {
	type fields struct {
		data func() DatabaseEngine
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
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					mock.ExpectQuery("SELECT name, isOn FROM lightbulbs").WillReturnError(errors.New("Database failed"))
					return db
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    map[string]bool{},
			wantErr: true,
		},
		{
			name: "WhenDatabaseReturnsOneRow_ShouldReturnMapWithSingleEntryAndNoError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("sala", true)

					mock.ExpectQuery("SELECT name, isOn FROM lightbulbs").WillReturnRows(rows)
					return db
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: map[string]bool{
				"sala": true,
			},
			wantErr: false,
		},
		{
			name: "WhenDatabaseReturnsThreeRows_ShouldReturnMapWithThreeEntriesAndNoError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("sala", true)
					rows.AddRow("cozinha", false)
					rows.AddRow("quarto", true)

					mock.ExpectQuery("SELECT name, isOn FROM lightbulbs").WillReturnRows(rows)
					return db
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: map[string]bool{
				"sala":    true,
				"cozinha": false,
				"quarto":  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RelationalRepository{
				data: tt.fields.data(),
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

func TestRelationalRepository_GetById(t *testing.T) {
	type fields struct {
		data func() DatabaseEngine
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "IfDatabaseReturnError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					mock.ExpectQuery(regexp.QuoteMeta("SELECT name, isOn FROM lightbulbs WHERE name=$1")).WithArgs("sala").WillReturnError(errors.New("Database failed"))
					return db
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "sala",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "WhenDatabaseUpdateValue_ShouldReturnNoError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("cozinha", true)

					mock.ExpectQuery(regexp.QuoteMeta("SELECT name, isOn FROM lightbulbs WHERE name=$1")).WithArgs("cozinha").WillReturnRows(rows)
					return db
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "cozinha",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RelationalRepository{
				data: tt.fields.data(),
			}
			got, err := l.GetById(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationalRepository.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RelationalRepository.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRelationalRepository_Delete(t *testing.T) {
	type fields struct {
		data func() DatabaseEngine
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "IfDatabaseReturnError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM lightbulbs WHERE name=$1")).WithArgs("sala").WillReturnError(errors.New("Delete failed!"))
					return db
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "sala",
			},
			wantErr: true,
		},
		{
			name: "WhenDatabaseDeletesValue_ShouldReturnNoError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("cozinha", true)

					mock.ExpectExec(regexp.QuoteMeta("DELETE FROM lightbulbs WHERE name=$1")).WithArgs("cozinha").WillReturnResult(sqlmock.NewResult(0, 0))

					return db
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "cozinha",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RelationalRepository{
				data: tt.fields.data(),
			}
			if err := l.Delete(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("RelationalRepository.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRelationalRepository_Create(t *testing.T) {
	type fields struct {
		data func() DatabaseEngine
	}
	type args struct {
		ctx   context.Context
		name  string
		value bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "IfDatabaseReturnError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO lightbulbs (name, isOn) VALUES ($1, $2)")).WithArgs("quarto", false).WillReturnError(errors.New("Creation failed!"))
					return db
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: false,
			},
			wantErr: true,
		},
		{
			name: "WhenDatabaseCreateNewEntry_ShouldReturnNoError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("cozinha", true)

					mock.ExpectExec(regexp.QuoteMeta("INSERT INTO lightbulbs (name, isOn) VALUES ($1, $2)")).WithArgs("quarto", true).WillReturnResult(sqlmock.NewResult(0, 0))

					return db
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RelationalRepository{
				data: tt.fields.data(),
			}
			if err := l.Create(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RelationalRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRelationalRepository_Update(t *testing.T) {
	type fields struct {
		data func() DatabaseEngine
	}
	type args struct {
		ctx   context.Context
		name  string
		value bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "IfDatabaseReturnError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()
					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("quarto", true)

					mock.ExpectExec(regexp.QuoteMeta("UPDATE lightbulbs SET isOn=$2 WHERE name=$1;")).WithArgs("quarto", false).WillReturnError(errors.New("Update Failed!"))
					return db
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: false,
			},
			wantErr: true,
		},
		{
			name: "WhenDatabaseUpdatesValue_ShouldReturnNoError",
			fields: fields{
				data: func() DatabaseEngine {
					db, mock, _ := sqlmock.New()

					rows := mock.NewRows([]string{
						"name", "isOn",
					})
					rows.AddRow("quarto", true)

					mock.ExpectExec(regexp.QuoteMeta("UPDATE lightbulbs SET isOn=$2 WHERE name=$1;")).WithArgs("quarto", false).WillReturnResult(sqlmock.NewResult(0, 0))

					return db
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RelationalRepository{
				data: tt.fields.data(),
			}
			if err := l.Update(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RelationalRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
