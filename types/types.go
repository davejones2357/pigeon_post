package types

import (
    "fmt"
    "strings"

    "pigeon_post/protocol"
)

const (
    CmdConnect = "CONNECT"
    CmdSend    = "SEND"
    CmdPing    = "PING"
    CmdPong    = "PONG"
    CmdQuit    = "QUIT"
)

// Command is the generic interface for network commands.
type Command interface {
    Type() string
    ClientID() string
    ToMessage() protocol.SimpleMessage
}

// Connect represents: CONNECT <client_id>
type Connect struct {
    ID string
}

func NewConnect(id string) Connect { return Connect{ID: strings.TrimSpace(id)} }
func (c Connect) Type() string     { return CmdConnect }
func (c Connect) ClientID() string { return c.ID }
func (c Connect) ToMessage() protocol.SimpleMessage {
    return protocol.SimpleMessage{
        MessageType: CmdConnect,
        ClientID:    c.ID,
        Message:     "",
    }
}

// Send represents: SEND <payload>
type Send struct {
    ID      string
    Payload string
}

func NewSend(id, payload string) Send {
    return Send{ID: strings.TrimSpace(id), Payload: payload}
}
func (s Send) Type() string     { return CmdSend }
func (s Send) ClientID() string { return s.ID }
func (s Send) ToMessage() protocol.SimpleMessage {
    return protocol.SimpleMessage{
        MessageType: CmdSend,
        ClientID:    s.ID,
        Message:     s.Payload,
    }
}

// Ping represents: PING (server should reply with PONG)
type Ping struct {
    ID string
}

func NewPing(id string) Ping { return Ping{ID: strings.TrimSpace(id)} }
func (p Ping) Type() string  { return CmdPing }
func (p Ping) ClientID() string {
    return p.ID
}
func (p Ping) ToMessage() protocol.SimpleMessage {
    return protocol.SimpleMessage{
        MessageType: CmdPing,
        ClientID:    p.ID,
        Message:     "",
    }
}

// Pong represents server reply to PING: PONG
type Pong struct {
    ID string
}

func NewPong(id string) Pong { return Pong{ID: strings.TrimSpace(id)} }
func (p Pong) Type() string  { return CmdPong }
func (p Pong) ClientID() string {
    return p.ID
}
func (p Pong) ToMessage() protocol.SimpleMessage {
    return protocol.SimpleMessage{
        MessageType: CmdPong,
        ClientID:    p.ID,
        Message:     "",
    }
}

// Quit represents: QUIT (graceful disconnect)
type Quit struct {
    ID string
}

func NewQuit(id string) Quit { return Quit{ID: strings.TrimSpace(id)} }
func (q Quit) Type() string  { return CmdQuit }
func (q Quit) ClientID() string {
    return q.ID
}
func (q Quit) ToMessage() protocol.SimpleMessage {
    return protocol.SimpleMessage{
        MessageType: CmdQuit,
        ClientID:    q.ID,
        Message:     "",
    }
}

// Parse converts a protocol.SimpleMessage into a typed Command.
func Parse(msg protocol.SimpleMessage) (Command, error) {
    mt := strings.ToUpper(strings.TrimSpace(msg.MessageType))
    cid := strings.TrimSpace(msg.ClientID)

    switch mt {
    case CmdConnect:
        return Connect{ID: cid}, nil
    case CmdSend:
        return Send{ID: cid, Payload: msg.Message}, nil
    case CmdPing:
        return Ping{ID: cid}, nil
    case CmdPong:
        return Pong{ID: cid}, nil
    case CmdQuit:
        return Quit{ID: cid}, nil
    default:
        return nil, fmt.Errorf("types: unknown message type %q", mt)
    }
}