package repositories

import (
	"context"
	"reflect"
	"testing"
)

func TestInMemoryDB_Get(t *testing.T) {
	type fields struct {
		data map[string]bool
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
			name: "WhenEmptyReturnsEmpty",
			fields: fields{
				map[string]bool{},
			},
			args: args{
				ctx: context.Background(),
			},
			want:    map[string]bool{},
			wantErr: false,
		},

		{
			name: "WhenHasValuesReturnValues",
			fields: fields{
				map[string]bool{
					"sala":   true,
					"quarto": false,
				},
			},
			args: args{
				ctx: context.Background(),
			},
			want: map[string]bool{
				"sala":   true,
				"quarto": false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &InMemoryDB{
				data: tt.fields.data,
			}
			got, err := l.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryDB.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InMemoryDB.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryDB_GetById(t *testing.T) {
	type fields struct {
		data map[string]bool
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
			name: "WhenNameExistsReturnValue",
			fields: fields{
				data: map[string]bool{
					"sala":   true,
					"quarto": false,
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "sala",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "WhenNameDontExistReturnNothing",
			fields: fields{
				map[string]bool{
					"sala":   true,
					"quarto": false,
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "cozinha",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &InMemoryDB{
				data: tt.fields.data,
			}
			got, err := l.GetById(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("InMemoryDB.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InMemoryDB.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryDB_Create(t *testing.T) {
	type fields struct {
		data map[string]bool
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
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "CreateNewLightbulb",
			fields: fields{
				data: map[string]bool{
					"sala": true,
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: false,
			},
			want: map[string]bool{
				"sala":   true,
				"quarto": false,
			},
			wantErr: false,
		},
		{
			name: "CreateLightbulbThatAlreadyExist",
			fields: fields{
				data: map[string]bool{
					"sala":   true,
					"quarto": false,
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: false,
			},
			want: map[string]bool{
				"sala":   true,
				"quarto": false,
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &InMemoryDB{
				data: tt.fields.data,
			}
			if err := l.Create(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryDB.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.fields.data, tt.want) {
				t.Errorf("InMemoryDB.Create() = %v, want %v", tt.fields.data, tt.want)
			}
		})
	}
}

func TestInMemoryDB_Update(t *testing.T) {
	type fields struct {
		data map[string]bool
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
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "UpdateLightbulbThatExist",
			fields: fields{
				data: map[string]bool{
					"quarto": false,
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: true,
			},
			want: map[string]bool{
				"quarto": true,
			},
			wantErr: false,
		},

		{
			name: "UpdateLightbulbThatDoNotExist",
			fields: fields{
				data: map[string]bool{
					"sala": false,
				},
			},
			args: args{
				ctx:   context.Background(),
				name:  "quarto",
				value: true,
			},
			want: map[string]bool{
				"sala":   false,
				"quarto": true,
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &InMemoryDB{
				data: tt.fields.data,
			}
			if err := l.Update(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryDB.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.fields.data, tt.want) {
				t.Errorf("InMemoryDB.Update() = %v, want %v", tt.fields.data, tt.want)
			}
		})
	}
}

func TestInMemoryDB_Delete(t *testing.T) {
	type fields struct {
		data map[string]bool
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "DeleteLightbulbThatExist",
			fields: fields{
				data: map[string]bool{
					"quarto": false,
					"sala":   true,
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "quarto",
			},
			want: map[string]bool{
				"sala": true,
			},
			wantErr: false,
		},
		{
			name: "DeleteLightbulbThatDoNotExist",
			fields: fields{
				data: map[string]bool{
					"sala":    true,
					"cozinha": false,
				},
			},
			args: args{
				ctx:  context.Background(),
				name: "quarto",
			},

			want: map[string]bool{
				"sala":    true,
				"cozinha": false,
			},
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &InMemoryDB{
				data: tt.fields.data,
			}
			if err := l.Delete(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("InMemoryDB.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.fields.data, tt.want) {
				t.Errorf("InMemoryDB.Delete() = %v, want %v", tt.fields.data, tt.want)
			}
		})
	}
}
