package repository

import (
	"AutomobilesAPI/models"
	"context"
	"reflect"
	"testing"
)

func TestAutosRepository_FindById(t *testing.T) {
	type args struct {
		ctx    context.Context
		autoID int
	}
	tests := []struct {
		name    string
		p       *AutosRepository
		args    args
		want    models.Autos
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.FindById(tt.args.ctx, tt.args.autoID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutosRepository.FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutosRepository.FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}
