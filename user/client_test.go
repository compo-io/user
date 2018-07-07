package user

import (
	"github.com/skamenetskiy/jsonapi"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

type ClientTestSuite struct {
	suite.Suite
}

func (t *ClientTestSuite) SetupSuite() {
}

func (t *ClientTestSuite) TestCreate() {
	s1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(http.MethodPost, r.Method)
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"id":10,"login":"login1"}`))
	}))
	defer s1.Close()
	client = jsonapi.NewClient(s1.URL[7:])
	u := &User{
		Login:    "login",
		Password: "password",
	}
	u2, err := Create(u)
	t.NoError(err)
	t.NotNil(u2)
	t.Equal(uint64(10), u2.ID)
	t.Equal("login1", u2.Login)
	t.Empty(u2.Password)
	t.Empty(u2.PasswordHash)
	s2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(http.MethodPost, r.Method)
		w.WriteHeader(500)
		w.Header().Add("Content-Type", "application/json")
	}))
	client = jsonapi.NewClient(s2.URL[7:])
	u3 := &User{
		ID:       1,
		Login:    "hello",
		Password: "world",
	}
	u4, err := Create(u3)
	t.Error(err)
	t.Nil(u4)
	client = jsonapi.NewClient("bad:request:uri")
	u6, err := Create(new(User))
	t.Error(err)
	t.Nil(u6)
	u7, err := Create(nil)
	t.Error(err)
	t.Equal("user is nil", err.Error())
	t.Nil(u7)
}

func (t *ClientTestSuite) TestGetByID() {
	s1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(http.MethodGet, r.Method)
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"id":1,"login":"login","password":"password"}`))
	}))
	defer s1.Close()
	client = jsonapi.NewClient(s1.URL[7:])
	u1, err := GetByID(1)
	t.NoError(err)
	t.NotNil(u1)
	t.IsType(new(User), u1)
	t.Equal(uint64(1), u1.ID)
	t.Equal("login", u1.Login)
	t.Equal("password", u1.Password)
	t.Empty(u1.PasswordHash)
	client = jsonapi.NewClient("bad_url")
	u2, err := GetByID(1)
	t.Error(err)
	t.Nil(u2)
	s2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Equal(http.MethodGet, r.Method)
		w.Write([]byte("bad json"))
	}))
	defer s2.Close()
	client = jsonapi.NewClient(s2.URL[7:])
	u3, err := GetByID(1)
	t.Error(err)
	t.Nil(u3)
}

func (t *ClientTestSuite) TestGetByLogin() {
	s1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"id":1,"login":"login","password":"password"}`))
	}))
	defer s1.Close()
	client = jsonapi.NewClient(s1.URL[7:])
	u1, err := GetByLogin("login")
	t.NoError(err)
	t.NotNil(u1)
	t.IsType(new(User), u1)
	t.Equal(uint64(1), u1.ID)
	t.Equal("login", u1.Login)
	t.Equal("password", u1.Password)
	t.Empty(u1.PasswordHash)
	client = jsonapi.NewClient("bad_url")
	u2, err := GetByLogin("login")
	t.Error(err)
	t.Nil(u2)
	s2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bad json"))
	}))
	defer s2.Close()
	client = jsonapi.NewClient(s2.URL[7:])
	u3, err := GetByID(1)
	t.Error(err)
	t.Nil(u3)
}
