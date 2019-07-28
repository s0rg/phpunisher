package scanners

import (
	"testing"
)

func TestScannerStub(t *testing.T) {

	const n = "test-name"

	s := stub{n}

	if s.Name() != n {
		t.Fatal("names not match")
	}
}
