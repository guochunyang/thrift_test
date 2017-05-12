package main

import (
	"thrift_test/gen-go/tutorial"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)


type CalculatorHandler struct {
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{}
}


func (p *CalculatorHandler) Ping() (err error) {
	fmt.Println("ping()")
	return nil
}

func (p *CalculatorHandler) Add(a int32, b int32) (ret int32, err error) {
	return a + b, nil
}

func main() {
	handler := NewCalculatorHandler()
	processor := tutorial.NewCalculatorServiceProcessor(handler)

	network := "127.0.0.1:8080"

	serverTransport, err := thrift.NewTServerSocket(network)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()


	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in ", network)
	server.Serve()
}
