package context

import (
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	app "github.com/becosuke/tasks-api/application/grpc/server/context"
	"github.com/becosuke/tasks-api/config"
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
	if req, err = http.NewRequest(http.MethodGet, "/v1/context/1", nil); err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("code: %d", w.Code)
	}

	t.Log(w.Body.String())
}
