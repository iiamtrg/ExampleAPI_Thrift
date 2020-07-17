

package models

import (
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"os"
	"time"
)

const BS_PERSON = "Person"


type PersonClient struct {
	TransportFactory thrift.TTransportFactory
	ProtocolFactory  thrift.TProtocolFactory
	Transport        thrift.TTransport
}

func (this *PersonClient) InitSocket()  {

	this.TransportFactory = thrift.NewTBufferedTransportFactory(8192)

	this.ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	var err error
	this.Transport, err = thrift.NewTSocket(addr)
	if err != nil {
		log.Println(err, "Error Opening socket: ")
	}
	if this.Transport == nil{
		log.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	this.Transport, err = this.TransportFactory.GetTransport(this.Transport)
	if err != nil {
		log.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	err = this.Transport.Open()
	if err !=nil {
		log.Println(err, "error opening transport")
	}
}

//GetItemsAll Person
func (this *PersonClient) GetItemsAll() (*myGeneric.TPeronSetResult_, error) {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)

	result, err := client.GetItemsPerson(ctx, BS_PERSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	}
	fmt.Println("result is: ",result)
	return result, nil
}

func (this *PersonClient) GetItemById(id string) (*myGeneric.TPersonResult_, error)  {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)
	result, err := client.GetItemPerson(ctx, BS_PERSON, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	}
	return result, err
}

func (this *PersonClient) Put(item *myGeneric.TPerson) error {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)
	err := client.PutItemPerson(ctx, BS_PERSON, item)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	return nil
}

func (this *PersonClient) Remove(id string) error {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)
	err := client.RemoveItem(ctx, BS_PERSON, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	return nil
}