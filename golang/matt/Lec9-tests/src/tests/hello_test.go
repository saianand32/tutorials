package tests

import (
	"testHello/src/hello"
	"testing"
)

// commands to run the tests - go test ./src/tests -v   //v is for verbose
// commands to run the benchmark - go test ./src/tests -bench=.

func TestSayHello(t *testing.T) {
	want := "Hiii saishwar"
	got := hello.Say("saishwar")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
func TestSayHelloNultiple(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "saishwar",
			want: "Hiii saishwar",
		},
		{
			name: "anand",
			want: "Hiii anand",
		},
		{
			name: "sai",
			want: "Hiii sai",
		},
	}

	for _, tt := range tests {
		if got := hello.Say(tt.name); got != tt.want {
			t.Errorf("wanted %s, got %s", tt.want, got)
		}
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hello.Say("saishwar")
	}
}
