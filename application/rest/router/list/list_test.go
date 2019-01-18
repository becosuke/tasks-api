package list

import (
	"bytes"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	app "github.com/becosuke/tasks-api/application/grpc/server/list"
	"github.com/becosuke/tasks-api/config"
	"github.com/becosuke/tasks-api/domain/entity/common"
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
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, "/v1/list/1", nil); err != nil {
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
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, "/v1/lists/document/1,2,3", nil); err != nil {
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
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, "/v1/lists/all", nil); err != nil {
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
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, "/v1/lists/all/count", nil); err != nil {
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

	var req *http.Request
	var err error
	param := []byte(`{"title": "created"}`)
	if req, err = http.NewRequest(http.MethodPost, "/v1/list", bytes.NewBuffer(param)); err != nil {
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
	conf := config.GetConfig()
	if conf.TasksEnv == common.ENV_PRODUCTION {
		t.Log("skip test")
		return
	}

	var req *http.Request
	var err error
	param := []byte(`{"title": "updated"}`)
	if req, err = http.NewRequest(http.MethodPut, "/v1/list/1", bytes.NewBuffer(param)); err != nil {
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

	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodDelete, "/v1/list/1", nil); err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}
