package entry

import "flag"

const (
	HostPort string = "192.168.165.134:6677"
)

var (
	ServerName = flag.String("service", "/grpc/RtbServer/providers", "rtb service name")
	Port       = flag.String("port", "50051", "listening port")
	Target     = flag.String("target", "10.69.56.55:2181", "register zookeeper address")
)
