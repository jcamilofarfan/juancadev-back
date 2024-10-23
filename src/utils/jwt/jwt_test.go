package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGenerate(t *testing.T) {
	payload := TokenPayload{
		ID: 1,
	}
	payloadErr := TokenPayload{}
	type args struct {
		payload *TokenPayload
	}
	tests := []struct {
		description string
		args        args
		wantErr     bool
	}{
		{
			description: "Generate Token Succes",
			args: args{
				payload: &payload,
			},
			wantErr: false,
		},
		{
			description: "Generate Token Succes",
			args: args{
				payload: &payloadErr,
			},
			wantErr: true,
		},
	}
	for _, test := range tests {
		_, errorToken := Generate(test.args.payload)

		if errorToken != nil {
			assert.Equalf(t, test.wantErr, true, test.description)
			continue
		}
		assert.Equalf(t, test.wantErr, false, test.description)
	}
}

func TestVerify(t *testing.T) {
	payload := TokenPayload{
		ID: 1,
	}
	token, _ := Generate(&payload)
	type args struct {
		token string
	}
	tests := []struct {
		description    string
		args    args
		want    *TokenPayload
		wantErr bool
	}{
		{
			description: "Generate Token Succes",
			args: args{
				token: token,
			},
			want: &payload,
			wantErr: false,
		},
		{
			description: "Generate Token Succes",
			args: args{
				token: "token",
			},
			want: &payload,
			wantErr: true,
		},
	}
	for _, test := range tests {

		tokenPayload, err :=Verify(test.args.token)

		if (err != nil){
			assert.Equalf(t, test.wantErr, true, test.description)
			continue
		}

		assert.Equalf(t, test.wantErr, false, test.description)
		assert.Equalf(t, test.want, tokenPayload, test.description)

	}
}
