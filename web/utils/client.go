/*
 *@time       2021/10/19 22:25
 *@version    1.0.0
 *@author     11726
 */

package utils

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func NewClient() micro.Service {
	newRegistry := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	// New Service
	client := micro.NewService(
		micro.Registry(newRegistry),
	)
	return client
}
