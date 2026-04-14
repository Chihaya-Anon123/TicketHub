package service

import (
	"strings"
	"unicode"

	"github.com/Chihaya-Anon123/TicketHub/internal/code"
	"github.com/Chihaya-Anon123/TicketHub/internal/dao"
	"github.com/Chihaya-Anon123/TicketHub/internal/errs"
	"github.com/Chihaya-Anon123/TicketHub/internal/model"
	"github.com/Chihaya-Anon123/TicketHub/internal/utils"
)

type RegisterInput struct {
	Username string
	Email    string
	Password string
}

type RegisterOutput struct {
	ID       uint
	Username string
	Email    string
	Status   uint8
}

// 用户注册
func Register(input RegisterInput) (*RegisterOutput, error) {
	//检验用户名
	if input.Username == "" {
		return nil, errs.New(code.CodeInvalidParams, "username should not be empty")
	}
	if strings.IndexFunc(input.Username, unicode.IsSpace) != -1 {
		return nil, errs.New(code.CodeInvalidParams, "username should not has space")
	}
	existUser1, err := dao.GetUserByUsername(input.Username)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if existUser1 != nil {
		return nil, errs.New(code.CodeInvalidParams, "username already exists")
	}

	//检验邮箱
	if input.Email == "" {
		return nil, errs.New(code.CodeInvalidParams, "email should not be empty")
	}
	existUser2, err := dao.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errs.ErrDBError
	}
	if existUser2 != nil {
		return nil, errs.New(code.CodeInvalidParams, "email already exists")
	}

	//检验密码
	if input.Password == "" {
		return nil, errs.New(code.CodeInvalidParams, "password should not be empty")
	}
	if strings.IndexFunc(input.Password, unicode.IsSpace) != -1 {
		return nil, errs.New(code.CodeInvalidParams, "password should not has space")
	}
	if len(input.Password) < 6 || len(input.Password) > 20 {
		return nil, errs.New(code.CodeInvalidParams, "length of password should be in range 6-20")
	}

	hashedpassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, errs.ErrInternalServer
	}

	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedpassword,
		Status:   1,
	}

	if err := dao.CreateUser(user); err != nil {
		return nil, errs.ErrDBError
	}

	return &RegisterOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Status:   user.Status,
	}, nil
}
