package test

import (
	"testing"
	"os"
	"github.com/federicoleon/golang-testing/src/api/app"
	"github.com/mercadolibre/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	go app.StartApp()
	os.Exit(m.Run())
}
