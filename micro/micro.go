package micro

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"net"
)

var NamingClient *naming_client.INamingClient

type Micro interface {
	// NewServiceRegister 服务注册 /**
	NewServiceRegister(serviceName string, port int)
}
type NacosMicro struct {
}

func TestHello() {
	println("test Hello invoke ................. ")
}
func (f *NacosMicro) NewServiceRegister(serviceName string, port int) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         Cfg.Nacos.NameSpaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "localhost",
			Port:        8848,
			ContextPath: "/nacos",
		},
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		panic("error config .... ")
	}

	namingClient, err = clients.CreateNamingClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	// Register Gin service to Nacos
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          GetLocalIPAddress(),
		Port:        uint64(port),
		ServiceName: serviceName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    map[string]string{"version": "1.0"},
		ClusterName: "DEFAULT",
		GroupName:   Cfg.Nacos.GroupName,
		Ephemeral:   true,
	})

	if err != nil || !success {
		fmt.Println("Failed to register Gin service to Nacos:", err)
		panic("error")
	}
	fmt.Println("service register successful on nacos")
	NamingClient = &namingClient

}
func GetLocalIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.String()
			}
		}
	}
	return ""
}
