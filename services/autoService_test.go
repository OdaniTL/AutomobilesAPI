package services

import (
	"AutomobilesAPI/configs"
	"AutomobilesAPI/dtos/request"
	"AutomobilesAPI/dtos/response"
	"AutomobilesAPI/repository"
	"context"
	"reflect"
	"testing"
)

/*
"brand": "Toyota",
"model": "Corolla",
"year": 2023,
"body_type": "sedan",
"engine_type": "gasoline",
"displacement": 1598,
"price": 25000.50,
"state": true
*/
func TestAutoServices_Create(t *testing.T) {

	connection, _ := configs.NewConnection()
	autosRepository := repository.NewAutosRepository(connection)
	srv := NewAutoServices(autosRepository)

	type args struct {
		request request.AutosCreateRequest
	}
	ctx := context.Background()

	//srv := &AutoServicesMock{}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Empty model AutosCreateRequest",
			args: args{
				request: request.AutosCreateRequest{
					Brand:      "",
					Model:      "",
					Year:       10,
					EngineType: "",
					Price:      0,
				},
			},
			wantErr: true,
		},
		{
			name: "Nil Brand and Model",
			args: args{
				request: request.AutosCreateRequest{},
			},
			wantErr: true,
		},
		{
			name: "Empty BodyType",
			args: args{
				request: request.AutosCreateRequest{
					Brand:      "testBran",
					Model:      "TestModel",
					Year:       1999,
					EngineType: "EngineTypeTest",
					Price:      1505,
					BodyType:   "",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := srv.Create(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AutoServices.Create() expected = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestAutoServices_Delete(t *testing.T) {
	type args struct {
		autoId int
	}
	ctx := context.Background()
	tests := []struct {
		name    string
		p       *AutoServices
		args    args
		wantErr bool
	}{
		{
			name:    "Auto by Id 0",
			p:       &AutoServices{},
			args:    args{autoId: 0},
			wantErr: true,
		},
		{
			name:    "Invalid AutoId",
			p:       &AutoServices{},
			args:    args{autoId: -1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.p.Delete(ctx, tt.args.autoId); (err != nil) != tt.wantErr {
				t.Errorf("AutoServices.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAutoServices_FindByName(t *testing.T) {
	type args struct {
		brandName string
	}
	ctx := context.Background()
	srvs := &AutoServices{}

	tests := []struct {
		name string

		args    args
		want    []response.AutosResponse
		wantErr bool
	}{
		{
			name:    "Invali Brand",
			args:    args{brandName: " "},
			wantErr: true,
		},
		{
			name:    "Empty BrandName",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//log.Println(len(tt.args.brandName))
			got, err := srvs.FindByName(ctx, tt.args.brandName)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoServices.FindByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoServices.FindByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoServices_Udate(t *testing.T) {
	type args struct {
		request request.AutosUpdateRequest
	}
	ctx := context.Background()
	srvs := &AutoServices{}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Nil model AutosUpdateRequest",
			args: args{
				request: request.AutosUpdateRequest{},
			},
			wantErr: true,
		},
		{
			name: "Empty AutoID AutosUpdateRequest",
			args: args{
				request: request.AutosUpdateRequest{
					AutoID:     0,
					Brand:      "testBran",
					Model:      "TestModel",
					Year:       1999,
					EngineType: "EngineTypeTest",
					Price:      15105,
					BodyType:   "BodyTypeTest",
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid AutoID AutosUpdateRequest",
			args: args{
				request: request.AutosUpdateRequest{
					AutoID:     -1,
					Brand:      "testBran",
					Model:      "TestModel",
					Year:       1999,
					EngineType: "EngineTypeTest",
					Price:      15105,
					BodyType:   "BodyTypeTest",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := srvs.Udate(ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("AutoServices.Udate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAutoServices_FindById(t *testing.T) {
	type args struct {
		autoId int
	}
	ctx := context.Background()
	srvs := &AutoServices{}
	tests := []struct {
		name    string
		args    args
		want    response.AutosResponse
		wantErr bool
	}{
		{
			name:    "AutoId Cero",
			args:    args{autoId: 0},
			wantErr: true,
		},
		{
			name:    "Invalid AutoId",
			args:    args{autoId: -1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srvs.FindById(ctx, tt.args.autoId)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoServices.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoServices.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}
