package httphandler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"httphandler/src"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type MockUserService struct {
	RegisterFunc    func(user httphandler.User) (string, error)
	UsersRegistered []httphandler.User
}

func (m *MockUserService) Register(user httphandler.User) (insertedID string, err error) {
	m.UsersRegistered = append(m.UsersRegistered, user)
	return m.RegisterFunc(user)
}

func userToJSON(user httphandler.User) io.Reader {
	arrayOfBytes, err := json.Marshal(user)
	if err != nil {
		return nil
	}
	return bytes.NewReader(arrayOfBytes)
}

func assertStatus(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d sent to stdout but expected %d", got, want)
	}
}

func TestRegisterUser(t *testing.T) {
	t.Run("can register valid users", func(t *testing.T) {
		user := httphandler.User{Name: "CJ"}
		expectedInsertedID := "whatever"
		service := &MockUserService{
			RegisterFunc: func(user httphandler.User) (string, error) {
				return expectedInsertedID, nil
			},
		}
		server := httphandler.NewUserServer(service)
		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
		res := httptest.NewRecorder()
		server.RegisterUser(res, req)
		assertStatus(t, res.Code, http.StatusCreated)

		if res.Body.String() != expectedInsertedID {
			t.Errorf("expected body of %q but got %q", res.Body.String(), expectedInsertedID)
		}
		if len(service.UsersRegistered) != 1 {
			t.Fatalf("expected 1 user added but got %d", len(service.UsersRegistered))
		}
		if !reflect.DeepEqual(service.UsersRegistered[0], user) {
			t.Errorf("the user registered %+v was not what was expected %+v", service.UsersRegistered[0], user)
		}
	})
	t.Run("returns 400 bad request if body is not valid user JSON", func(t *testing.T) {
		server := httphandler.NewUserServer(nil)
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("trouble will find me"))
		res := httptest.NewRecorder()
		server.RegisterUser(res, req)
		assertStatus(t, res.Code, http.StatusBadRequest)
	})
	t.Run("returns a 500 internal server error if the service fails", func(t *testing.T) {
		user := httphandler.User{Name: "CJ"}
		service := &MockUserService{
			RegisterFunc: func(user httphandler.User) (string, error) {
				return "", errors.New("couldn't add new user")
			},
		}
		server := httphandler.NewUserServer(service)
		req := httptest.NewRequest(http.MethodGet, "/", userToJSON(user))
		res := httptest.NewRecorder()
		server.RegisterUser(res, req)
		assertStatus(t, res.Code, http.StatusInternalServerError)
	})
}
