package sso

import "testing"

func TestSSO(t *testing.T) {
	t.Run("assign/sso", func(t *testing.T) {
		s := New("foobar")
		if s.String() != "foobar" {
			t.Fail()
		}
	})
	t.Run("assign/heap", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
	t.Run("append/sso", func(t *testing.T) {
		s := New("hello")
		s.AppendString(" world!")
		if s.String() != "hello world!" {
			t.Fail()
		}
	})
	t.Run("append/sso_to_heap", func(t *testing.T) {
		s := New("Lorem ipsum")
		s.AppendString(" dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
	t.Run("append/heap", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet,")
		s.AppendString(" consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
}

func BenchmarkSSO(b *testing.B) {
	b.Run("assign/sso", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("foobar")
			_ = s
		}
	})
	b.Run("assign/heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
			_ = s
		}
	})
	b.Run("append/sso", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("hello")
			s.AppendString(" world!")
		}
	})
	b.Run("append/sso_to_heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum")
			s.AppendString(" dolor sit amet, consectetur adipiscing elit.")
		}
	})
	b.Run("append/heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet,")
			s.AppendString(" consectetur adipiscing elit.")
		}
	})
}
