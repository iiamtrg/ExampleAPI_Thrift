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


const addr = "127.0.0.1:9090"

const BS_TEAM = "Team"

type TeamClient struct {
	TransportFactory thrift.TTransportFactory
	ProtocolFactory  thrift.TProtocolFactory
	Transport        thrift.TTransport
}

func (this *TeamClient) InitSocket()  {

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

//GetItemsAll Team
func (this *TeamClient) GetItemsAll() (*myGeneric.TTeamSetResult_, error) {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)

	result, err := client.GetItemsTeam(ctx,BS_TEAM)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	}
	fmt.Println("result is: ",result)
	return result, nil
}

func (this *TeamClient) GetItemById(id string) (*myGeneric.TTeamResult_, error)  {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)
	result, err := client.GetItemTeam(ctx, BS_TEAM, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	}
	return result, err
}

func (this *TeamClient) Put(item *myGeneric.TTeam) error {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)
	err := client.PutItemTeam(ctx, BS_TEAM, item)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	return nil
}

func (this *TeamClient) Remove(id string) error {

	this.InitSocket()
	defer this.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(this.Transport, this.ProtocolFactory)
	err := client.RemoveItem(ctx, BS_TEAM, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	return nil
}