package protocol

import (
    "errors"
    "strings"
)

const (
    MsgTypeSize = 10
    ClientIDSize = 10
    HeaderSize = MsgTypeSize + ClientIDSize
    FieldSize  = 10
)

type SimpleMessage struct {
    MessageType string
    ClientID    string
    Message     string
}

// padOrTrim ensures a field is exactly <FieldSize> bytes
func padOrTrim(s string) string {
    if len(s) < FieldSize {
        return s + strings.Repeat(" ", FieldSize-len(s))
    }
    return s[:FieldSize]
}

// Serialize converts a SimpleMessage into bytes + length
func Serialize(msg SimpleMessage) ([]byte, int) {
    mt := padOrTrim(msg.MessageType)
    cid := padOrTrim(msg.ClientID)
    raw := mt + cid + msg.Message

    b := []byte(raw)
    return b, len(b)
}

// Deserialize converts raw bytes into a SimpleMessage
func Deserialize(data []byte) (SimpleMessage, error) {
    if len(data) < HeaderSize {
        return SimpleMessage{}, errors.New("message too short to contain header")
    }

    mt := strings.TrimSpace(string(data[0:MsgTypeSize]))
    cid := strings.TrimSpace(string(data[MsgTypeSize:HeaderSize]))
    body := string(data[HeaderSize:])

    return SimpleMessage{
        MessageType: mt,
        ClientID:    cid,
        Message:     body,
    }, nil
}