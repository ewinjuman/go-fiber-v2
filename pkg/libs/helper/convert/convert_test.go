package convert

import "testing"

type InStruct struct {
	ID string `json:"id"`
}

var formData map[string]string

func TestObjectToObject(t *testing.T) {
	type args struct {
		in  interface{}
		out interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"Success",
			args{
				in:  InStruct{ID: "1"},
				out: "{ID:1}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ObjectToObject(tt.args.in, tt.args.out)
		})
	}
}

func TestObjectToString(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success 1",
			args{
				data: InStruct{ID: "1"},
			},
			"{\"id\":\"1\"}",
		},
		{"Success 2",
			args{
				data: InStruct{ID: "2"},
			},
			"{\"id\":\"2\"}",
		},
		{"Failed",
			args{
				data: make(chan int),
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjectToString(tt.args.data); got != tt.want {
				t.Errorf("ObjectToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToObject(t *testing.T) {
	type args struct {
		in  string
		out interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"Success 1",
			args{
				in:  "{\"id\":\"1\"}",
				out: InStruct{ID: "1"},
			},
		},
		{"Success 2",
			args{
				in:  "{\"id\":\"2\"}",
				out: InStruct{ID: "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringToObject(tt.args.in, tt.args.out)
		})
	}
}
