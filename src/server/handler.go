package main

import (
	"ExampleAPI_Bigset_Thrift/src/helps"
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

const BS_PERSON = "Person"
const BS_TEAM = "Team"

type GenericServiceHandler struct {
	myGeneric.TGenericService
}

func (this *GenericServiceHandler) String(obj interface{}) string {
	return fmt.Sprintf("%s", obj)
}

func (this *GenericServiceHandler) GetBsKey(bsKey string) generic.TStringKey {
	return generic.TStringKey(bsKey)
}

// check item is exist in bigset
func (this *GenericServiceHandler) ItemIsExist(ctx context.Context, bsKey string, rooId string) (bool, error) {

	_ = ctx
	item, err := bigsetIf.BsGetItem(this.GetBsKey(bsKey), generic.TItemKey(rooId))
	if err != nil {
		return false, err
	}
	if item != nil {
		return true, nil
	}
	return false, nil
}

func (this *GenericServiceHandler) GetItemPerson(ctx context.Context, bsKey string, rootID string) (*myGeneric.TPersonResult_, error) {

	_ = ctx
	bPerson, err := bigsetIf.BsGetItem(this.GetBsKey(bsKey), generic.TItemKey(this.String(rootID)))
	if err != nil {
		log.Println("error get Item: ", err)
		r := &myGeneric.TPersonResult_{
			Error: myGeneric.TErrorCode_ITEM_NOT_EXISTED,
			Item:  nil,
		}
		return r, err
	}
	var item myGeneric.TPerson
	err = json.Unmarshal(bPerson.GetValue(), &item)
	r := &myGeneric.TPersonResult_{
		Error: myGeneric.TErrorCode_SUCCESS,
		Item:  &item,
	}
	return r, nil
}

func (this *GenericServiceHandler) GetItemsPerson(ctx context.Context, bsKey string) (*myGeneric.TPeronSetResult_, error) {
	_ = ctx
	totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey(bsKey))
	if err != nil {
		log.Println("err bigset", err)
		return nil, err
	}
	slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(bsKey), 0, int32(totalCount))
	if err != nil {
		log.Println("err bigset", err)
		return nil, err
	}
	arrayTemp := make([]*myGeneric.TPerson, 0)
	for _, v := range slice {
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

func (this *GenericServiceHandler) GetPersonsPagination(ctx context.Context, bsKey string, offset int32, limit int32) (*myGeneric.TPeronSetResult_, error) {

	_ = ctx
	slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(bsKey), offset, limit)
	if err != nil {
		log.Println("err bigset", err)
		return nil, err
	}
	arrayTemp := make([]*myGeneric.TPerson, 0)
	for _, v := range slice {
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

func (this *GenericServiceHandler) GetPersonsOfTeam(ctx context.Context, teamID string) (*myGeneric.TPeronSetResult_, error) {

	_ = ctx
	personCount, err := bigsetIf.GetTotalCount(this.GetBsKey(teamID))
	if err != nil {
		log.Println(err, "bigset error: ")
		return nil, err
	}
	personSlice, err := bigsetIf.BsGetSliceR(this.GetBsKey(teamID), 0, int32(personCount))
	if err != nil {
		log.Println(err, "bigset error: ")
		return nil, err
	}
	arrayTemp := make([]*myGeneric.TPerson, 0)
	for _, v := range personSlice {
		bPerson, err := this.GetItemPerson(ctx, BS_PERSON, string(v.GetKey()))
		if err != nil {
			log.Println(err, "bigset error: ")
		} else {
			arrayTemp = append(arrayTemp, bPerson.GetItem())
		}
	}
	r := &myGeneric.TPeronSetResult_{
		Error: myGeneric.TErrorCode_SUCCESS,
		Items: arrayTemp,
	}
	return r, nil
}

func (this *GenericServiceHandler) GetPersonsOfTeamPagination(ctx context.Context, teamID string, offset int32, limit int32) (*myGeneric.TPeronSetResult_, error) {

	_ = ctx
	personSlice, err := bigsetIf.BsGetSliceR(this.GetBsKey(teamID), offset, limit)
	if err != nil {
		log.Println("err bigset", err)
		return nil, err
	}
	arrayTemp := make([]*myGeneric.TPerson, 0)
	for _, v := range personSlice {
		bPerson, err := this.GetItemPerson(ctx, BS_PERSON, string(v.GetKey()))
		if err != nil {
			log.Println(err, "bigset error: ")
		} else {
			arrayTemp = append(arrayTemp, bPerson.GetItem())
		}
	}
	r := &myGeneric.TPeronSetResult_{
		Error: myGeneric.TErrorCode_SUCCESS,
		Items: arrayTemp,
	}
	return r, nil
}

//@param bsKey is person id
func (this *GenericServiceHandler) PutPersonToTeam(ctx context.Context, teamId string, personID string) error {

	_ = ctx
	// add person to team
	bTime, err := helps.MarshalBytes(time.Now().Format("2006-01-02"))
	if err != nil {
		log.Println(err, "format time error")
	} else {
		err = bigsetIf.BsPutItem(this.GetBsKey(teamId), &generic.TItem{
			Key:   []byte(personID),
			Value: bTime,
		})
		if err != nil {
			log.Println(err, "bigset err: ")
		}
	}
	return err
}

func (this *GenericServiceHandler) PutItemPerson(ctx context.Context, bsKey string, item *myGeneric.TPerson) error {

	_ = ctx
	bPerson, err := helps.MarshalBytes(item)
	if err != nil {
		return err
	}
	key := []byte(this.String(item.GetPersonId()))
	//put person to bigset
	err = bigsetIf.BsPutItem(this.GetBsKey(bsKey), &generic.TItem{
		Key:   key,
		Value: bPerson,
	})
	if err != nil {
		log.Println(err, "bigset error")
	} else {
		// put node personID to bigset (bsKey : teamID)
		err = this.PutPersonToTeam(ctx, item.GetTeamId(), item.GetPersonId())
		if err != nil {
			log.Println(err, "PutPerson to Team Fbigset error")
		}
	}
	return err
}

func (this *GenericServiceHandler) GetItemTeam(ctx context.Context, bsKey string, rootID string) (*myGeneric.TTeamResult_, error) {

	_ = ctx
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

func (this *GenericServiceHandler) GetPersonIsTeam(ctx context.Context, personID string) (*myGeneric.TTeamResult_, error) {

	_ = ctx
	person, err := this.GetItemPerson(ctx, BS_PERSON, personID)
	if err != nil {
		log.Println(err, "bigset error:")
		return nil, err
	} else if person != nil {
		team, err := this.GetItemTeam(ctx, BS_TEAM, person.GetItem().GetTeamId())
		if err != nil {
			log.Println(err, " get Team Item bigset error:")
			return nil, err
		}
		return team, nil
	}
	return nil, nil
}

func (this *GenericServiceHandler) GetItemsTeam(ctx context.Context, bsKey string) (*myGeneric.TTeamSetResult_, error) {

	_ = ctx
	totalCount, err := bigsetIf.GetTotalCount(this.GetBsKey(bsKey))
	if err != nil {
		return nil, err
	} else {
		slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(bsKey), 0, int32(totalCount))
		if err != nil {
			return nil, err
		}

		arrayTemp := make([]*myGeneric.TTeam, 0)
		for _, v := range slice {
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

func (this *GenericServiceHandler) GetTeamsPagination(ctx context.Context, bsKey string, offset int32, limit int32) (*myGeneric.TTeamSetResult_, error) {

	_ = ctx
	count, err := bigsetIf.GetTotalCount(this.GetBsKey(bsKey))
	if err != nil {
		log.Println("err bigset", err)
	}
	if limit <= 0 {
		limit = int32(count)
	}
	slice, err := bigsetIf.BsGetSliceR(this.GetBsKey(bsKey), offset, limit)
	if err != nil {
		return nil, err
	}

	arrayTemp := make([]*myGeneric.TTeam, 0)
	for _, v := range slice {
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

func (this *GenericServiceHandler) PutItemTeam(ctx context.Context, bsKey string, item *myGeneric.TTeam) error {

	_ = ctx
	bTeam, err := helps.MarshalBytes(item)
	if err != nil {
		return err
	}
	key := []byte(item.TeamId)
	err = bigsetIf.BsPutItem(this.GetBsKey(bsKey), &generic.TItem{
		Key:   key,
		Value: bTeam,
	})
	if err != nil {
		log.Println(err, "bigset err: ")
		return err
	}
	return nil
}

// func (this *GenericServiceHandler) PutPersonToTeam(ctx context.Context, bsKey string, personId string) error {

// 	_ = ctx
// 	// add person to team
// 	bTime, err := helps.MarshalBytes(time.Now().Format("2006-01-02"))
// 	if err != nil {
// 		log.Println(err, "format time error")
// 		return err
// 	}
// 	err = bigsetIf.BsPutItem(this.GetBsKey(bsKey), &generic.TItem{
// 		Key:   []byte(fmt.Sprintf("%s", personId)),
// 		Value: bTime,
// 	})
// 	if err != nil {
// 		log.Println(err, "put person to team bigset err: ")
// 		return err
// 	}
// 	return nil
// }

// func (this *GenericServiceHandler) PutMultiPersonsToTeam(ctx context.Context, bsKey string, personIds []string) error {

// 	_ = ctx
// 	// add person to team
// 	bTime, err := helps.MarshalBytes(time.Now().Format("2006-01-02"))
// 	if err != nil {
// 		log.Println(err, "format time error")
// 		return err
// 	}
// 	items := make([]*generic.TItem, 0)
// 	for _, v := range personIds {
// 		items = append(items, &generic.TItem{
// 			Key:   []byte(v),
// 			Value: bTime,
// 		})
// 	}
// 	return bigsetIf.BsMultiPut(this.GetBsKey(bsKey), items)
// }

func (this *GenericServiceHandler) RemoveItem(ctx context.Context, bsKey string, rooID string) error {

	_ = ctx
	err := bigsetIf.BsRemoveItem(this.GetBsKey(bsKey), generic.TItemKey(rooID))
	if err != nil {
		log.Printf("can not remove bigset of %s", rooID)
	}
	return nil

}

// func (this *GenericServiceHandler) RemoveAll(ctx context.Context, bsKey string) error {

// 	_ = ctx
// 	err := bigsetIf.RemoveAll(this.GetBsKey(bsKey))
// 	if err != nil {
// 		log.Printf("can not remove bigset of %s", bsKey)
// 	}
// 	return nil

// }
