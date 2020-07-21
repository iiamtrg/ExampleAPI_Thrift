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

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

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
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		} else {
			p.Ctx.ResponseWriter.WriteHeader(200)
			p.Data["json"] = result
		}
	} else if off != "" && limit == "" {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		result, err := sv.GetItemsPagination(int32(offInt), 10)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		p.Data["json"] = result
		return
	} else if off == "" && limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		result, err := sv.GetItemsPagination(0, int32(limitInt))
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
		}
		p.Data["json"] = result
		return
	} else {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		result, err := sv.GetItemsPagination(int32(offInt), int32(limitInt))
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
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
		return
	} else {
		p.Data["json"] = result
		return
	}

}

// @Title Get persons of team
// @Description  Get persons of team
// @Param	uid 	path	string	true
// @Param	offset	query	int	false
// @Param	limit	query	int	false
// @Success 200 {object} teams
// @router /team/:uid [get]
func (p *PersonController) GetPersonOfTeam() {

	defer p.ServeJSON()
	sv := &models.PersonClient{}
	off := p.GetString("offset")
	limit := p.GetString("limit")
	teamID := p.GetString(":uid")

	if off == "" && limit == "" {
		result, err := sv.GetPersonsOfTeam(teamID)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(404)
			errJson := &Error{}
			errJson.Code = "404"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		} else {
			p.Ctx.ResponseWriter.WriteHeader(200)
			p.Data["json"] = result
			return
		}
	} else if off != "" && limit == "" {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		result, err := sv.GetPersonOfTeamPagination(teamID, int32(offInt), 0)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(404)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		p.Data["json"] = result
		return
	} else if off == "" && limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		result, err := sv.GetPersonOfTeamPagination(teamID, 0, int32(limitInt))
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(404)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		p.Data["json"] = result
		return
	} else {
		offInt, err := strconv.Atoi(off)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(403)
			errJson := &Error{}
			errJson.Code = "403"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		result, err := sv.GetPersonOfTeamPagination(teamID, int32(offInt), int32(limitInt))
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(404)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		p.Data["json"] = result
		return
	}

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
		errJson := &Error{}
		errJson.Code = "400"
		errJson.Message = err.Error()
		p.Data["json"] = errJson
		return
	} else {
		matched, err := regexp.Match(`^p-\d+$`, []byte(person.GetPersonId()))
		if err != nil || !matched {
			p.Ctx.ResponseWriter.WriteHeader(400)
			obj := make(map[string]string)
			obj["Code"] = "400"
			obj["Message"] = "personID is not valid. Pattern: p-[0-9]+"
			p.Data["json"] = obj
			return
		}
		sv := &models.PersonClient{}
		err = sv.PutItem(&person)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		p.Ctx.ResponseWriter.WriteHeader(201)
		p.Data["json"] = "create success"
	}
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
		errJson := &Error{}
		errJson.Code = "500"
		errJson.Message = err.Error()
		p.Data["json"] = errJson
		return
	} else {
		sv := &models.PersonClient{}
		_, err1 := sv.GetItemById(person.GetPersonId())
		err2 := sv.PutItem(person)
		if err2 != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		} else if err1 != nil {
			p.Ctx.ResponseWriter.WriteHeader(201)
			p.Ctx.ResponseWriter.Header().Set("location", fmt.Sprintf("%s/%s/%s", p.Ctx.Input.Host(), p.Ctx.Input.URL(), person.GetPersonId()))
			return
		} else {
			p.Ctx.ResponseWriter.WriteHeader(204)
			return
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
		errJson := &Error{}
		errJson.Code = "404"
		errJson.Message = err.Error()
		p.Data["json"] = errJson
		return
	} else {
		err = sv.RemoveItem(uid)
		if err != nil {
			p.Ctx.ResponseWriter.WriteHeader(500)
			errJson := &Error{}
			errJson.Code = "500"
			errJson.Message = err.Error()
			p.Data["json"] = errJson
			return
		}
		p.Ctx.ResponseWriter.WriteHeader(204)
	}
}
