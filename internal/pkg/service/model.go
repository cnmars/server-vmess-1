package service

const (
	protocol = "vmess"
	TLS      = "tls"
	TCP      = "tcp"
	WS       = "ws"
	GRPC     = "grpc"
	H2       = "h2"
)

// Service is the interface of all the services running in the panel
type Service interface {
	Start() error
	Close() error
}

type CertConfig struct {
	CertFile string
	KeyFile  string
}
