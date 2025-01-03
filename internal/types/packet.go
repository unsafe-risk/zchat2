package types

//go:generate go tool gobe --allow-unsafe --allow-zero-copy-string .

type ProtocolVersion uint32

//go:generate go tool stringer -linecomment -type=ProtocolVersion
const (
	ProtocolVersionUnknown    ProtocolVersion = 0          // PROTOCOL_UNKNOWN
	ProtocolVersion1          ProtocolVersion = 0x5065dd0a // PROTOCOL_V1
	ProtocolVersion1Encrypted ProtocolVersion = 0x7d5fd22a // PROTOCOL_V1_ENCRYPTED
)

type Packet struct {
	Version            ProtocolVersion
	MessageV1          MessageV1          `gobe_enum:"Version=ProtocolVersion1"`
	MessageV1Encrypted MessageV1Encrypted `gobe_enum:"Version=ProtocolVersion1Encrypted"`
}

//go:generate go tool stringer -linecomment -type=MessageType
type MessageType uint16

const (
	MessageTypeUnknown MessageType = 0x0000 // MSG_UNKNOWN

	MessageTypeClientHello MessageType = 0x0101 // MSG_CLIENT_HELLO
	MessageTypeServerHello MessageType = 0x0102 // MSG_SERVER_HELLO

	MessageTypeSyncRequest  MessageType = 0x0201 // MSG_SYNC_REQUEST
	MessageTypeSyncResponse MessageType = 0x0202 // MSG_SYNC_RESPONSE
)

type MessageV1 struct {
	SessionID uint64
	Type      MessageType

	MessageV1SyncRequest  MessageV1SyncRequest  `gobe_enum:"Type=MessageTypeSyncRequest"`
	MessageV1SyncResponse MessageV1SyncResponse `gobe_enum:"Type=MessageTypeSyncResponse"`
}

type MessageV1SyncRequest struct {
	LastOffset int64 // last committed offset
	MaxEntries int64 // max entries to sync
}

type MessageV1SyncResponse struct {
	CommittedOffset int64    // committed offset
	Entries         [][]byte // entries
}

type MessageV1Encrypted struct {
	SessionID uint64
	KeyID     [16]byte
	Nonce     [24]byte
	AuthData  [16]byte
	Payload   []byte
}
