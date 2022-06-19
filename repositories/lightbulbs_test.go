package repositories

import (
	"context"
	"reflect"
	"testing"
)

func TestLightbulbs_Get(t *testing.T) {
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
			l := &Lightbulbs{
				data: tt.fields.data,
			}
			got, err := l.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lightbulbs.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lightbulbs.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightbulbs_GetById(t *testing.T) {
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
			l := &Lightbulbs{
				data: tt.fields.data,
			}
			got, err := l.GetById(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lightbulbs.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Lightbulbs.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightbulbs_Create(t *testing.T) {
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
			l := &Lightbulbs{
				data: tt.fields.data,
			}
			if err := l.Create(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Lightbulbs.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.fields.data, tt.want) {
				t.Errorf("Lightbulbs.Create() = %v, want %v", tt.fields.data, tt.want)
			}
		})
	}
}

func TestLightbulbs_Update(t *testing.T) {
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
			l := &Lightbulbs{
				data: tt.fields.data,
			}
			if err := l.Update(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Lightbulbs.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.fields.data, tt.want) {
				t.Errorf("Lightbulbs.Update() = %v, want %v", tt.fields.data, tt.want)
			}
		})
	}
}

func TestLightbulbs_Delete(t *testing.T) {
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
			l := &Lightbulbs{
				data: tt.fields.data,
			}
			if err := l.Delete(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Lightbulbs.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.fields.data, tt.want) {
				t.Errorf("Lightbulbs.Delete() = %v, want %v", tt.fields.data, tt.want)
			}
		})
	}
}
