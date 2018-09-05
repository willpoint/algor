package experiments

import (
	"testing"
)

func TestForward(t *testing.T) {
	type arg struct {
		alphabet string
		forward  int
	}
	tests := []struct {
		name string
		arg  arg
		want string
	}{
		{"aForward3", arg{"a", 3}, "d"},
		{"zForward3", arg{"z", 3}, "c"},
		{",Forward3", arg{",", 3}, ","},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := forward(tt.arg.alphabet, tt.arg.forward); got != tt.want {
				t.Errorf("forward() wants %s, got %s", tt.want, got)
			}
		})
	}
}

func TestBackward(t *testing.T) {
	type arg struct {
		alphabet string
		backward int
	}
	tests := []struct {
		name string
		arg  arg
		want string
	}{
		{"aBackward3", arg{"a", 3}, "x"},
		{"aBackward3", arg{"z", 3}, "w"},
		{"?Backward3", arg{"?", 3}, "?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backward(tt.arg.alphabet, tt.arg.backward); got != tt.want {
				t.Errorf("backward() wants %s, got %s", tt.want, got)
			}
		})
	}
}

func Test_alphaToNum(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"testA", args{"a"}, 1, true},
		{"testZ", args{"z"}, 26, true},
		{"testComma", args{","}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := alphaToNum(tt.args.a)
			if got != tt.want {
				t.Errorf("alphaToNum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("alphaToNum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_numToAlpha(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{"test1", args{1}, "A", true},
		{"test24", args{24}, "X", true},
		{"test38", args{38}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := numToAlpha(tt.args.i)
			if got != tt.want {
				t.Errorf("numToAlpha() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("numToAlpha() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_mod(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"testPositive", args{3, 26}, 3},
		{"testNegative", args{-3, 26}, 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mod(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("mod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCipher(t *testing.T) {
	type args struct {
		alpha string
		step  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"testABC", args{"abc?", 3}, "def?"},
		{"testYZA", args{"yza,", 3}, "bcd,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cipher(tt.args.alpha, tt.args.step); got != tt.want {
				t.Errorf("Cipher() = %v, want %v", got, tt.want)
			}
		})
	}
}
