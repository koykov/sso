package sso

import "testing"

func TestSSO(t *testing.T) {
	t.Run("assignByteseq/short", func(t *testing.T) {
		s := New("foobar")
		if s.String() != "foobar" {
			t.Fail()
		}
	})
	t.Run("assignByteseq/long", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
	t.Run("append/short+short", func(t *testing.T) {
		s := New("hello")
		s.Append(" world!")
		if s.String() != "hello world!" {
			t.Fail()
		}
	})
	t.Run("append/short+long", func(t *testing.T) {
		s := New("Lorem ipsum")
		s.Append(" dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
	t.Run("append/long+long", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet,")
		s.Append(" consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
}

func BenchmarkSSO(b *testing.B) {
	b.Run("assignByteseq/short", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("foobar")
			_ = s
		}
	})
	b.Run("assignByteseq/long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
			_ = s
		}
	})
	b.Run("append/short+short", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("hello")
			s.Append(" world!")
		}
	})
	b.Run("append/short+long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum")
			s.Append(" dolor sit amet, consectetur adipiscing elit.")
		}
	})
	b.Run("append/long+long", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet,")
			s.Append(" consectetur adipiscing elit.")
		}
	})
}
