//go:binary-only-package

package ngapType

// Need to import "free5gc/lib/aper" if it uses "aper"

type AMFUENGAPID struct {
	Value int64 `aper:"valueLB:0,valueUB:1099511627775"`
}
