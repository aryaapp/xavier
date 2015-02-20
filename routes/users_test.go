package routes

import (
	"io"
	"testing"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
	c.Assert(42, Equals, "42")
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}

func Test(t *testing.T) {
	TestingT(t)
}
