package models

import (
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"context"
	"log"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

const BS_PERSON = "Person"

type PersonClient struct {
	TransportFactory thrift.TTransportFactory
	ProtocolFactory  thrift.TProtocolFactory
	Transport        thrift.TTransport
}

func (p *PersonClient) InitSocket() {

	p.TransportFactory = thrift.NewTBufferedTransportFactory(8192)

	p.ProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	var err error
	p.Transport, err = thrift.NewTSocket(addr)
	if err != nil {
		log.Println(err, "Error Opening socket: ")
	}
	if p.Transport == nil {
		log.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	p.Transport, err = p.TransportFactory.GetTransport(p.Transport)
	if err != nil {
		log.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	err = p.Transport.Open()
	if err != nil {
		log.Println(err, "error opening transport")
	}
}

//GetItemsAll Person
func (p *PersonClient) GetItemsAll() (*myGeneric.TPeronSetResult_, error) {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	result, err := client.GetItemsPerson(ctx, BS_PERSON)
	if err != nil {

		return nil, err
	}
	return result, nil
}

func (p *PersonClient) GetItemsPagination(offset int32, limit int32) (*myGeneric.TPeronSetResult_, error) {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	result, err := client.GetPersonsPagination(ctx, BS_PERSON, offset, limit)
	if err != nil {

		return nil, err
	}
	return result, nil
}

func (p *PersonClient) GetPersonTeam(personID string) (*myGeneric.TTeamResult_, error) {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	result, err := client.GetPersonIsTeam(ctx, personID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonClient) GetPersonsOfTeam(teamId string) (*myGeneric.TPeronSetResult_, error) {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	result, err := client.GetPersonsOfTeam(ctx, teamId)
	if err != nil {

		return nil, err
	}
	return result, nil
}

func (p *PersonClient) GetPersonOfTeamPagination(teamId string, offset int32, limit int32) (*myGeneric.TPeronSetResult_, error) {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	result, err := client.GetPersonsOfTeamPagination(ctx, teamId, offset, limit)
	if err != nil {

		return nil, err
	}
	return result, nil
}

func (p *PersonClient) GetItemById(id string) (*myGeneric.TPersonResult_, error) {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)
	result, err := client.GetItemPerson(ctx, BS_PERSON, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonClient) PutItem(item *myGeneric.TPerson) error {

	p.InitSocket()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer p.Transport.Close()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	if item.GetTeamId() != "" {
		_, err := client.GetItemTeam(ctx, BS_TEAM, item.GetTeamId())
		if err != nil {
			return err
		} else {
			ok, err := client.ItemIsExist(ctx, BS_TEAM, item.GetTeamId())
			if err != nil || !ok {
				return err
			}
			err = client.PutItemPerson(ctx, BS_PERSON, item)
			if err != nil {
				return err
			}
			err = client.PutPersonToTeam(ctx, item.GetTeamId(), item.GetPersonId())
			if err != nil {
				return err
			}
			return nil
		}
	} else {
		err := client.PutItemPerson(ctx, BS_PERSON, item)
		if err != nil {
			return err
		}
		return nil
	}
}

func (p *PersonClient) RemoveItem(personId string) error {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)
	//get person
	person, err := client.GetItemPerson(ctx, BS_PERSON, personId)
	if err != nil {
		log.Println(err)
	} else {
		err := client.RemoveItem(ctx, BS_PERSON, personId)
		if err != nil {
			return err
		}
		err = client.RemoveItem(ctx, person.GetItem().GetTeamId(), personId)
		if err != nil {
			log.Printf("can not remove person(%s) in team(%s)", personId, person.GetItem().GetTeamId())
		}
	}
	return nil
}
