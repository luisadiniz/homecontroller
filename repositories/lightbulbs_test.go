package repositories

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// )

// func TestGetLightbulbs(t *testing.T) {
// 	type args struct {
// 		ctx context.Context
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		setup   func()
// 		want    map[string]bool
// 		wantErr bool
// 	}{

// 		{
// 			name: "WhenEmptyReturnsEmpty",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{}
// 			},
// 			want:    map[string]bool{},
// 			wantErr: false,
// 		},

// 		{
// 			name: "WhenHasValuesReturnValues",
// 			args: args{
// 				ctx: context.Background(),
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala":   true,
// 					"quarto": false,
// 				}
// 			},
// 			want: map[string]bool{
// 				"sala":   true,
// 				"quarto": false,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setup()
// 			got, err := GetLightbulbs(tt.args.ctx)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetLightbulbs() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetLightbulbs() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetLightbulbById(t *testing.T) {
// 	type args struct {
// 		ctx  context.Context
// 		name string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		setup   func()
// 		want    bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "WhenNameExistsReturnValue",
// 			args: args{
// 				ctx:  context.Background(),
// 				name: "sala",
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala":   true,
// 					"quarto": false,
// 				}
// 			},
// 			want:    true,
// 			wantErr: false,
// 		},
// 		{
// 			name: "WhenNameDontExistReturnNothing",
// 			args: args{
// 				ctx:  context.Background(),
// 				name: "cozinha",
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala":   true,
// 					"quarto": false,
// 				}
// 			},
// 			want:    false,
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setup()
// 			got, err := GetLightbulbById(tt.args.ctx, tt.args.name)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetLightbulbById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("GetLightbulbById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestCreateLightbulbs(t *testing.T) {
// 	type args struct {
// 		ctx   context.Context
// 		name  string
// 		value bool
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		setup   func()
// 		want    map[string]bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "CreateNewLightbulb",
// 			args: args{
// 				ctx:   context.Background(),
// 				name:  "quarto",
// 				value: false,
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala": true,
// 				}
// 			},
// 			want: map[string]bool{
// 				"sala":   true,
// 				"quarto": false,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "CreateLightbulbThatAlreadyExist",
// 			args: args{
// 				ctx:   context.Background(),
// 				name:  "quarto",
// 				value: false,
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala":   true,
// 					"quarto": false,
// 				}
// 			},
// 			want: map[string]bool{
// 				"sala":   true,
// 				"quarto": false,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setup()
// 			if err := CreateLightbulbs(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
// 				t.Errorf("CreateLightbulbs() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if !reflect.DeepEqual(lightbulbs, tt.want) {
// 				t.Errorf("CreateLightbulbs() = %v, want %v", lightbulbs, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUpdateLightbulb(t *testing.T) {
// 	type args struct {
// 		ctx   context.Context
// 		name  string
// 		value bool
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		setup   func()
// 		want    map[string]bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "UpdateLightbulbThatExist",
// 			args: args{
// 				ctx:   context.Background(),
// 				name:  "quarto",
// 				value: true,
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"quarto": false,
// 				}
// 			},
// 			want: map[string]bool{
// 				"quarto": true,
// 			},
// 			wantErr: false,
// 		},

// 		{
// 			name: "UpdateLightbulbThatDoNotExist",
// 			args: args{
// 				ctx:   context.Background(),
// 				name:  "quarto",
// 				value: true,
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala": false,
// 				}
// 			},
// 			want: map[string]bool{
// 				"sala":   false,
// 				"quarto": true,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setup()
// 			if err := UpdateLightbulb(tt.args.ctx, tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
// 				t.Errorf("UpdateLightbulb() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if !reflect.DeepEqual(lightbulbs, tt.want) {
// 				t.Errorf("UpdateLightbulb() = %v, want %v", lightbulbs, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDeleteLightbulb(t *testing.T) {
// 	type args struct {
// 		ctx  context.Context
// 		name string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		setup   func()
// 		want    map[string]bool
// 		wantErr bool
// 	}{
// 		{
// 			name: "DeleteLightbulbThatExist",
// 			args: args{
// 				ctx:  context.Background(),
// 				name: "quarto",
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala":   true,
// 					"quarto": false,
// 				}
// 			},
// 			want: map[string]bool{
// 				"sala": true,
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "DeleteLightbulbThatDoNotExist",
// 			args: args{
// 				ctx:  context.Background(),
// 				name: "quarto",
// 			},
// 			setup: func() {
// 				lightbulbs = map[string]bool{
// 					"sala":    true,
// 					"cozinha": false,
// 				}
// 			},
// 			want: map[string]bool{
// 				"sala":    true,
// 				"cozinha": false,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.setup()
// 			if err := DeleteLightbulb(tt.args.ctx, tt.args.name); (err != nil) != tt.wantErr {
// 				t.Errorf("DeleteLightbulb() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if !reflect.DeepEqual(lightbulbs, tt.want) {
// 				t.Errorf("DeleteLightbulb() = %v, want %v", lightbulbs, tt.want)
// 			}
// 		})
// 	}
// }
