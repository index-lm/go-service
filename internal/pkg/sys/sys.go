package sys

import (
	"context"
	"net"
)

var (
	ServerIP   string
	ServerPort uint64
	ServerName string
)

func init() {
	//fixme In the configuration file, sys takes precedence, if not, get the valid network card sys of the machine
	//if ServerIP != "" {
	//	ServerIP = config.Config.ServerIP
	//	return
	//}

	// see https://gist.github.com/jniltinho/9787946#gistcomment-3019898
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ServerIP = localAddr.IP.String()
}

func Initialize(serverPort uint64, serverName string) {
	ServerName = serverName
	ServerPort = serverPort
}

func GetCxt() context.Context {
	return context.Background()
}