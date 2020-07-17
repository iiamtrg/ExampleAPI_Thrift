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
type PersonController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all person
// @Success 200 {object} myGeneric.TPerson
// @router / [get]
func (this *PersonController) GetAll() {

	defer this.ServeJSON()
	sv := &models.PersonClient{}
	result, err := sv.GetItemsAll()
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Data["json"] = result
	}
}

// @Title GetPerson
// @Description get person by id
// @Param	uid 	path	string	true
// @Success 200 {object} myGeneric.TPerson
// @Failure 403 {string} uid is not int
// @router /:uid [get]
func (this *PersonController) Get() {

	defer this.ServeJSON()
	sv := &models.PersonClient{}
	result, err := sv.GetItemById(this.GetString(":uid"))
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Data["json"] = "id is not existed"
	}
	this.Data["json"] = result
}

// @Title CreatePerson
// @Description create new person
// @Param body	body	myGeneric.TPerson	true
// @Success	201 {string} id
// @router / [post]
func (this *PersonController) Post() {
	defer this.ServeJSON()
	person := myGeneric.TPerson{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &person)
	if err != nil{
		this.Ctx.ResponseWriter.WriteHeader(400)
		this.Data["json"] = "data is not mapping success"
	} else {
		sv := &models.PersonClient{}
		err = sv.Put(&person)
		if err != nil {
			this.Ctx.ResponseWriter.WriteHeader(500)
		}
		this.Data["json"] = "create success"
	}
}

// @Title UpdatePerson
// @Description update person
// @Param uid 	path	string	true
// @Param body	body	myGeneric.TPerson true
// @Success 201 | 204
// @router /:uid [put]
func (this *PersonController) Put(){

	defer this.ServeJSON()
	person := &myGeneric.TPerson{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &person)
	if err != nil || strings.Compare(person.GetPersonId(), this.GetString(":uid")) != 0{
		this.Ctx.ResponseWriter.WriteHeader(400)
	} else {
		sv := &models.PersonClient{}
		_, err1 := sv.GetItemById(person.GetPersonId())
		err2 := sv.Put(person)
		if err2 != nil {
			this.Ctx.ResponseWriter.WriteHeader(500)
		} else if err1 != nil {
			this.Ctx.ResponseWriter.WriteHeader(201)
			this.Ctx.ResponseWriter.Header().Set("location", fmt.Sprintf("%s/%s/%s", this.Ctx.Input.Host(),this.Ctx.Input.URL(), person.GetPersonId()))
		} else {
			this.Ctx.ResponseWriter.WriteHeader(204)
		}
	}
}

// @Title DeletePerson
// @Description delete person
// @Param	uid		path	string	true
// @Success	204
// @router /:uid [delete]
func (this *PersonController) Delete(){

	defer this.ServeJSON()
	sv := &models.PersonClient{}
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