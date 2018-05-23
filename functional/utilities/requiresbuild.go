package testutilities

import (
	"testing"

	"github.com/Microsoft/hcsshim/internal/osversion"
)

func RequiresBuild(t *testing.T, b uint16) {
	if osversion.GetOSVersion().Build < b {
		t.Skipf("Requires build %d+", b)
	}
}
