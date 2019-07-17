package suites

import (
	"github.com/MarconiProtocol/kyber/group/edwards25519"
)

func init() {
	register(edwards25519.NewBlakeSHA256Ed25519())
}
