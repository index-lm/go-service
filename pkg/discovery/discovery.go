package discovery

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	utils "go-service/internal/pkg/sys"
	"go-service/pkg/log"
)

var ServerCenter *naming_client.INamingClient

type InitConfig struct {
	ipAddr      string
	port        uint64
	namespaceId string
	serverName  string
}

func Initialize(ipAddr string, port uint64, namespaceId string, serverName string,logPath string, cachePath string) {
	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ipAddr, port, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(namespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(logPath),
		constant.WithCacheDir(cachePath),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("error"),
	)

	// create naming client
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		log.Info("sys", err.Error())
	}
	ServerCenter = &client
	//Register with default cluster and group
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          utils.ServerIP,
		Port:        utils.ServerPort,
		ServiceName: serverName,
		//GroupName:   "DEFAULT_GROUP",
		ClusterName: namespaceId,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"id": "si-1"},
	})

	//DeRegister with sys,port,serviceName
	//ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
	//Note:sys=10.0.0.10,port=8848 should belong to the cluster of DEFAULT and the group of DEFAULT_GROUP.
	//ExampleServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
	//	Ip:          "10.0.0.10",
	//	Port:        8848,
	//	ServiceName: "im-gateway",
	//	//GroupName:   "DEFAULT_GROUP",
	//	Cluster:     "lecare-develop",
	//	Ephemeral:   true, //it must be true
	//})

	//Register with default cluster and group
	//ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	//ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
	//	Ip:          "10.0.0.10",
	//	Port:        8848,
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	ClusterName: "lecare-develop",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	Metadata:    map[string]string{"idc": "shanghai"},
	//})

	//time.Sleep(1 * time.Second)

	//Get service with serviceName
	//ClusterName=DEFAULT, GroupName=DEFAULT_GROUP
	//ExampleServiceClient_GetService(client, vo.GetServiceParam{
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	Clusters:    []string{"lecare-develop"},
	//})
	//
	////SelectAllInstance
	////GroupName=DEFAULT_GROUP
	//ExampleServiceClient_SelectAllInstances(client, vo.SelectAllInstancesParam{
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	Clusters:    []string{"lecare-develop"},
	//})
	//
	////SelectInstances only return the instances of healthy=${HealthyOnly},enable=true and weight>0
	////ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	//ExampleServiceClient_SelectInstances(client, vo.SelectInstancesParam{
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	Clusters:    []string{"lecare-develop"},
	//})
	//
	////SelectOneHealthyInstance return one instance by WRR strategy for load balance
	////And the instance should be health=true,enable=true and weight>0
	////ClusterName=DEFAULT,GroupName=DEFAULT_GROUP
	//ExampleServiceClient_SelectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	Clusters:    []string{"lecare-develop"},
	//})
	//
	////Subscribe key=serviceName+groupName+cluster
	////Note:We call add multiple SubscribeCallback with the same key.
	//param := &vo.SubscribeParam{
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	Clusters:    []string{"lecare-develop"},
	//	SubscribeCallback: func(services []model.Instance, err error) {
	//		fmt.Printf("callback111 return services:%s \n\n", util.ToJsonString(services))
	//	},
	//}
	//ExampleServiceClient_Subscribe(client, param)

	//ExampleServiceClient_RegisterServiceInstance(client, vo.RegisterInstanceParam{
	//	Ip:          "10.0.0.10",
	//	Port:        8848,
	//	ServiceName: "im-gateway",
	//	GroupName:   "DEFAULT_GROUP",
	//	ClusterName: "lecare-develop",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	Metadata:    map[string]string{"idc": "beijing"},
	//})
	////wait for client pull change from server
	//time.Sleep(3 * time.Second)

	//Now we just unsubscribe callback1, and callback2 will still receive change event
	//ExampleServiceClient_UnSubscribe(client, param)
	//ExampleServiceClient_DeRegisterServiceInstance(client, vo.DeregisterInstanceParam{
	//	Ip:          "10.0.0.112",
	//	Ephemeral:   true,
	//	Port:        8848,
	//	ServiceName: "im-gateway",
	//	Cluster:     "cluster-b",
	//})
	////wait for client pull change from server
	//time.Sleep(3 * time.Second)

	//GeAllService will get the list of service name
	//NameSpace default value is public.If the client set the namespaceId, NameSpace will use it.
	//GroupName default value is DEFAULT_GROUP
	//ExampleServiceClient_GetAllService(client, vo.GetAllServiceInfoParam{
	//	PageNo:   1,
	//	PageSize: 10,
	//})

	//// 创建clientConfig的另一种方式
	//clientConfig := *constant.NewClientConfig(
	//	constant.WithNamespaceId("lecare-develop"), //当namespace是public时，此处填空字符串。
	//	constant.WithTimeoutMs(5000),
	//	constant.WithNotLoadCacheAtStart(true),
	//	constant.WithLogDir("D:\\opt\\goog"),
	//	constant.WithCacheDir("D:\\opt\\gocache"),
	//	constant.WithRotateTime("1h"),
	//	constant.WithMaxAge(3),
	//	constant.WithLogLevel("debug"),
	//)
	//// 创建serverConfig的另一种方式
	//serverConfigs := []constant.ServerConfig{
	//	*constant.NewServerConfig(
	//		"172.16.20.20",
	//		80,
	//		constant.WithScheme("http"),
	//		constant.WithContextPath("/nacos")),
	//}
	//// 创建服务发现客户端的另一种方式 (推荐)
	//namingClient, err := clients.NewNamingClient(
	//	vo.NacosClientParam{
	//		ClientConfig:  &clientConfig,
	//		ServerConfigs: serverConfigs,
	//	},
	//)
	//if err!=nil  {
	//	log.Info("sys",err.Error())
	//}
	//// 创建动态配置客户端的另一种方式 (推荐)
	////configClient, err := clients.NewConfigClient(
	////	vo.NacosClientParam{
	////		ClientConfig:  &clientConfig,
	////		ServerConfigs: serverConfigs,
	////	},
	////)
	////if err!=nil  {
	////	log.Info("sys","配置中心失败")
	////}
	//_, err = namingClient.RegisterInstance(vo.RegisterInstanceParam{
	//	Ip:          "172.16.20.20",
	//	Port:        8848,
	//	ServiceName: "im-gateway",
	//	//GroupName:   "DEFAULT_GROUP",
	//	ClusterName: "lecare-develop",
	//	Weight:      10,
	//	Enable:      true,
	//	Healthy:     true,
	//	Ephemeral:   true,
	//	//Metadata:    map[string]string{"idc": "shanghai"},
	//})
	//if err!=nil  {
	//	log.Info("sys",err.Error())
	//}

}

func ExampleServiceClient_RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	success, err := client.RegisterInstance(param)
	if err != nil {
		log.Info("sys", err.Error())
	}
	log.Info("nacos", fmt.Sprintf("RegisterServiceInstance,param:%+v,result:%+v", param, success))
}

func ExampleServiceClient_DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	success, _ := client.DeregisterInstance(param)
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func ExampleServiceClient_GetService(client naming_client.INamingClient, param vo.GetServiceParam) {
	service, _ := client.GetService(param)
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
}

func ExampleServiceClient_SelectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam) {
	instances, _ := client.SelectAllInstances(param)
	fmt.Printf("SelectAllInstance,param:%+v, result:%+v \n\n", param, instances)
}

func SelectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam) []model.Instance {
	instances, err := client.SelectInstances(param)
	if err != nil {
		log.Error("grpc", err.Error())
	}
	return instances
}

func SelectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) *model.Instance {
	instances, err := client.SelectOneHealthyInstance(param)
	if err != nil {
		log.Error("grpc", err.Error())
	}
	return instances
}

func ExampleServiceClient_Subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Subscribe(param)
}

func ExampleServiceClient_UnSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	client.Unsubscribe(param)
}

func ExampleServiceClient_GetAllService(client naming_client.INamingClient, param vo.GetAllServiceInfoParam) {
	service, _ := client.GetAllServicesInfo(param)
	fmt.Printf("GetAllService,param:%+v, result:%+v \n\n", param, service)
}
