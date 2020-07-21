package models

import (
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

const addr = "127.0.0.1:9090"

const BS_TEAM = "Team"

type TeamClient struct {
	TransportFactory thrift.TTransportFactory
	ProtocolFactory  thrift.TProtocolFactory
	Transport        thrift.TTransport
}

func (t *TeamClient) InitSocket() {

	t.TransportFactory = thrift.NewTBufferedTransportFactory(8192)

	t.ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	var err error
	t.Transport, err = thrift.NewTSocket(addr)
	if err != nil {
		log.Println(err, "Error Opening socket: ")
	}
	if t.Transport == nil {
		log.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	t.Transport, err = t.TransportFactory.GetTransport(t.Transport)
	if err != nil {
		log.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	err = t.Transport.Open()
	if err != nil {
		log.Println(err, "error opening transport")
	}
}

//GetItemsAll Team
func (t *TeamClient) GetItemsAll() (*myGeneric.TTeamSetResult_, error) {

	t.InitSocket()
	defer t.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(t.Transport, t.ProtocolFactory)

	result, err := client.GetItemsTeam(ctx, BS_TEAM)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return nil, err
	}
	return result, nil
}

func (t *TeamClient) GetItemsPagination(offset int32, limit int32) (*myGeneric.TTeamSetResult_, error) {

	t.InitSocket()
	defer t.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(t.Transport, t.ProtocolFactory)

	result, err := client.GetTeamsPagination(ctx, BS_PERSON, offset, limit)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return nil, err
	}
	return result, nil
}

func (t *TeamClient) GetItemById(id string) (*myGeneric.TTeamResult_, error) {

	t.InitSocket()
	defer t.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(t.Transport, t.ProtocolFactory)
	result, err := client.GetItemTeam(ctx, BS_TEAM, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return nil, err
	}
	return result, nil
}

func (t *TeamClient) PutItem(item *myGeneric.TTeam) error {

	t.InitSocket()
	defer t.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(t.Transport, t.ProtocolFactory)
	err := client.PutItemTeam(ctx, BS_TEAM, item)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	return nil
}

func (t *TeamClient) RemoveItem(teamId string) error {

	t.InitSocket()
	defer t.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(t.Transport, t.ProtocolFactory)
	err := client.RemoveItem(ctx, BS_TEAM, teamId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}

	//xóa tất cả các node ràng buộc giữa team và person
	personsOfTeam, err := client.GetPersonsOfTeam(ctx, teamId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	} else {
		// xóa các node bigset name : personID, teamID
		for _, v := range personsOfTeam.Items {
			_ = client.RemoveItem(ctx, v.GetPersonId(), teamId)
			_ = client.RemoveItem(ctx, teamId, v.GetPersonId())
		}
	}

	return nil
}
