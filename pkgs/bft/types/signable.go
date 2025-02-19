package types

import (
	"github.com/gnolang/gno/pkgs/crypto/ed25519"
	"github.com/gnolang/gno/pkgs/maths"
)

// MaxSignatureSize is a maximum allowed signature size for the Proposal
// and Vote.
// XXX: secp256k1 does not have Size nor MaxSize defined.
var MaxSignatureSize = maths.MaxInt(ed25519.SignatureSize, 64)

// Signable is an interface for all signable things.
// It typically removes signatures before serializing.
// SignBytes returns the bytes to be signed
// NOTE: chainIDs are part of the SignBytes but not
// necessarily the object themselves.
// NOTE: Expected to panic if there is an error marshalling.
type Signable interface {
	SignBytes(chainID string) []byte
}
