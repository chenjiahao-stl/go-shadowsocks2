package socks

import (
	"fmt"
	"testing"
)

func TestParseAddr(t *testing.T) {
	var target = "8.8.8.8:53"
	tgt := ParseAddr(target)
	fmt.Println(tgt)
}
