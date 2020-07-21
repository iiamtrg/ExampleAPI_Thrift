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
		fmt.Fprintln(os.Stderr, "bigset error", err)
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
		fmt.Fprintln(os.Stderr, "bigset error", err)
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
		fmt.Fprintln(os.Stderr, "bigset error", err)
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
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return nil, err
	}
	return result, nil
}

func (p *PersonClient) PutItem(item *myGeneric.TPerson) error {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)
	err := client.PutItemPerson(ctx, BS_PERSON, item)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	return nil
}

func (p *PersonClient) PutPersonIsTeam(personId string, teamId string) error {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	//check team exist in bigset
	ok, err := client.ItemIsExist(ctx, BS_TEAM, teamId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	if ok {
		// check person exist
		ok2, err := client.ItemIsExist(ctx, BS_PERSON, personId)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bigset error", err)
			return err
		}
		if ok2 {
			// thêm / cập nhật team cho person
			err := client.PutPersonIsTeam(ctx, personId, teamId)
			if err != nil {
				fmt.Fprintln(os.Stderr, "bigset error", err)
				return err
			}
			// thêm / cập nhật person vào team
			err = client.PutPersonToTeam(ctx, teamId, personId)
			if err != nil {
				fmt.Fprintln(os.Stderr, "bigset error", err)
				return err
			}
			return nil
		}
		return fmt.Errorf("person is not exist")
	}
	return fmt.Errorf("team is not exist")
}

func (p *PersonClient) PutPersonToTeam(personId string, teamId string) error {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	//check team exist in bigset
	ok, err := client.ItemIsExist(ctx, BS_TEAM, teamId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	if ok {
		ok2, err := client.ItemIsExist(ctx, BS_PERSON, personId)
		if err != nil {
			fmt.Fprintln(os.Stderr, "bigset error", err)
			return err
		}
		if ok2 {

			// thêm person vào team
			err := client.PutPersonToTeam(ctx, teamId, personId)
			if err != nil {
				fmt.Fprintln(os.Stderr, "bigset error", err)
				return err
			}

			// thêm team cho person
			err = client.PutPersonIsTeam(ctx, personId, teamId)
			if err != nil {
				fmt.Fprintln(os.Stderr, "bigset error", err)
				return err
			}
			return nil
		}
		return fmt.Errorf("person is not exist")
	}
	return fmt.Errorf("team is not exist")
}

func (p *PersonClient) PutMultiPersonsToTeam(personIds []string, teamId string) error {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)

	//check team exist in bigset
	ok, err := client.ItemIsExist(ctx, BS_TEAM, teamId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	if ok {
		for _, v := range personIds {
			ok2, err := client.ItemIsExist(ctx, BS_PERSON, v)
			if err != nil || !ok2 {
				fmt.Fprintln(os.Stderr, "bigset error", err)
				return err
			}
			if ok2 {
				err := client.PutPersonToTeam(ctx, teamId, v)
				if err != nil {
					fmt.Fprintln(os.Stderr, "bigset error", err)
					return err
				}
				err = client.PutPersonIsTeam(ctx, v, teamId)
				if err != nil {
					fmt.Fprintln(os.Stderr, "bigset error", err)
					return err
				}
			}
		}
		return nil
	}
	return fmt.Errorf("%s", "team is not exist")
}

func (p *PersonClient) RemoveItem(personId string) error {

	p.InitSocket()
	defer p.Transport.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client := myGeneric.NewTGenericServiceClientFactory(p.Transport, p.ProtocolFactory)
	err := client.RemoveItem(ctx, BS_PERSON, personId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
		return err
	}
	//get person's team
	team, err := client.GetPersonIsTeam(ctx, personId)
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	} else {
		err := client.RemoveItem(ctx, personId, team.GetItem().GetTeamId())
		if err != nil {
			fmt.Fprintln(os.Stderr, "bigset error", err)
		}
	}

	return nil
}
