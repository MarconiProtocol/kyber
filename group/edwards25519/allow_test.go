// +build !vartime

package edwards25519

import (
	"testing"

	"github.com/MarconiProtocol/kyber"
)

func TestNotVartime(t *testing.T) {
	p := tSuite.Point()
	if _, ok := p.(kyber.AllowsVarTime); ok {
		t.Fatal("expected Point to NOT allow var time")
	}
}
