package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils/jwt"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	payload := jwt.TokenPayload{
		ID: 1,
	}
	tk, _ := jwt.Generate(&payload)
	token := "Bearer " + tk
	bodySingUP := []byte(`{
		"email": "email@email.com",
		"password": "password",
		"name": "name"
	}`)
	bodyLogin := []byte(`{
		"email": "email@email.com",
		"password": "password"
	}`)
	bodyCreateTodo := []byte(`{
		"task": "first Task"
	}`)
	bodyUpdateTodo := []byte(`{
		"task": "create mock go"
	}`)
	bodyCheckTodo := []byte(`{
		"completed": true
	}`)

	tests := []struct {
		description     string
		route           string
		method          string
		expectedError   bool
		body            io.Reader
		expectedCode    int
		expectedBody    string
		dontcomparebody bool
		token           bool
	}{
		{
			description:   "index route",
			route:         "/",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "OK",
		},
		{
			description:   "health route",
			route:         "/health",
			method:        "GET",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "healthy",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			method:        "POST",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  `{"statusCode":404,"error":"Cannot POST /i-dont-exist"}`,
		},
		{
			description:     "sing Up succes",
			route:           "/auth/signup",
			method:          "POST",
			expectedError:   false,
			body:            bytes.NewBuffer(bodySingUP),
			expectedCode:    200,
			dontcomparebody: true,
		},
		{
			description:   "sing Up Error",
			route:         "/auth/signup",
			method:        "POST",
			expectedError: false,
			body:          bytes.NewBuffer(bodySingUP),
			expectedCode:  409,
			expectedBody:  `{"statusCode":409,"error":"Email already exists"}`,
		},
		{
			description:     "Login succes",
			route:           "/auth/login",
			method:          "POST",
			expectedError:   false,
			body:            bytes.NewBuffer(bodyLogin),
			expectedCode:    200,
			dontcomparebody: true,
		},
		{
			description:   "Login Error",
			route:         "/auth/login",
			method:        "POST",
			expectedError: false,
			body: bytes.NewBuffer([]byte(`{
				"email": "email@email.com",
				"password": "passworderror"
			}`)),
			expectedCode: 401,
			expectedBody: `{"statusCode":401,"error":"Invalid email or password"}`,
		},
		{
			description:     "Create Todo",
			route:           "/todo/create",
			method:          "POST",
			expectedError:   false,
			body:            bytes.NewBuffer(bodyCreateTodo),
			expectedCode:    200,
			expectedBody:    `{"todo":{"id":1,"task":"first Task","completed":false}}`,
			token:           true,
		},
		{
			description:     "Get a Todo",
			route:           "/todo/1",
			method:          "GET",
			expectedError:   false,
			expectedCode:    200,
			expectedBody:    `{"todo":{"id":1,"task":"first Task","completed":false}}`,
			token:           true,
		},
		{
			description:     "Update Task",
			route:           "/todo/1",
			method:          "PATCH",
			expectedError:   false,
			body:            bytes.NewBuffer(bodyUpdateTodo),
			expectedCode:    200,
			expectedBody:    `{"message":"Todo successfully updated"}`,
			token:           true,
		},
		{
			description:     "Create Todo",
			route:           "/todo/1/check",
			method:          "PATCH",
			expectedError:   false,
			body:            bytes.NewBuffer(bodyCheckTodo),
			expectedCode:    200,
			expectedBody:    `{"message":"Todo successfully updated"}`,
			token:           true,
		},
		{
			description:     "Get Todos",
			route:           "/todo/list",
			method:          "GET",
			expectedError:   false,
			expectedCode:    200,
			expectedBody:    `{"todos":[{"id":1,"task":"create mock go","completed":true}]}`,
			token:           true,
		},
		{
			description:     "Delete Todo",
			route:           "/todo/1",
			method:          "DELETE",
			expectedError:   false,
			body:            bytes.NewBuffer(bodyCheckTodo),
			expectedCode:    200,
			expectedBody:    `{"message":"Todo successfully deleted"}`,
			token:           true,
		},
	}

	app := Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			test.body,
		)
		req.Header.Add("Content-Type", "application/json")
		if test.token {
			req.Header.Add("Authorization", token)
		}

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		if test.dontcomparebody {
			continue
		}

		body, err := ioutil.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
