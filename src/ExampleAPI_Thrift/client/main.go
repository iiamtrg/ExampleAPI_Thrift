package main

import (
	"ExampleAPI_Bigset_Thrift/thrift/gen-go/myGeneric"
	"context"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"os"
)


func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

var ctx = context.Background()
const addr = "127.0.0.1:9090"

func main()  {


	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	fmt.Printf("addr   = '%v'\n", addr)
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		log.Println(err, "Error Openning socket: ")
	}
	if transport == nil{
		fmt.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		 fmt.Println("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}
	err = transport.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer transport.Close()
	client := myGeneric.NewTGenericServiceClientFactory(transport, protocolFactory)

	result, err := client.GetItemsTeam(ctx,"Team")
	if err != nil {
		fmt.Fprintln(os.Stderr, "bigset error", err)
	}
	fmt.Println("result is: ",result)

}