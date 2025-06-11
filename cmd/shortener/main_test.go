package main

import "testing"

func TestSum(t *testing.T) {
	if sum := sum(1, 3); sum != 5 {
		t.Errorf("sum expected to be 5; got %d", sum)
	}
}