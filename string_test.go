package sso

import "testing"

func TestSSO(t *testing.T) {
	t.Run("assign/short", func(t *testing.T) {
		s := New("foobar")
		if s.String() != "foobar" {
			t.Fail()
		}
	})
	t.Run("assign/long", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
}

func BenchmarkSSO(b *testing.B) {
	b.Run("assign/short", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("foobar")
			_ = s
		}
	})
	b.Run("assign/long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
			_ = s
		}
	})
}
