package discovery

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-service/pkg/log"
	"google.golang.org/grpc"
)

func GetServiceClient(serviceName string, groupName string, namespace string) *model.Instance {
	instance := SelectOneHealthyInstance(*ServerCenter, vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName,
		Clusters:    []string{namespace},
	})
	return instance
}

func GetGrpcServiceDial(serviceName string, groupName string, namespace string) *grpc.ClientConn {
	client := GetServiceClient(serviceName, groupName, namespace)
	address := fmt.Sprintf("%s:%d", client.Ip, client.Port)
	dial, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error("grpc", err.Error())
	}
	return dial
}
