```
  __________  ____  _   __
 /_  __/ __ \/ __ \/ | / /
  / / / /_/ / / / /  |/ / 
 / / / _, _/ /_/ / /|  /  
/_/ /_/ |_|\____/_/ |_/   
keygen
```

# Go-Tron-Keygen

Go-Tron-Keygen is a simple Go-based project designed to demonstrate the process of generating Tron keys and addresses.

## Overview

1. Generate a random private key (32-byte array).
2. Convert a private key to a public key using the Secp256k1 curve.
3. Calculate the Keccak-256 hash of the binary format of the public key (Keccak-256 is the hashing algorithm used by Tron to derive addresses).
4. Take the first byte of the Keccak-256 hash (should be 41 in hexadecimal).
5. Append the rest 20 bytes of the Keccak-256 hash to the first byte (41).
6. Convert the 21-byte binary string into a Base58Check encoded string. Base58Check is a variant of Base58 encoding that includes a checksum for verifying the validity of the address.
7. The resulting Base58Check encoded string is the Tron address associated with the public key.


## Usage

```
go run main.go


priv key: 1234...
pub key: 5678...
address hex: 9abc...
address base58: TDef...
```


## Dependencies

This project uses several Go packages:

- `crypto/ecdsa`: Implements the elliptic curve Digital Signature Algorithm.
- `crypto/elliptic`: Provides the elliptic curve.
- `crypto/rand`: Offers a cryptographically secure random number generator.
- `encoding/hex`: Functions for encoding and decoding hexadecimal strings.
- `github.com/decred/dcrd/dcrec/secp256k1/v4`: Supplies the Secp256k1 elliptic curve.
- `golang.org/x/crypto/sha3`: Includes the Keccak-256 hash function.
- `go-tron-keygen/base58`: Base58Check encoding and decoding.

## License

Go-Tron-Keygen is licensed under the MIT License.
