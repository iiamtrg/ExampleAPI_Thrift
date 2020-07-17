package controllers

import (
	"ExampleAPI_Bigset_Thrift/src/ExampleAPI_Thrift/models"
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	_ "ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

// Operations about Users
type TeamController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all teams
// @Success 200 {object} teams
// @router / [get]
func (this *TeamController) GetAll() {

	defer this.ServeJSON()
	sv := &models.TeamClient{}
	result, err := sv.GetItemsAll()
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = result
	}
}

// @Title GetItem
// @Description get team by id
// @Param	uid 	path	string	true
// @Success 200 {object} teams
// @Failure 403 {string} uid is not int
// @router /:uid [get]
func (this *TeamController) Get() {

	defer this.ServeJSON()
	sv := &models.TeamClient{}
	result, err := sv.GetItemById(this.GetString(":uid"))
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Data["json"] = "id is not existed"
	}
	this.Data["json"] = result
}

// @Title CreateTeam
// @Description create new team
// @Param body	body	myGeneric.TTeam	true
// @Success	201 {string} id
// @router / [post]
func (this *TeamController) Post() {
	defer this.ServeJSON()
	team := myGeneric.TTeam{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &team)
	if err != nil{
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = "data is not mapping success"
	} else {
		sv := &models.TeamClient{}
		err = sv.Put(&team)
		if err != nil {
			this.Ctx.ResponseWriter.WriteHeader(500)
		}
		this.Data["json"] = "create success"
	}
}

// @Title UpdateTeam
// @Description update team
// @Param uid 	path	string	true
// @Param body	body	myGeneric.TTeam true
// @Success 201 | 204
// @router /:uid [put]
func (this *TeamController) Put(){

	defer this.ServeJSON()
	team := &myGeneric.TTeam{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &team)
	if err != nil || strings.Compare(team.GetTeamId(), this.GetString(":uid")) != 0{
		this.Ctx.ResponseWriter.WriteHeader(400)
	} else {
		sv := &models.TeamClient{}
		_, err1 := sv.GetItemById(team.GetTeamId())
		err2 := sv.Put(team)
		if err2 != nil {
			this.Ctx.ResponseWriter.WriteHeader(500)
		} else if err1 != nil {
			this.Ctx.ResponseWriter.WriteHeader(201)
			this.Ctx.ResponseWriter.Header().Set("location", fmt.Sprintf("%s/%s/%s", this.Ctx.Input.Host(),this.Ctx.Input.URL(), team.GetTeamId()))
		} else {
			this.Ctx.ResponseWriter.WriteHeader(204)
		}
	}
}

// @Title delete team
// @Description delete team
// @Param	uid		path	string	true
// @Success	204
// @router /:uid [delete]
func (this *TeamController) Delete(){

	defer this.ServeJSON()
	sv := &models.TeamClient{}
	uid := this.GetString(":uid")
	_, err := sv.GetItemById(uid)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		err = sv.Remove(uid)
		if err != nil {
			this.Ctx.ResponseWriter.WriteHeader(500)
		}
		this.Ctx.ResponseWriter.WriteHeader(204)
	}
}