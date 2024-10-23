package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGenerate(t *testing.T) {
	pass := "password"
	type args struct {
		raw string
	}
	tests := []struct {
		description string
		args        args
		want        bool
	}{
		{
			description: "Generate password OK",
			args: args{
				raw: pass,
			},
			want: true,
		},
		{
			description: "Generate password Fail",
			args:        args{},
			want:        false,
		},
	}
	for _, test := range tests {
		result := true
		hash := Generate(test.args.raw)
		err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
		if err != nil {
			result = false
		}
		assert.Equalf(t, test.want, result, test.description)
	}
}

func TestVerify(t *testing.T) {
	pass := "password"
	hash, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	type args struct {
		hash string
		raw  string
	}
	tests := []struct {
		description string
		args        args
		wantErr     bool
	}{
		{
			description: "Generate password OK",
			args: args{
				raw:  pass,
				hash: string(hash),
			},
			wantErr: false,
		},
		{
			description: "Generate password OK",
			args: args{
				raw:  "pass",
				hash: string(hash),
			},
			wantErr: true,
		},
		{
			description: "Generate password OK",
			args: args{
				raw:  pass,
				hash: "string(hash)",
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		err := Verify(test.args.hash, test.args.raw)

		if (err != nil) {
			assert.Equalf(t, test.wantErr, true, test.description )
			continue
		}
		assert.Equalf(t, test.wantErr, false, test.description )

	}
}
