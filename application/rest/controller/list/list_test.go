package list

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	controller "github.com/becosuke/tasks-api/application/grpc/controller/list"
	"github.com/becosuke/tasks-api/config"
	entity "github.com/becosuke/tasks-api/domain/entity/list"
	service "github.com/becosuke/tasks-api/domain/service/list"
)

var server *grpc.Server
var mux *runtime.ServeMux

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		log.Fatal(err)
	}
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func setup() error {
	conf := config.GetConfig()

	listen, err := net.Listen("tcp", conf.GrpcAddr)
	if err != nil {
		return err
	}
	server = grpc.NewServer()
	controller.Register(server)

	go func() {
		_ = server.Serve(listen)
	}()

	router := NewRouter()
	if err = router.Setup(); err != nil {
		return err
	}

	mux = router.GetMux()

	return nil
}

func teardown() {
	server.GracefulStop()
}

func create(title string) (*entity.Document, error) {
	if !config.IsLocal() {
		return nil, errors.New("skip test")
	}

	document, err := service.Create(title)
	if err != nil {
		return &entity.Document{}, err
	}

	bs, _ := json.Marshal(document.Message())
	log.Print(string(bs))
	return document, nil
}

func TestCreate(t *testing.T) {
	if !config.IsLocal() {
		t.Log("skip test")
		return
	}

	param := `{"title": "created"}`
	req, err := http.NewRequest(http.MethodPost, "/v1/list", bytes.NewBufferString(param))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}

func TestUpdate(t *testing.T) {
	if !config.IsLocal() {
		t.Log("skip test")
		return
	}

	created, err := create("update")
	if err != nil {
		t.Error(err)
	}

	param := `{"title": "updated"}`
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/list/%d", created.Id), bytes.NewBufferString(param))
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}

func TestDelete(t *testing.T) {
	if !config.IsLocal() {
		t.Log("skip test")
		return
	}

	created, err := create("delete")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/list/%d", created.Id), nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}

func TestGetDocument(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/list/1", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}

func TestGetDocuments(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/lists/document/1,2,3", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}

func TestGetDocumentsAll(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/lists/all?limit=10", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}

func TestGetCountAll(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/lists/all/count", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}
