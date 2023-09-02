package main

import "testing"

func TestStringFlags_Set(t *testing.T) {
	cc := []struct {
		name string
		vv   []string
		out  string
		err  error
	}{
		{
			name: "single value: a",
			vv:   []string{"a"},
			out:  "a",
		},
		{
			name: "three values: a, b, c",
			vv:   []string{"a", "b", "c"},
			out:  "a b c",
		},
	}

	for _, c := range cc {
		t.Run(c.name, func(t *testing.T) {
			f := &StringFlags{}

			for _, v := range c.vv {
				if err := f.Set(v); err != c.err {
					t.Errorf("received error %v and expected %v", err, c.err)
				}
			}

			if f.String() != c.out {
				t.Errorf("expected output %s and received %s", c.out, f.String())
			}
		})
	}
}
