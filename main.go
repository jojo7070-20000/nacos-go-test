package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	var rootPath = "/Users/apple/work/golang/src/nacos-go-test"

	// 从控制台命名空间管理的"命名空间详情"中拷贝 End Point、命名空间 ID
	//var endpoint = "192.168.5.14"
	var namespaceId = "e80cb224-cf7d-4f57-a253-b05d382beeca" // dev namespaceId

	// 推荐使用 RAM 用户的 accessKey、secretKey
	//var accessKey = ""
	//var secretKey = ""

	clientConfig := constant.ClientConfig{
		//Endpoint:    endpoint + ":8848",
		NamespaceId: namespaceId,
		//AccessKey:      accessKey,
		//SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
		CacheDir:       rootPath + "/CacheDir",
		LogDir:         rootPath + "/LogDir",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "192.168.5.14",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		fmt.Println("configClient err:", err)
		return
	}
	success1, err1 := configClient.PublishConfig(vo.ConfigParam{
		AppName: "APP NAME",

		DataId:  "xxx.test",
		Group:   "DEFAULT_GROUP",
		Content: "hello world!222222, yes!!! write by Gopher!"})

	fmt.Println(success1, err1)

	fmt.Println("End Write....")

	fmt.Println("Begin Read ....")

	configClientR, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr:      "192.168.5.14",
				ContextPath: "/nacos",
				Port:        8848,
			},
		},
		"clientConfig": constant.ClientConfig{
			//Endpoint:    endpoint + ":8848",
			NamespaceId: "", // 默认 public, 没有NamespaceId
			//AccessKey:      accessKey,
			//SecretKey:      secretKey,
			TimeoutMs:      5 * 1000,
			ListenInterval: 30 * 1000,
			CacheDir:       rootPath + "/CacheDir",
			LogDir:         rootPath + "/LogDir",
		},
	})

	var dataId = "redis.yaml"
	var group = "DEFAULT_GROUP"

	success, err := configClientR.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})
	fmt.Println("config get error:", err)
	fmt.Println(success)

	fmt.Println("begin 获取服务！！")

	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig": constant.ClientConfig{
			//Endpoint:    endpoint + ":8848",
			NamespaceId: "", // public
			//AccessKey:      accessKey,
			//SecretKey:      secretKey,
			TimeoutMs:      5 * 1000,
			ListenInterval: 30 * 1000,
			CacheDir:       rootPath + "/CacheDir",
			LogDir:         rootPath + "/LogDir",
		},
	})

	// 获取服务：GetService
	service, _ := namingClient.GetService(vo.GetServiceParam{
		ServiceName: "xd-user-server",
		Clusters:    []string{"DEFAULT"}, // 集群名字 默认是default
	})

	fmt.Println("获取注册的服务器：", service.Hosts)

	instances, _ := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		ServiceName: "xd-devops-server",
		Clusters:    []string{"DEFAULT"},
	})

	for _, v := range instances {
		fmt.Println("获取所有的实例列表：SelectAllInstances：", v.Ip, v.Port, v.Weight)
	}

	instancesH, _ := namingClient.SelectInstances(vo.SelectInstancesParam{
		ServiceName: "xd-devops-server",
		Clusters:    []string{"DEFAULT"},
		HealthyOnly: true,
	})

	for _, v := range instancesH {
		fmt.Println("获取实例列表：SelectInstances：", v.Ip, v.Port, v.Weight)
	}

	instanceHOne, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "xd-devops-server",
		Clusters:    []string{"DEFAULT"},
	})
	fmt.Println("获取一个健康的实例（加权轮训负载均衡）：SelectOneHealthyInstance", instanceHOne.Ip, instanceHOne.Port, instanceHOne.Weight)

	instanceHOne2, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "xd-devops-server",
		Clusters:    []string{"DEFAULT"},
	})
	fmt.Println("获取一个健康的实例（加权轮训负载均衡）：SelectOneHealthyInstance", instanceHOne2.Ip, instanceHOne2.Port, instanceHOne2.Weight)

	instanceHOne3, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: "xd-devops-server",
		Clusters:    []string{"DEFAULT"},
	})
	fmt.Println("获取一个健康的实例（加权轮训负载均衡）：SelectOneHealthyInstance", instanceHOne3.Ip, instanceHOne3.Port, instanceHOne3.Weight)


}
