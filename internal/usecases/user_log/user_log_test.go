package user_log

import (
	"context"
	"github.com/nocturna-ta/ums/internal/domain/repository"
	"github.com/nocturna-ta/ums/internal/usecases/response"
	"reflect"
	"testing"
)

func TestModule_GetAllUserLog(t *testing.T) {
	type fields struct {
		userLogRepo repository.UserLogRepository
	}
	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []response.UserLogResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Module{
				userLogRepo: tt.fields.userLogRepo,
			}
			got, err := m.GetAllUserLog(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUserLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUserLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}
