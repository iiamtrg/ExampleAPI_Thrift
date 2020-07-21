package controllers

import (
	"ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/models"
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	_ "ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// Operations about Users
type TeamController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all person
// @Param	offset	query	int	false
// @Param	limit	query	int false
// @Success 200 {object} myGeneric.TPerson
// @router / [get]
func (t *TeamController) Get() {

	defer t.ServeJSON()
	sv := &models.TeamClient{}
	off := t.GetString("offset")
	limit := t.GetString("limit")
	if off == "" && limit == "" {
		result, err := sv.GetItemsAll()
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
		} else {
			t.Ctx.ResponseWriter.WriteHeader(200)
			t.Data["json"] = result
		}
	} else if off != "" && limit == "" {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		result, err := sv.GetItemsPagination(int32(offInt), 0)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
		}
		t.Data["json"] = result
		return
	} else if off == "" && limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		result, err := sv.GetItemsPagination(0, int32(limitInt))
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
		}
		t.Data["json"] = result
		return
	} else {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		result, err := sv.GetItemsPagination(int32(offInt), int32(limitInt))
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
		}
		t.Data["json"] = result
		return
	}

}

// @Title GetItem
// @Description get team by id
// @Param	uid 	path	string	true
// @Success 200 {object} teams
// @router /:uid [get]
func (t *TeamController) GetById() {

	defer t.ServeJSON()
	sv := &models.TeamClient{}
	result, err := sv.GetItemById(t.GetString(":uid"))
	if err != nil {
		t.Ctx.ResponseWriter.WriteHeader(404)
		errJson := &Error{}
		errJson.code = "404"
		errJson.message = err.Error()
		t.Data["json"] = errJson
		return
	}
	t.Data["json"] = result
}

// @Title get person's team
// @Description get person's team
// @Param uid	path	string	true
// @Success	200 {string} myGeneric.TTeamResult_
// @router /person/:uid [get]
func (t *TeamController) GetPersonTeam() {

	defer t.ServeJSON()
	uid := t.GetString(":uid")
	sv := &models.PersonClient{}
	result, err := sv.GetPersonTeam(uid)
	if err != nil {
		t.Ctx.ResponseWriter.WriteHeader(404)
		errJson := &Error{}
		errJson.code = "404"
		errJson.message = err.Error()
		t.Data["json"] = errJson
		return
	}
	t.Data["json"] = result
}

// @Title CreateTeam
// @Description create new team
// @Param body	body	myGeneric.TTeam	true
// @Success	201 {string} id
// @router / [post]
func (t *TeamController) Post() {
	defer t.ServeJSON()
	team := myGeneric.TTeam{}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &team)
	if err != nil {
		t.Ctx.ResponseWriter.WriteHeader(400)
		errJson := &Error{}
		errJson.code = "400"
		errJson.message = err.Error()
		t.Data["json"] = errJson
		return
	} else {
		matched, err := regexp.Match(`^t-\d+$`, []byte(team.GetTeamId()))
		if err != nil || !matched {
			t.Ctx.ResponseWriter.WriteHeader(400)
			obj := make(map[string]string, 0)
			obj["code"] = "400"
			obj["message"] = "teamID is not valid. Pattern: t-[0-9]+"
			t.Data["json"] = obj
			return
		}
		sv := &models.TeamClient{}
		err = sv.PutItem(&team)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.code = "500"
			errJson.message = err.Error()
			t.Data["json"] = errJson
			return
		}
		t.Ctx.ResponseWriter.WriteHeader(201)
		t.Data["json"] = "create success"
	}
}

// @Title UpdateTeam
// @Description update team
// @Param uid 	path	string	true
// @Param body	body	myGeneric.TTeam true
// @Success 201 | 204
// @router /:uid [put]
func (t *TeamController) Put() {

	defer t.ServeJSON()
	team := &myGeneric.TTeam{}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &team)
	if err != nil || strings.Compare(team.GetTeamId(), t.GetString(":uid")) != 0 {
		t.Ctx.ResponseWriter.WriteHeader(400)
		errJson := &Error{}
		errJson.code = "400"
		errJson.message = err.Error()
		t.Data["json"] = errJson
	} else {
		sv := &models.TeamClient{}
		_, err1 := sv.GetItemById(team.GetTeamId())
		err2 := sv.PutItem(team)
		if err2 != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.code = "500"
			errJson.message = err.Error()
			t.Data["json"] = errJson
			return
		} else if err1 != nil {
			t.Ctx.ResponseWriter.WriteHeader(201)
			t.Ctx.ResponseWriter.Header().Set("location", fmt.Sprintf("%s/%s/%s", t.Ctx.Input.Host(), t.Ctx.Input.URL(), team.GetTeamId()))
		} else {
			t.Ctx.ResponseWriter.WriteHeader(204)
		}
	}
}

// @Title delete team
// @Description delete team
// @Param	uid		path	string	true
// @Success	204
// @router /:uid [delete]
func (t *TeamController) Delete() {

	defer t.ServeJSON()
	sv := &models.TeamClient{}
	uid := t.GetString(":uid")
	_, err := sv.GetItemById(uid)
	if err != nil {
		t.Ctx.ResponseWriter.WriteHeader(404)
		errJson := &Error{}
		errJson.code = "404"
		errJson.message = err.Error()
		t.Data["json"] = errJson
		return
	} else {
		err = sv.RemoveItem(uid)
		if err != nil {
			t.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.code = "500"
			errJson.message = err.Error()
			t.Data["json"] = errJson
			return
		}
		t.Ctx.ResponseWriter.WriteHeader(204)
	}
}
