package rcon

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"net"
	"time"
)

const (
	SERVERDATA_AUTH           int32 = 3
	SERVERDATA_AUTH_RESPONSE  int32 = 2
	SERVERDATA_EXECCOMMAND    int32 = 2
	SERVERDATA_RESPONSE_VALUE int32 = 0

	HEADER_SIZE     int = 4
	MIN_PACKET_SIZE int = 10
)

type Packet struct {
	ID   int32
	Type int32
	Body []byte
}

func newServerBoundPacket(packetType int32, body string) *Packet {
	rand.Seed(time.Now().UnixNano())
	return &Packet{
		ID:   rand.Int31(),
		Type: packetType,
		Body: []byte(body),
	}
}

func (p *Packet) Bytes() []byte {
	packetLength := int32(len(p.Body) + 10)
	packetBuffer := bytes.NewBuffer(make([]byte, 0, packetLength+4))

	// write the packet length
	binary.Write(packetBuffer, binary.LittleEndian, packetLength)
	// write the packet id
	binary.Write(packetBuffer, binary.LittleEndian, p.ID)
	// write the packet type
	binary.Write(packetBuffer, binary.LittleEndian, p.Type)
	// write the packet body
	packetBuffer.Write(p.Body)
	// write two zero bytes to null terminate the string and the packet
	packetBuffer.Write([]byte{0x00, 0x00})

	return packetBuffer.Bytes()
}

func newClientBoundPacket(conn net.Conn) (*Packet, error) {
	packet := &Packet{}
	var packetSize int32

	if err := binary.Read(conn, binary.LittleEndian, &packetSize); err != nil {
		return nil, err
	}
	if err := binary.Read(conn, binary.LittleEndian, &packet.ID); err != nil {
		return nil, err
	}
	if err := binary.Read(conn, binary.LittleEndian, &packet.Type); err != nil {
		return nil, err
	}

	if packetSize - int32(HEADER_SIZE) < 2 {
		return nil, errPacketBodyToShort
	}
	body := make([]byte, packetSize - int32(HEADER_SIZE))
	length, err := conn.Read(body)

	if err != nil {
		return nil, err
	}
	packet.Body = body[:length]

	return packet, nil
}

func (p *Packet) getBody() string {
	return string(p.Body)
}
