package test

import (
	_ "ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/routers"
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	log "github.com/astaxie/beego/logs"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// test get all team
func TestGetTeams(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/team", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n%s", w.Code, w.Body.String())

	var response myGeneric.TTeamSetResult_
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Team Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

// test get team by teamid
func TestGetTeamById(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/team/t-12", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n%s", w.Code, w.Body.String())

	var response myGeneric.TTeamResult_
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Team Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

// test create new team
func TestCreateTeam(t *testing.T) {
	var jsonStr = []byte(`{"teamId":"t-2","teamName":"Mobile","teamAddress":"HN"}`)

	req, err := http.NewRequest("POST", "/v1/team", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Team Endpoint\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

// test update a team
func TestPutTeam(t *testing.T) {
	var jsonStr = []byte(`{"teamId":"t-2","teamName":"Mobile_Put","teamAddress":"HN2"}`)

	req, err := http.NewRequest("PUT", "/v1/team/t-2", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n%s", w.Code)
	Convey("Subject: Test Team Endpoint\n", t, func() {
		Convey("Status Code Should Be 204", func() {
			So(w.Code, ShouldEqual, 204)
		})
	})
}

// test delete a team
func TestDeleteTeam(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/v1/team/t-2", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n", w.Code)

	Convey("Subject: Test Person Endpoint\n", t, func() {
		Convey("Status Code Should Be 204", func() {
			So(w.Code, ShouldEqual, 204)
		})

	})
}

// Get the person's team list
func TestGetPersonIsTeam(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/team/person/p-1", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n%s", w.Code, w.Body.String())

	var response myGeneric.TTeamResult_
	json.Unmarshal(w.Body.Bytes(), &response)

	Convey("Subject: Test Team Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})

	})
}

// test post person to team
func TestPostPersonToTeam(t *testing.T) {

	form := url.Values{}
	form.Set("personId", "p-1")
	req, err := http.NewRequest("POST", "/v1/team/t-12/person", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.PostForm = form
	req.Header.Set("Content-Type", "multipart/form-data")
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	log.Trace("testing", "TestTeam", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Team Endpoint\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})

	})
}
