// +build !vartime

package edwards25519

import (
	"testing"

	"gitlab.neji.vm.tc/marconi/kyber"
)

func TestNotVartime(t *testing.T) {
	p := tSuite.Point()
	if _, ok := p.(kyber.AllowsVarTime); ok {
		t.Fatal("expected Point to NOT allow var time")
	}
}
