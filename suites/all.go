package suites

import (
	"gitlab.neji.vm.tc/marconi/kyber/group/edwards25519"
)

func init() {
	register(edwards25519.NewBlakeSHA256Ed25519())
}
