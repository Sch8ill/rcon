package rcon

import (
	"errors"
)

var (
	errAlreadyConnected     error = errors.New("the client is already connected")
	errAuthenticationFailed error = errors.New("authentication failed (wrong password?)")
	errWrongPacketID        error = errors.New("packet send by server has the wrong packet id")
	errPacketBodyToShort    error = errors.New("body of packet send by server is to short")
	errWrongConenctionState error = errors.New("operation canceld because of wrong connection state")
)
