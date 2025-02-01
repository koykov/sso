package sso

import "testing"

func TestSSO(t *testing.T) {
	t.Run("copy/sso", func(t *testing.T) {
		s := New("foobar")
		if s.String() != "foobar" {
			t.Fail()
		}
	})
	t.Run("copy/heap", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
	t.Run("concat/sso", func(t *testing.T) {
		s := New("hello")
		s.Concat(" world!")
		if s.String() != "hello world!" {
			t.Fail()
		}
	})
	t.Run("concat/sso_to_heap", func(t *testing.T) {
		s := New("Lorem ipsum")
		s.Concat(" dolor sit amet, consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
	t.Run("concat/heap", func(t *testing.T) {
		s := New("Lorem ipsum dolor sit amet,")
		s.Concat(" consectetur adipiscing elit.")
		if s.String() != "Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
			t.Fail()
		}
	})
}

func BenchmarkSSO(b *testing.B) {
	b.Run("copy/sso", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("foobar")
			_ = s
		}
	})
	b.Run("copy/heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")
			_ = s
		}
	})
	b.Run("concat/sso", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("hello")
			s.Concat(" world!")
		}
	})
	b.Run("concat/sso_to_heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum")
			s.Concat(" dolor sit amet, consectetur adipiscing elit.")
		}
	})
	b.Run("concat/heap", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := New("Lorem ipsum dolor sit amet,")
			s.Concat(" consectetur adipiscing elit.")
		}
	})
}

func BenchmarkString(b *testing.B) {
	b.Run("copy", func(b *testing.B) {
		b.ReportAllocs()
		x := "foobar"
		for i := 0; i < b.N; i++ {
			s := x + ""
			_ = s
		}
	})
	b.Run("concat", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := "hello"
			s = s + " world!"
		}
	})
	b.Run("concat", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			s := "Lorem ipsum dolor sit amet,"
			s = s + " consectetur adipiscing elit."
		}
	})
}
