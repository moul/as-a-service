package moul

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test(t *testing.T) {
	Convey("Testing package", t, func() {
		// RegisterAction
		// PlainResponse
		So(len(Actions()) > 0, ShouldBeTrue)
	})
}
