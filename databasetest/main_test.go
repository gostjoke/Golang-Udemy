package main

import "testing"

func TestSaveUser(t *testing.T) {
	md := MockDatabse{
		Users: map[int]User{
			2: {ID: 2, First: "Alice"},
		},
	}

	s := &Service{ds: md}

	u, err := s.GetUser(2)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.First != "Alice" {
		t.Fatalf("expected user with First name 'Alice', got %v", u.First)
	}
}

//go test -bench=.
func BenchmarkSaveUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md := MockDatabse{
			Users: map[int]User{
				2: {ID: 2, First: "Alice"},
			},
		}
		s := &Service{ds: md}
		u, err := s.GetUser(2)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
		if u.First != "Alice" {
			b.Fatalf("unexpected user first name: %v", u.First)
		}
	}
}
