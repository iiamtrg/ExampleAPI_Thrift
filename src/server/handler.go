package main

import (
	"ExampleAPI_Bigset_Thrift/src/helps"
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"context"
	"encoding/json"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
)



type GenericServiceHandler struct {
	myGeneric.TGenericService
}

func (this *GenericServiceHandler) String(obj interface{}) string{
	return fmt.Sprintf("%s", obj)
}

func (this *GenericServiceHandler) GetBsKey(bsKey string) generic.TStringKey {
	return generic.TStringKey(fmt.Sprintf("%s", bsKey))
}

func (this *GenericServiceHandler) GetItemPerson(ctx context.Context, bsKey string, rootID string) (*myGeneric.TPersonResult_, error){

	_ = ctx
	bPerson, err := bigsetIf.BsGetItem(this.GetBsKey(bsKey), generic.TItemKey(this.String(rootID)))
	if err != nil {
		log.Println("error get Item: ", err)
		r := &myGeneric.TPersonResult_{
			Error: myGeneric.TErrorCode_ITEM_NOT_EXISTED,
			Item: nil,
		}
		return r, err
	}
	var item myGeneric.TPerson
	err = json.Unmarshal(bPerson.GetValue(), &item)
	r := &myGeneric.TPersonResult_{
		Error: myGeneric.TErrorCode_SUCCESS,
		Item: &item,
	}
	return r, nil
}

func (this *GenericServiceHandler) GetItemsPerson(ctx context.Context, bsKey string) (*myGeneric.TPeronSetResult_, error){
	fmt.Println("ok")
	_ = ctx
	totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey(bsKey))
	if err != nil {
		log.Println("err bigset", err)
		return nil, err
	} else {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(bsKey), 0, int32(totalCount))
		if err != nil{
			log.Println("err bigset", err)

			return nil,err
		}
		arrayTemp := make([]*myGeneric.TPerson, 0)
		for _,v := range slice {
			var temp myGeneric.TPerson
			_ = json.Unmarshal(v.GetValue(), &temp)
			arrayTemp = append(arrayTemp, &temp)
		}
		r := &myGeneric.TPeronSetResult_{
			Error: myGeneric.TErrorCode_SUCCESS,
			Items: arrayTemp,
		}
		return r, nil
	}
}

func (this *GenericServiceHandler) PutItemPerson(ctx context.Context, bsKey string, item *myGeneric.TPerson) error {

	_ = ctx

	bPerson, _ , err := helps.MarshalBytes(item)
	key := []byte(item.PersonId)
	if err != nil {
		return err
	}

	return bigsetIf.BsPutItem(this.GetBsKey(bsKey), &generic.TItem{
		Key: key,
		Value: bPerson,
	})
}

func (this *GenericServiceHandler) GetItemTeam(ctx context.Context, bsKey string, rootID string) (*myGeneric.TTeamResult_, error){
	_ = ctx
	fmt.Println(rootID)
	bTeam, err := bigsetIf.BsGetItem(this.GetBsKey(bsKey), generic.TItemKey(this.String(rootID)))
	r := &myGeneric.TTeamResult_{}
	if err != nil {
		log.Println("error get Item: ", err)
		r.Error = myGeneric.TErrorCode_ITEM_NOT_EXISTED
		r.Item = nil
		return r, err
	}
	item := &myGeneric.TTeam{}
	_ = json.Unmarshal(bTeam.GetValue(), item)
	r.Error = myGeneric.TErrorCode_SUCCESS
	r.Item = item
	return r, nil
}

func (this *GenericServiceHandler) GetItemsTeam(ctx context.Context, bsKey string) (*myGeneric.TTeamSetResult_, error){

	_ = ctx
	totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey(bsKey))
	if err != nil {
		return nil, err
	} else {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(bsKey), 0, int32(totalCount))
		if err != nil{
			return nil,err
		}

		arrayTemp := make([]*myGeneric.TTeam, 0)
		for _,v := range slice {
			var temp myGeneric.TTeam
			_ = json.Unmarshal(v.GetValue(), &temp)
			arrayTemp = append(arrayTemp, &temp)
		}
		r := &myGeneric.TTeamSetResult_{
			Error: myGeneric.TErrorCode_SUCCESS,
			Items: arrayTemp,
		}
		return r, nil
	}
}

func (this *GenericServiceHandler) PutItemTeam(ctx context.Context, bsKey string, item *myGeneric.TTeam) error {

	_ = ctx

	bTeam, _, err := helps.MarshalBytes(item)
	if err != nil {
		return err
	}
	key := []byte(item.TeamId)
	return bigsetIf.BsPutItem(this.GetBsKey(bsKey), &generic.TItem{
		Key: key,
		Value: bTeam,
	})
}

func (this *GenericServiceHandler) RemoveItem(ctx context.Context, bsKey string, rooID string) error {

	_ = ctx
	return bigsetIf.BsRemoveItem(this.GetBsKey(bsKey), generic.TItemKey(this.String(rooID)))
}

func (this *GenericServiceHandler) RemoveAll(ctx context.Context, bsKey string) error {
	_ = ctx
	return bigsetIf.RemoveAll(this.GetBsKey(bsKey))
}

