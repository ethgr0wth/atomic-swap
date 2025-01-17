package mcrypto

import (
	"github.com/noot/atomic-swap/common"
	"github.com/noot/atomic-swap/crypto"
)

const (
	addressPrefixMainnet  byte = 18
	addressPrefixStagenet byte = 24
)

// Address represents a base58-encoded string
type Address string

func getChecksum(data ...[]byte) (result [4]byte) {
	keccak256 := crypto.Keccak256(data...)
	copy(result[:], keccak256[:4])
	return
}

// AddressBytes returns the address as bytes for a PrivateKeyPair with the given environment (ie. mainnet or stagenet)
func (kp *PrivateKeyPair) AddressBytes(env common.Environment) []byte {
	return kp.PublicKeyPair().AddressBytes(env)
}

// Address returns the base58-encoded address for a PrivateKeyPair with the given environment
// (ie. mainnet or stagenet)
func (kp *PrivateKeyPair) Address(env common.Environment) Address {
	return Address(EncodeMoneroBase58(kp.AddressBytes(env)))
}

// AddressBytes returns the address as bytes for a PublicKeyPair with the given environment (ie. mainnet or stagenet)
func (kp *PublicKeyPair) AddressBytes(env common.Environment) []byte {
	psk := kp.sk.key.Bytes()
	pvk := kp.vk.key.Bytes()
	c := append(psk, pvk...)

	var prefix byte
	switch env {
	case common.Mainnet, common.Development:
		prefix = addressPrefixMainnet
	case common.Stagenet:
		prefix = addressPrefixStagenet
	}

	// address encoding is:
	// (network_prefix) + (32-byte public spend key) + (32-byte-byte public view key)
	// + first_4_Bytes(Hash(network_prefix + (32-byte public spend key) + (32-byte public view key)))
	checksum := getChecksum(append([]byte{prefix}, c...))
	addr := append(append([]byte{prefix}, c...), checksum[:4]...)
	return addr
}

// Address returns the base58-encoded address for a PublicKeyPair with the given environment
// (ie. mainnet or stagenet)
func (kp *PublicKeyPair) Address(env common.Environment) Address {
	return Address(EncodeMoneroBase58(kp.AddressBytes(env)))
}
