package bls12381

import (
	"crypto/subtle"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	blst "github.com/supranational/blst/bindings/go"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"io"
)

const (
	PrivKeyName = "tendermint/PrivKeyBls12381"
	PubKeyName  = "tendermint/PubKeyBls12381"
	// PubKeySize is is the size, in bytes, of public keys as used in this package.
	PubKeySize = 96
	// PrivKeySize is the size, in bytes, of private keys as used in this package.
	// Uncompressed public key
	PrivKeySize = 32
	// Compressed signature
	SignatureSize = 96
	keyType       = "bls12381"
	SeedSize      = 32
)

var _ cryptotypes.PrivKey = &PrivKey{}
var _ codec.AminoMarshaler = &PrivKey{}

// Bytes returns the byte representation of the Private Key.
func (privKey *PrivKey) Bytes() []byte {
	return privKey.Key
}

// PubKey performs the point-scalar multiplication from the privKey on the
// generator point to get the pubkey.
func (privKey *PrivKey) PubKey() cryptotypes.PubKey {
	sk := new(blst.SecretKey).Deserialize(privKey.Key)
	if sk == nil {
		panic("Failed to deserialize secret key!")
	}
	pk := new(blst.P1Affine).From(sk)
	pk_bytes := pk.Serialize()

	return &PubKey{Key: pk_bytes}
}

// Sign produces a signature on the provided message.
// This assumes the privkey is wellformed in the golang format.

func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	dst := []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")
	sk := new(blst.SecretKey).Deserialize(privKey.Key)
	if sk == nil {
		panic("Failed to deserialize secret key!")
	}

	sig := new(blst.P2Affine).Sign(sk, msg, dst)
	if sig == nil {
		panic("Failed to sign message!")
	}

	sig_bytes := sig.Compress()

	return sig_bytes, nil
}

// Equals - you probably don't need to use this.
// Runs in constant time based on length of the
func (privKey *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	return privKey.Type() == other.Type() && subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

func (privKey *PrivKey) Type() string {
	return keyType
}

// MarshalAmino overrides Amino binary marshalling.
func (privKey PrivKey) MarshalAmino() ([]byte, error) {
	return privKey.Key, nil
}

// UnmarshalAmino overrides Amino binary marshalling.
func (privKey *PrivKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PrivKeySize {
		return fmt.Errorf("invalid privkey size")
	}
	privKey.Key = bz

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (privKey PrivKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return privKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (privKey *PrivKey) UnmarshalAminoJSON(bz []byte) error {
	return privKey.UnmarshalAmino(bz)
}

// GenPrivKey generates a new BLS private key on curve bls12-381 private key.
// It uses OS randomness to generate the private key.
func GenPrivKey() *PrivKey {
	return &PrivKey{Key: genPrivKey(crypto.CReader())}
}

// genPrivKey generates a new secp256k1 private key using the provided reader.
func genPrivKey(rand io.Reader) []byte {
	var ikm [SeedSize]byte
	_, err := io.ReadFull(rand, ikm[:])
	if err != nil {
		panic(err)
	}

	sk := blst.KeyGen(ikm[:])
	if sk == nil {
		panic("Failed to generate secret key!")
	}

	sk_bytes := sk.Serialize()

	return sk_bytes
}

// GenPrivKeyFromSecret hashes the secret with SHA2, and uses
// that 32 byte output to create the private key.
// NOTE: secret should be the output of a KDF like bcrypt,
// if it's derived from user input.
func GenPrivKeyFromSecret(secret []byte) *PrivKey {
	ikm := crypto.Sha256(secret) // Not Ripemd160 because we want 32 bytes.
	sk_bytes := blst.KeyGen(ikm[:]).Serialize()

	return &PrivKey{Key: sk_bytes}
}

var _ cryptotypes.PubKey = &PubKey{}
var _ codec.AminoMarshaler = &PubKey{}

// Validate public key, infinity and subgroup checking
func (pubKey *PubKey) Validate() bool {
	pk := new(blst.P1Affine).Deserialize(pubKey.Key)
	return pk.KeyValidate()
}

// Address is the SHA256-20 of the raw pubkey bytes.
func (pubKey *PubKey) Address() crypto.Address {
	if len(pubKey.Key) != PubKeySize {
		panic("pubkey is incorrect size")
	}
	return crypto.Address(tmhash.SumTruncated(pubKey.Key))
}

// Bytes returns the PubKey byte format.
func (pubKey *PubKey) Bytes() []byte {
	return pubKey.Key
}

// Assume public key is validated
func (pubKey *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	// make sure we use the same algorithm to sign
	pk := new(blst.P1Affine).Deserialize(pubKey.Key)
	if pk == nil {
		panic("Failed to deserialize public key")
	}

	sigma := new(blst.P2Affine).Uncompress(sig)
	if sigma == nil {
		panic("Failed to deserialize signature")
	}

	dst := []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")

	return sigma.Verify(true, pk, false, msg, dst)
}

func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeyBls12381{%X}", pubKey.Key)
}

func (pubKey *PubKey) Type() string {
	return keyType
}

func (pubKey *PubKey) Equals(other cryptotypes.PubKey) bool {
	if pubKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(pubKey.Bytes(), other.Bytes()) == 1
}

// MarshalAmino overrides Amino binary marshalling.
func (pubKey PubKey) MarshalAmino() ([]byte, error) {
	return pubKey.Key, nil
}

// UnmarshalAmino overrides Amino binary marshalling.
func (pubKey *PubKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PubKeySize {
		return errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size")
	}
	pubKey.Key = bz

	return nil
}

// MarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey PubKey) MarshalAminoJSON() ([]byte, error) {
	// When we marshal to Amino JSON, we don't marshal the "key" field itself,
	// just its contents (i.e. the key bytes).
	return pubKey.MarshalAmino()
}

// UnmarshalAminoJSON overrides Amino JSON marshalling.
func (pubKey *PubKey) UnmarshalAminoJSON(bz []byte) error {
	return pubKey.UnmarshalAmino(bz)
}
