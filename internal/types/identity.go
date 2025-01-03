package types

//go:generate go tool stringer -linecomment -type=IdentityVersion
type IdentityVersion uint32

const (
	IdentityVersionUnknown IdentityVersion = 0          // IDENTITY_UNKNOWN
	IdentityVersion1       IdentityVersion = 0xda048c6b // IDENTITY_V1
)

type Identity struct {
	Version IdentityVersion

	IdentityV1 IdentityV1 `gobe_enum:"Version=IdentityVersion1"`

	ValidFrom  int64 // unix timestamp (seconds)
	ValidUntil int64 // unix timestamp (seconds)
}

type IdentityV1 struct {
	Ed25519PublicKey [32]byte   // Ed25519
	X25519PublicKey  [32]byte   // X25519
	MLKEMPublicKey   [1568]byte // ML-KEM-1024
	MLDSAPublicKey   [2592]byte // ML-DSA-87
}
