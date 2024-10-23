package utils

import (
	"testing"

	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/types"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	payloadTodo := types.CreateDTO{
		Task: "prueba",
	}
	payloadCheckTodo := types.CheckTodoDTO{
		Completed: false,
	}
	payloadLogin := types.LoginDTO {
		Email: "email@email.com",
		Password: "password",
	}
	payloadLoginWithPasswordLessThanSixDigits := types.LoginDTO {
		Email: "email@email.com",
		Password: "pas",
	}
	fiberErrorWithPasswordLessThanSixDigits := &fiber.Error{
		Code: 400,
		Message: "`Password` with value `pas` doesn't satisfy the `password` constraint",
	}
	payloadSignUp := types.SignupDTO {
		LoginDTO: payloadLogin,
		Name: "name",
	}
	payloadErrorBothEmpty := types.SignupDTO{}
	fiberErrorBothEmpty := &fiber.Error{
		Code: 400,
		Message: "`Email` with value `` doesn't satisfy the `required` constraint,`Password` with value `` doesn't satisfy the `password` constraint,`Name` with value `` doesn't satisfy the `required` constraint",
	}
	payloadErrorAnEmptyField := types.LoginDTO{
		Password: "password",
	}
	fiberErrorAnEmptyField := &fiber.Error{
		Code: 400,
		Message: "`Email` with value `` doesn't satisfy the `required` constraint",
	}

	type args struct {
		payload interface{}
	}
	tests := []struct {
		description string
		args args
		want *fiber.Error
	}{
		{
			description: "Create task",
			args: args{
				payload: payloadTodo,
			},
			want: nil,
		},
		{
			description: "Create task",
			args: args{
				payload: payloadCheckTodo,
			},
			want: nil,
		},
		{
			description: "Login User",
			args: args{
				payload: payloadLogin,
			},
			want: nil,
		},
		{
			description: "SignUp User",
			args: args{
				payload: payloadSignUp,
			},
			want: nil,
		},
		{
			description: "SignUp User",
			args: args{
				payload: payloadErrorBothEmpty,
			},
			want: fiberErrorBothEmpty,
		},
		{
			description: "SignUp User",
			args: args{
				payload: payloadLoginWithPasswordLessThanSixDigits,
			},
			want: fiberErrorWithPasswordLessThanSixDigits,
		},
		{
			description: "SignUp User",
			args: args{
				payload: payloadErrorAnEmptyField,
			},
			want: fiberErrorAnEmptyField,
		},
	}
	for _, test := range tests {

		result := Validate(test.args.payload)

		assert.Equalf(t, test.want, result, test.description)
	}
}
