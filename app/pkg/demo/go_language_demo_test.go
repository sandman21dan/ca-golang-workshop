package demo_test

import (
	"fmt"
	"testing"
)

func TestGoDemo(t *testing.T) {
	t.Run("Variables", func(t *testing.T) {
		// declare and assign an int
		var a int
		a = 200

		// also declare and assign an int
		b := 100

		// declare and assign an array of strings
		var l []string
		l = []string{"a", "b", "c"}

		// also declare and assing an array of strings
		l2 := []string{"a", "b", "c"}

		// declare and assign a map
		m := map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
		}

		fmt.Println(a, b, l, l2, m)
	})

	t.Run("More values", func(t *testing.T) {
		// values are not nullable
		var a int    // implicitly 0
		var s string // implicitly ""
		var b bool   // implicitly false

		fmt.Println(a, s, b)

		// Only pointers are nullable
		var p *int // points nowhere AKA nil

		if p == nil {
			fmt.Println("p is not set")
		}

		v := 100
		// assign p to point to the value of v
		p = &v

		if p != nil {
			fmt.Println("p is set and now points to value", *p)
		}
	})

	t.Run("Conditionals", func(t *testing.T) {
		cond := true
		otherCond := false

		if cond {
			fmt.Println("cond is true")
		} else if otherCond {
			fmt.Println("otherCond is true")
		} else {
			fmt.Println("cond is false")
		}

		switch cond {
		case true:
			fmt.Println("cond is true")
			// no need to break
		case false:
			fmt.Println("cond is false")
		default:
			fmt.Println("cond is neither true nor false ????")
		}
	})

	t.Run("Loops", func(t *testing.T) {
		// create an array from 0 to n
		l := []int{}
		n := 10

		for i := 0; i < n; i++ {
			l = append(l, i)
		}

		fmt.Println("l is", l)

		// map an array to another
		l2 := []int{}

		for _, i := range l {
			l2 = append(l2, i*2)
		}

		fmt.Println("l2 is", l2)
	})
}
