package test

import (
	"bytes"
	_ "mutants-project/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

func TestStats(t *testing.T) {
	r, _ := http.NewRequest("GET", "/stats", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

func TestMutant(t *testing.T) {
	var jsonStr = []byte(`{ "dna": [ "ATGC", "CAGT", "TTTT", "AGAA" ] }`)
	r, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestMutantHumanCase(t *testing.T) {
	var jsonStr = []byte(`{ "dna": [ "ATGC", "CAGT", "TGTT", "AGAA" ] }`)
	r, _ := http.NewRequest("POST", "/mutant", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 403", func() {
			So(w.Code, ShouldEqual, 403)
		})
	})
}
