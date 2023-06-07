package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"sesi1-crud/entity"
	"sesi1-crud/repository"
	"sesi1-crud/utils/database"
	"testing"
)

type UsecaseMock struct {
	mock mock.Mock
}

func TestUsecase_GetUserByID(t *testing.T) {
	db := database.GetDatabase()
	fmt.Println("OKe")
	uc := Usecase{
		userRepo: repository.NewRepoUser(db),
	}
	tests := []struct {
		name    string
		args    int
		want    *entity.Users
		wantErr bool
	}{
		{
			name: "Success",
			args: 1,
			want: &entity.Users{
				Name:     "latiflijdsbgiubdspgiab",
				Email:    "kubonagisa@gmail.com",
				Password: "1234",
			},
			wantErr: false,
		},
		{
			name:    "Failed get User because not found",
			args:    100,
			want:    &entity.Users{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := uc.GetUserByID(tt.args)
			assert.Equal(t, tt.want.Name, got.Name)

			//if (err != nil) != tt.wantErr {
			//	t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			//}
		})
	}

}
