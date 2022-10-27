package nacos_registry

import (
	"base-framework/inf/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

/**
 * @description:
 * @author:xy
 * @date:2022/9/1 11:19
 * @Version: 1.0
 */

var ProviderSet = wire.NewSet(NewNacosConf, NewDiscovery, NewRegistrar)

func NewNacosConf(c *conf.Nacos) vo.NacosClientParam {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(c.Service.Ip, c.Service.Port),
	}

	cc := &constant.ClientConfig{
		NamespaceId: c.Discovery.Namespace,
	}

	return vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
}

// NewDiscovery nacos服务发现注入
func NewDiscovery(param vo.NacosClientParam) registry.Discovery {
	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}

// NewRegistrar 服务注册业务注入
func NewRegistrar(param vo.NacosClientParam) registry.Registrar {

	client, err := clients.NewNamingClient(param)
	if err != nil {
		panic(err)
	}
	return nacos.New(client)
}
