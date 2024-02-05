package helper

import "testing"

func TestContainsInArrNoCaseSensitive(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in slice", args{
			s: []string{"exp"},
			e: "exp",
		}, true},
		{"not in slice", args{
			s: []string{"exp"},
			e: "exmp",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInArrNoCaseSensitive(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInArrNoCaseSensitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInArr(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"in slice", args{
			s: []string{"exp"},
			e: "exp",
		}, true},
		{"in slice", args{
			s: []string{"exp"},
			e: "exP",
		}, false},
		{"not in slice", args{
			s: []string{"exp"},
			e: "exmp",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInArr(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetweenString(t *testing.T) {
	type args struct {
		value string
		a     string
		b     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"get between", args{
			value: "string for get between",
			a:     "string",
			b:     "get",
		}, " for "},
		{"get between invalid var", args{
			value: "string for get between",
			a:     "one",
			b:     "string",
		}, ""},
		{"get between no var", args{
			value: "string for get between",
			a:     "",
			b:     "",
		}, ""},
		{"get between one var", args{
			value: "string for get between",
			a:     "string",
			b:     "",
		}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BetweenString(tt.args.value, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("BetweenString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeforeString(t *testing.T) {
	type args struct {
		value string
		a     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"get before", args{
			value: "string for get before",
			a:     "for",
		}, "string "},
		{"get before but empty", args{
			value: "string for get before",
			a:     "string",
		}, ""},
		{"get before no variable", args{
			value: "string for get before",
			a:     "",
		}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BeforeString(tt.args.value, tt.args.a); got != tt.want {
				t.Errorf("BeforeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfterString(t *testing.T) {
	type args struct {
		value string
		a     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"get after", args{
			value: "string for get after",
			a:     "for",
		}, " get after"},
		{"get after but empty", args{
			value: "string for get after",
			a:     "after",
		}, ""},
		{"get after no variable", args{
			value: "string for get after",
			a:     "",
		}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AfterString(tt.args.value, tt.args.a); got != tt.want {
				t.Errorf("AfterString() = %v, want %v", got, tt.want)
			}
		})
	}
}
