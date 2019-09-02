package chainResponsibility

import "testing"

func chainRepFactory() Manager {
	c1 := NewProjectManager()
	c2 := NewDepManager()
	c3 := NewGeneralManager()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	return c1
}

func TestChainResponse(t *testing.T) {
	c := chainRepFactory()

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 100000)
	c.HandleFeeRequest("floar", 400)
}
