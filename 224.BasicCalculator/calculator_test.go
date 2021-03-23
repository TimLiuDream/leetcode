package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	s := "1 + 1"
	assert.Equal(t, 2, calculate(s))
	s = " 2-1 + 2 "
	assert.Equal(t, 3, calculate(s))
	s = "(1+(4+5+2)-3)+(6+8)"
	assert.Equal(t, 23, calculate(s))
}

func calculate(s string) int {
	var s1 = rune('(')
	var s2 = rune(')')
	s = strings.TrimSpace(s)
	stack := make([]rune, 0)
	for _, r := range s {
		if r == s1 {
			sta
		}
		if len(stack) == 0 {
			stack = append(stack, r)
		} else {
			if r == s2 {

			}
		}
	}
	return 1
}
