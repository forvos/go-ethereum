package vm

import "testing"

func TestStatck(t *testing.T) {
	st := newstack()
	t.Log(st.len())

}
