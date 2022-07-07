package user

import (
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/pkg/constant"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/repository"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	Update(ctx context.Context, ID uint, payload *dto.UpdatePasswordRequest) (string, error)
}

func NewService(f *factory.Factory) *service {
	return &service{f.UserRepository}
}

func (s *service) Update(ctx context.Context, ID uint, payload *dto.UpdatePasswordRequest) (string, error) {

	// check old password
	data_user, err := s.UserRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return "", res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data_user.Password), []byte(payload.OldPassword)); err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.OldPasswordIncorrect, err)
	}

	var data = make(map[string]interface{})
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	data["password"] = string(bytes)

	err = s.UserRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}