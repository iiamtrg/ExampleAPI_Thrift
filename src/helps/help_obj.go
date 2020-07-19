package helps

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"log"
	"reflect"
)
func MarshalArrayBytes(obj interface{}) ([][]byte, [][]byte, error){
	var objs, keys [][]byte
	if reflect.ValueOf(obj).Kind().String() != "slice" {
		return nil, nil, errors.New("kind of objects doesn't match slice type")
	}
	for _, object := range obj.([]interface{}) {
		marshal, err := json.Marshal(object)
		if err != nil {
			return nil, nil, err
		}
		objs = append(objs, marshal)
		keys = append(keys, []byte(fmt.Sprintf("%v", object)))
	}

	return objs, keys, nil
}

func MarshalBytes(object interface{}) ([]byte, error) {
	if object == nil {
		return nil, errors.New("object must be not nil")
	}

	obj, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}

	return obj, nil
}


func UnMarshalArrayTItem(objects []*generic.TItem) ([]interface{}, error) {
	log.Println("method @UnMarshalArrayTItem -- begin")
	objs := make([]interface{}, 0)
	var obj interface{}
	for _, object := range objects {
		err := json.Unmarshal(object.GetValue(), &obj)
		log.Println(string(object.GetValue()), "-- string(object.GetValue())")
		if err != nil {
			log.Println(err.Error(), "-- err.Error() helps/unmarshal_obj_help.go:18")
			return make([]interface{}, 0), err
		}

		objs = append(objs, obj)
	}

	log.Println("method @UnMarshalArrayTItem -- end")
	return objs, nil
}

func UnMarshalArrayTItemToStringKey(objects []*generic.TItem) []string {
	log.Println("method @UnMarshalArrayTItemToStringKey -- begin")
	objs := make([]string, 0)
	for _, object := range objects {
		objs = append(objs, string(object.GetKey()))
	}

	log.Println("method @UnMarshalArrayTItemToStringKey -- end")
	return objs
}

func UnMarshalArrayTItemToStringVal(objects []*generic.TItem) []string {
	log.Println("method @UnMarshalArrayTItemToStringVal -- begin")
	objs := make([]string, 0)
	for _, object := range objects {
		objs = append(objs, string(object.GetKey()))
	}

	log.Println("method @UnMarshalArrayTItemToStringVal -- end")
	return objs
}

func UnMarshalTItem(object *generic.TItem) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(object.GetValue(), &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func UnMarshalBytes(bytes []byte) (interface{}, error) {
	var obj interface{}

	err := json.Unmarshal(bytes, &obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
