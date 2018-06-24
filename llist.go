
package llist

import (
	"strings"
	"fmt"
	"math/rand"
)

// START OMIT

type item struct {
	x    int
	next *item
}
// END OMIT

type List struct {
	items *item
}


func (s *List) Insert(x int) {
	var p **item
	for p = &(s.items); *p != nil; p = &((*p).next) {
		if (*p).x >= x {
			break
		}
	}
	it := &item{x, *p}
	*p = it
}

func (s *List)String()string {
	w := new(strings.Builder)
	for p := s.items; p != nil; p = p.next {

		fmt.Fprintf(w, "%d ", p.x)
	}
	return w.String()
}

// START2 OMIT
func (s *List)Sum()int {
	var t int
	for p := s.items; p != nil; p = p.next {
		t += p.x
	}
	return t
}
// END2 OMIT


func (s *List)Fill(n int) {
	for i:= 0; i < n; i++ {
		s.Insert(rand.Int())
	}
}

// STARTS OMIT
func sumV(v []int) int {
	t := 0
	for _, x := range v {
		t += x
	}
	return t
}
// ENDS OMIT

func fillV(v []int) {
	for i := 0; i < cap(v); i++ {
		v[i] = rand.Int()
	}
}
