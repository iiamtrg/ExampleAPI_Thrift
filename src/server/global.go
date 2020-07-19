package main

import (
	"ExampleAPI_Bigset_Thrift/src/appconfig"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
	"log"
	"sync"
)

var (
	bigsetIf StringBigsetService.StringBigsetServiceIf
	Once     sync.Once
)

const ETCDENDPOINT = "127.0.0.1:2379"

const SID = "test/test"
const HOST = "127.0.0.1"
const PORT = "18407"

func init()  {
	// bigset config
	config := &appconfig.AppConfig{}
	etcdEndpoints := []string{ETCDENDPOINT}
	config.EtcdServerEndpoints = etcdEndpoints
	config.SourceMappingNewSID = SID
	config.SourceMappingNewHost = HOST
	config.SourceMappingNewPort = PORT
	config.SourceMappingNewProtocol = "binary"

	appconfig.Config = config
}

func InitBigSetIf() {

	log.Println(appconfig.Config, "appconfig.Config")
	Once.Do(func() {
		log.Println(appconfig.Config, "appconfig.Config")
		bigsetIf = StringBigsetService.NewStringBigsetServiceModel(appconfig.Config.SourceMappingNewSID,
			appconfig.Config.EtcdServerEndpoints,
			GoEndpointBackendManager.EndPoint{
				Host:      appconfig.Config.SourceMappingNewHost,
				Port:      appconfig.Config.SourceMappingNewPort,
				ServiceID: appconfig.Config.SourceMappingNewSID,
			})
		log.Println(bigsetIf, "bigsetIf")
	})
}
