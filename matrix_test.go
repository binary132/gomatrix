package matrix_test

import gc "gopkg.in/check.v1"

type matrixSuite struct{}

var _ = gc.Suite(matrixSuite{})

func (s *matrixSuite) TestInit(c *gc.C) {
	c.Assert(true, gc.Equals, true)
}
