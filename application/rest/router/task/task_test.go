package task

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	app "github.com/becosuke/tasks-api/application/grpc/server/task"
	"github.com/becosuke/tasks-api/config"
	"github.com/becosuke/tasks-api/domain/entity/common"
	message "github.com/becosuke/tasks-api/protogen/message/task"
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
	app.Register(server)

	go func() {
		server.Serve(listen)
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

func TestGetDocument(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/task/1", nil)
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

func TestCreate(t *testing.T) {
	conf := config.GetConfig()
	if conf.TasksEnv == common.ENV_PRODUCTION {
		t.Log("skip test")
		return
	}

	param := `{"title": "created", "list_id": 1}`
	req, err := http.NewRequest(http.MethodPost, "/v1/task", bytes.NewBufferString(param))
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

func createDocument() (*message.Document, error) {
	conf := config.GetConfig()
	if conf.TasksEnv == common.ENV_PRODUCTION {
		return nil, errors.New("skip test")
	}

	param := `{"title": "created", "list_id": 1}`
	req, err := http.NewRequest(http.MethodPost, "/v1/task", bytes.NewBufferString(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		return nil, err
	}

	created := &message.Document{}
	err = jsonpb.Unmarshal(w.Body, created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func TestUpdate(t *testing.T) {
	conf := config.GetConfig()
	if conf.TasksEnv == common.ENV_PRODUCTION {
		t.Log("skip test")
		return
	}

	created, err := createDocument()
	if err != nil {
		t.Error(err)
	}

	param := `{"title": "updated", "list_id": 1}`
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/v1/task/%d", created.Id), bytes.NewBufferString(param))
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
	conf := config.GetConfig()
	if conf.TasksEnv == common.ENV_PRODUCTION {
		t.Log("skip test")
		return
	}

	created, err := createDocument()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/task/%d", created.Id), nil)
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
	req, err := http.NewRequest(http.MethodGet, "/v1/tasks/document/1,2,3", nil)
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
	req, err := http.NewRequest(http.MethodGet, "/v1/tasks/all?limit=10", nil)
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
	req, err := http.NewRequest(http.MethodGet, "/v1/tasks/all/count", nil)
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

func TestGetDocumentsByList(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/tasks/list/1?limit=10", nil)
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

func TestGetCountByList(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/v1/tasks/list/1/count", nil)
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
