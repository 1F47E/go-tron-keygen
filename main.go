//
// – Generate a random private key and return it as a 32 byte array
// – Convert a private key to a public key using Secp256k1 curve
// - Take the Keccak-256 hash of the binary public key. Keccak-256 is a hashing algorithm used by Tron to derive addresses.
// - Take the first byte of the Keccak-256 hash, which should be 41 in hexadecimal.
// - Append the remaining 20 bytes of the Keccak-256 hash to the first byte 41.
// - Convert the resulting 21-byte binary string to a Base58Check encoded string. Base58Check is a variation of Base58 encoding that adds a checksum to ensure that the address is valid.
// - The resulting Base58Check encoded string is the Tron address corresponding to the public key.
//

package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	// "github.com/btcsuite/btcd/btcutil/base58"
	"go-tron-keygen/base58"

	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/sha3"
)

// ===== Private key

// generate a random private key and return it as a 32 byte array
func privateKey() (*ecdsa.PrivateKey, error) {
	// S256 returns a Curve which implements secp256k1.
	curve := secp.S256()
	return ecdsa.GenerateKey(curve, rand.Reader)
}

func privateKeyToHex(key *ecdsa.PrivateKey) string {
	return fmt.Sprintf("%x", key.D.Bytes())
}

// ===== Public key

// convert a private key to a public key using Secp256k1 curve
func publicKey(key *ecdsa.PrivateKey) ecdsa.PublicKey {

	curve := secp.S256()
	pub := ecdsa.PublicKey{Curve: curve, X: key.X, Y: key.Y}
	return pub

	// return fromPublicKey(&pub)
}

func publicKeyToHex(pubKey *ecdsa.PublicKey) string {
	// convert to bytes
	pubKeyBytes := elliptic.Marshal(pubKey.Curve, pubKey.X, pubKey.Y)
	// convert to hex
	return hex.EncodeToString(pubKeyBytes)
}

// ===== Address

func addressFromPublicKey(pub *ecdsa.PublicKey) [21]byte {
	// shop off the first byte
	// xy := crypto.FromECDSAPub(pub)[1:]
	var xy []byte
	if pub == nil || pub.X == nil || pub.Y == nil {
		xy = nil
	} else {
		xy = elliptic.Marshal(secp.S256(), pub.X, pub.Y)
	}

	h := sha3.NewLegacyKeccak256()
	if _, err := h.Write(xy); err != nil {
		panic("address: unexpected error encountered while writing key")
	}

	// create a byte array of size 21
	var addr [21]byte

	// prefix the address with 0x41, TRON's version byte
	addr[0] = 0x41

	// copy the last 20 bytes of the hash into addr
	// h.Sum(nil) gives you the hash of the data
	// [12:] gives you the last 20 bytes of the 32-byte hash, commonly used in cryptocurrency for shortening addresses
	copy(addr[1:], h.Sum(nil)[12:])

	return addr
}

func addressToHex(address [21]byte) string {
	return fmt.Sprintf("%x", address)
}

// Encode the address as base58 with checksum
// Checksum is first 4 bytes of a sha256 double hash of the address.
func addressToBase58(address [21]byte) string {
	return base58.CheckEncode(address[1:], 0x41)
}

func main() {
	privateKey, err := privateKey()
	if err != nil {
		log.Fatalf("Error generating private key: %v", err)
	}
	privHex := privateKeyToHex(privateKey)
	fmt.Printf("priv key: %s\n", privHex)

	publicKey := publicKey(privateKey)
	pubHex := publicKeyToHex(&publicKey)
	fmt.Printf("pub key: %s\n", pubHex)

	addressBytes := addressFromPublicKey(&publicKey)
	addressHex := addressToHex(addressBytes)
	fmt.Printf("address hex: %x\n", addressHex)

	addressB58 := addressToBase58(addressBytes)
	fmt.Printf("address base58: %s\n", addressB58)
}
