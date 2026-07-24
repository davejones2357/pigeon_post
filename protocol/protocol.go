package protocol

import (
    "errors"
    "strings"
)

const (
    HeaderSize = 20 // 10 bytes type + 10 bytes client ID
    FieldSize  = 10
)

type SimpleMessage struct {
    MessageType string
    ClientID    string
    Message     string
}

// padOrTrim ensures a field is exactly 10 bytes
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

    mt := strings.TrimSpace(string(data[0:10]))
    cid := strings.TrimSpace(string(data[10:20]))
    body := string(data[20:])

    return SimpleMessage{
        MessageType: mt,
        ClientID:    cid,
        Message:     body,
    }, nil
}