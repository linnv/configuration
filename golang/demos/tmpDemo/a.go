package demo

import (
	"fmt"
	"strconv"
	"strings"
)

type EtcdConfig struct {
	PortPeer     string `json:"PortPeer"`
	PortClient   string `json:"PortClient"`
	AddrIP       string `json:"AddrIP"`
	AddrHostName string `json:"AddrHostName"`
}

const count = 3

func genereateEtcdConfig(subnet string) []EtcdConfig {
	ls := make([]EtcdConfig, count)
	const portBasic = 2480
	for i := 0; i < count; i++ {
		p := i*2 + portBasic
		ls[i].PortClient = strconv.Itoa(p)
		ls[i].PortPeer = strconv.Itoa(p + 1)
		ls[i].AddrIP = subnet + strconv.Itoa(i+2)
		ls[i].AddrHostName = "e" + strconv.Itoa(i+2)
	}
	return ls
}

func generateETCDCluster() {
	const (
		templateEtcdStartupConfig = `etcd --name {nodename} -advertise-client-urls http://{hostIP}:{clientPort} -listen-client-urls http://{hostIP}:{clientPort} -initial-advertise-peer-urls http://{hostIP}:{peerPort} -listen-peer-urls http://{hostIP}:{peerPort} -initial-cluster-token {clusterToken} -initial-cluster {nodelist} -initial-cluster-state new`
		// templateEtcdStartupConfig = `etcd --name {nodename} -advertise-client-urls http://0.0.0.0:{clientPort} -listen-client-urls http://0.0.0.0:{clientPort} -initial-advertise-peer-urls http://{hostIP}:{peerPort} -listen-peer-urls http://{hostIP}:{peerPort} -initial-cluster-token {clusterToken} -initial-cluster {nodelist} -initial-cluster-state new`
		templateDockerNetwork = `docker network create --subnet={subnet}/24 {subnetWorkName}`
		// templateDockerRunning     = `docker run -d --net {subnetWorkName} --ip {containerIP} -p {hostPublicPort}:{guestPort} -it -h {hostName} --name {hostName} {imageName}`
		templateDockerRunning = `docker run --net {subnetWorkName} --ip {containerIP} -p {hostPublicPort}:{guestPort} -it -h {hostName} --name {hostName} {imageName}`
	)
	const (
		peerPort     = "2480"
		clientPort   = "2479"
		clusterToken = "justOneEetcdCluster"
	)
	const (
		subnet         = "10.12.0."
		subnetWorkName = "netEtcd"
		imageName      = "xetcd"
	)

	list := genereateEtcdConfig(subnet)
	var nodelist, hostList string
	for i := 0; i < count-1; i++ {
		v := list[i]
		nodelist += v.AddrHostName + "=http://" + v.AddrIP + ":" + v.PortPeer + ","
		hostList += "http://" + v.AddrIP + ":" + v.PortPeer + ","

	}
	nodelist += list[count-1].AddrHostName + "=http://" + list[count-1].AddrIP + ":" + list[count-1].PortPeer
	hostList += "http://" + list[count-1].AddrIP + ":" + list[count-1].PortPeer
	cmdNetwork := strings.NewReplacer([]string{
		"{subnet}", subnet + "0",
		"{subnetWorkName}", subnetWorkName,
	}...).Replace(templateDockerNetwork)
	fmt.Printf("%s\n", cmdNetwork)

	for _, v := range list {
		ret := strings.NewReplacer([]string{
			"{hostIP}", v.AddrIP,
			"{nodelist}", nodelist,
			"{clientPort}", v.PortClient,
			"{peerPort}", v.PortPeer,
			"{nodename}", v.AddrHostName,
			"{clusterToken}", clusterToken,
		}...).Replace(templateEtcdStartupConfig)

		cmdContainer := strings.NewReplacer([]string{
			"{hostName}", v.AddrHostName,
			"{hostPublicPort}", v.PortClient,
			"{guestPort}", v.PortClient,
			"{imageName}", imageName,
			"{subnetWorkName}", subnetWorkName,
			"{containerIP}", v.AddrIP,
		}...).Replace(templateDockerRunning)
		fmt.Printf("%s %s\n", cmdContainer, ret)
	}

	fmt.Printf("etcdctl --endpoint= %+v member list\n", hostList)
}

func JustDemo() {
	generateETCDCluster()
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return
}
