package main

import (
	"testing"
	"unique"
)

// Compare strings.

var (
	s1 = "just a decent length string"
	s2 = "just a decent length string"
)

func BenchmarkString(b *testing.B) {
	b.Run("regular", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = s1 == s2
		}
	})

	b.Run("unique", func(b *testing.B) {
		u1 := unique.Make(s1)
		u2 := unique.Make(s2)

		for i := 0; i < b.N; i++ {
			_ = u1 == u2
		}
	})
}

// Compare simple structs.

type Payload struct {
	A string
	B int
	C bool
	D float64
	E [2]int
}

func BenchmarkSimpleStruct(b *testing.B) {
	sp1 := Payload{A: "string", B: 42, C: true, D: 3.14, E: [2]int{1, 2}}
	sp2 := Payload{A: "string", B: 42, C: true, D: 3.14, E: [2]int{1, 2}}

	b.Run("regular", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = sp1 == sp2
		}
	})

	b.Run("unique", func(b *testing.B) {
		u1 := unique.Make(sp1)
		u2 := unique.Make(sp2)

		for i := 0; i < b.N; i++ {
			_ = u1 == u2
		}
	})
}

// Compare complex structs.

type Complex struct {
	A Payload
	B *Payload
	C [2]Payload
}

func BenchmarkComplexStruct(b *testing.B) {
	cp1 := Complex{
		A: Payload{A: "string", B: 42, C: true, D: 0.0006, E: [2]int{1, -43534512}},
		B: &Payload{A: "string", B: 42, C: true, D: 3.14, E: [2]int{54546456, 2}},
		C: [2]Payload{
			{A: "string", B: 42, C: true, D: -3547645, E: [2]int{1, 2}},
			{A: "string", B: 42, C: true, D: 3.14234234, E: [2]int{1, 2}},
		},
	}
	cp2 := cp1

	b.Run("regular", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = cp1 == cp2
		}
	})

	b.Run("unique", func(b *testing.B) {
		u1 := unique.Make(cp1)
		u2 := unique.Make(cp2)

		for i := 0; i < b.N; i++ {
			_ = u1 == u2
		}
	})
}
