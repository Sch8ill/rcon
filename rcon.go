package rcon

import (
	"net"
	"strings"
	"time"

	"github.com/Sch8ill/rcon/config"
)

type ConnectionState int

// an "enum" that represents the current connection state
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

// creates a RCON client and connects and authenticates it
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

// creates a new RCON client
func NewClient(addr string, password string, timeout time.Duration) *RconClient {
	return &RconClient{
		Address:   addr,
		Password:  password,
		Timeout:   timeout,
		ConnState: Disconnected,
	}
}

// establishes the underlying TCP connenction of the RCON client
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

// authenticates the client using the clients password
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

// executes a command on the remote server
func (rc *RconClient) ExecuteCmd(cmd string) (string, error) {
	// check if the client is connected and authenticated
	if rc.ConnState != Authenticated {
		return "", errWrongConenctionState
	}

	// construct the packet and write it to the socket
	cmdPacket := newServerBoundPacket(SERVERDATA_EXECCOMMAND, cmd)
	rc.Conn.Write(cmdPacket.Bytes())

	// parse the received packet
	resPacket, err := newClientBoundPacket(rc.Conn)

	if err != nil {
		return "", err
	}

	// the response packet has to have the same packet id as the request packet
	if cmdPacket.ID != resPacket.ID {
		return "", errWrongPacketID
	}

	return resPacket.getBody(), nil
}

// "closes" the RCON client by terminating the underlying TCP connection and setting the clients connection state to disconnected
func (rc *RconClient) Close() {
	rc.Conn.Close()
	rc.ConnState = Disconnected
}
