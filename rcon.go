package rcon

import (
	"net"
	"strings"
	"time"

	"github.com/Sch8ill/rcon/config"
)

type ConnectionState int

const (
	Connected ConnectionState = iota
	Disconnected
	Authenticated
)

type RconClient struct {
	Address   string
	Password  string
	Timeout   time.Duration
	Conn      net.Conn
	ConnState ConnectionState
}

func Dial(addr string, password string, timeout time.Duration) (*RconClient, error) {
	rconClient := NewClient(addr, password, timeout)

	if err := rconClient.Connect(); err != nil {
		return nil, err
	}
	if err := rconClient.Authenticate(); err != nil {
		return nil, err
	}
	return rconClient, nil
}

func NewClient(addr string, password string, timeout time.Duration) *RconClient {
	return &RconClient{
		Address:   addr,
		Password:  password,
		Timeout:   timeout,
		ConnState: Disconnected,
	}
}

func (rc *RconClient) Connect() error {
	if !strings.Contains(rc.Address, ":") {
		rc.Address = rc.Address + ":" + config.DefaultPort
	}

	if rc.ConnState != Disconnected {
		return errAlreadyConnected
	}

	conn, err := net.DialTimeout("tcp", rc.Address, rc.Timeout)
	if err != nil {
		return err
	}

	rc.ConnState = Connected
	rc.Conn = conn
	return nil
}

func (rc *RconClient) Authenticate() error {
	authPacket := newServerBoundPacket(SERVERDATA_AUTH, rc.Password)
	rc.Conn.Write(authPacket.Bytes())

	resPacket, err := newClientBoundPacket(rc.Conn)

	if err != nil {
		return err
	}

	if resPacket.ID == -1 {
		// packet id -1 means authentication failed
		return errAuthenticationFailed
	}
	rc.ConnState = Authenticated
	return nil
}

func (rc *RconClient) ExecuteCmd(cmd string) (string, error) {
	cmdPacket := newServerBoundPacket(SERVERDATA_EXECCOMMAND, cmd)
	rc.Conn.Write(cmdPacket.Bytes())

	resPacket, err := newClientBoundPacket(rc.Conn)

	if err != nil {
		return "", err
	}

	// the packet send by the the server has to have the same packet id as the request packet
	if cmdPacket.ID != resPacket.ID {
		return "", errWrongPacketID
	}

	return resPacket.getBody(), nil
}

func (rc *RconClient) Close() {
	rc.Conn.Close()
	rc.ConnState = Disconnected
}
