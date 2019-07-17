// +build vartime

package suites

import (
	"github.com/MarconiProtocol/kyber/group/curve25519"
	"github.com/MarconiProtocol/kyber/group/nist"
	"github.com/MarconiProtocol/kyber/pairing/bn256"
)

func init() {
	register(curve25519.NewBlakeSHA256Curve25519(false))
	register(curve25519.NewBlakeSHA256Curve25519(true))
	register(nist.NewBlakeSHA256P256())
	register(nist.NewBlakeSHA256QR512())
	register(bn256.NewSuiteG1())
	register(bn256.NewSuiteG2())
	register(bn256.NewSuiteGT())
}
