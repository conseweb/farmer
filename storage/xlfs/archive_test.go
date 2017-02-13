package xlfs

import (
	"testing"

	"gopkg.in/check.v1"
)

type XlfsSuite struct {
}

func init() {
	check.Suite(&XlfsSuite{})
}

func Test(t *testing.T) {
	check.TestingT(t)
}

func (x *XlfsSuite) TestArchive(c *check.C) {
}
