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
type PersonController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all person
// @Param	offset	query	int	false
// @Param	limit	query	int false
// @Success 200 {object} myGeneric.TPerson
// @router / [get]
func (p *PersonController) Get() {

	defer p.ServeJSON()
	sv := &models.PersonClient{}
	off := p.GetString("offset")
	limit := p.GetString("limit")
	if off == "" && limit == "" {
		result, err := sv.GetItemsAll()
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		} else {
			p.Ctx.ResponseWriter.WriteHeader(200)
			p.Data["json"] = result
		}
	} else if off != "" && limit == "" {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		result, err := sv.GetItemsPagination(int32(offInt), 0)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		}
		p.Data["json"] = result
		return
	} else if off == "" && limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		result, err := sv.GetItemsPagination(0, int32(limitInt))
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		}
		p.Data["json"] = result
		return
	} else {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(400)
			return
		}
		result, err := sv.GetItemsPagination(int32(offInt), int32(limitInt))
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		}
		p.Data["json"] = result
		return
	}

}

// @Title GetPerson
// @Description get person by id
// @Param	uid 	path	string	true
// @Success 200 {object} myGeneric.TPerson
// @Failure 403 {string} uid is not int
// @router /:uid [get]
func (p *PersonController) GetById() {

	defer p.ServeJSON()
	sv := &models.PersonClient{}
	result, err := sv.GetItemById(p.GetString(":uid"))
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(404)
		p.Data["json"] = "id is not exist"
	}
	p.Data["json"] = result
}

// @Title Get persons of team
// @Description  Get persons of team
// @Param	uid 	path	string	true
// @Success 200 {object} teams
// @router /team/:uid [get]
func (p *PersonController) GetPersonOfTeam() {

	defer p.ServeJSON()
	teamId := p.GetString(":uid")
	sv := &models.PersonClient{}
	result, err := sv.GetPersonsOfTeam(teamId)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(404)
		return
	}
	p.Data["json"] = result
}

// @Title CreatePerson
// @Description create new person
// @Param body	body	myGeneric.TPerson	true
// @Success	201 {string} id
// @router / [post]
func (p *PersonController) Post() {
	defer p.ServeJSON()
	person := myGeneric.TPerson{}
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &person)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(400)
	} else {
		matched, err := regexp.Match(`^p-\d+$`, []byte(person.GetPersonId()))
		if err != nil || !matched {
			p.Ctx.ResponseWriter.WriteHeader(400)
			obj := make(map[string]string, 0)
			obj["code"] = "400"
			obj["message"] = "personID is not valid. Pattern: p-[0-9]+"
			p.Data["json"] = obj
			return
		}
		sv := &models.PersonClient{}
		err = sv.PutItem(&person)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		}
		p.Ctx.ResponseWriter.WriteHeader(201)
		p.Data["json"] = "create success"
	}
}

// @Title create person's team
// @Description create person's team
// @Param uid	path	string	true
// @Param teamId	body	string	true
// @Success	200 {string} string
// @router /:uid/team/ [post]
func (p *PersonController) PostTeam() {

	defer p.ServeJSON()
	uid := p.GetString(":uid")
	teamId := p.GetString("teamId")
	sv := models.PersonClient{}
	err := sv.PutPersonIsTeam(uid, teamId)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(400)
		return
	}
	p.Ctx.ResponseWriter.WriteHeader(201)
	p.Data["json"] = "create success"
	return
}

// @Title UpdatePerson
// @Description update person
// @Param uid 	path	string	true
// @Param body	body	myGeneric.TPerson true
// @Success 201 | 204
// @router /:uid [put]
func (p *PersonController) Put() {

	defer p.ServeJSON()
	person := &myGeneric.TPerson{}
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &person)
	if err != nil || strings.Compare(person.GetPersonId(), p.GetString(":uid")) != 0 {
		p.Ctx.ResponseWriter.WriteHeader(400)
	} else {
		sv := &models.PersonClient{}
		_, err1 := sv.GetItemById(person.GetPersonId())
		err2 := sv.PutItem(person)
		if err2 != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		} else if err1 != nil {
			p.Ctx.ResponseWriter.WriteHeader(201)
			p.Ctx.ResponseWriter.Header().Set("location", fmt.Sprintf("%s/%s/%s", p.Ctx.Input.Host(), p.Ctx.Input.URL(), person.GetPersonId()))
		} else {
			p.Ctx.ResponseWriter.WriteHeader(204)
		}
	}
}

// @Title DeletePerson
// @Description delete person
// @Param	uid		path	string	true
// @Success	204
// @router /:uid [delete]
func (p *PersonController) Delete() {

	defer p.ServeJSON()
	sv := &models.PersonClient{}
	uid := p.GetString(":uid")
	_, err := sv.GetItemById(uid)
	if err != nil {
		p.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		err = sv.RemoveItem(uid)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
		}
		p.Ctx.ResponseWriter.WriteHeader(204)
	}
}
