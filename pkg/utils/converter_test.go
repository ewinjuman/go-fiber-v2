package utils

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

func TestConvertPhoneNumber(t *testing.T) {
	type args struct {
		mobilePhoneNumber string
	}
	tests := []struct {
		name                     string
		args                     args
		wantNewMobilePhoneNumber string
		wantErr                  bool
	}{
		{
			"62 test",
			args{mobilePhoneNumber: "6281272702504"},
			"081272702504",
			false,
		},
		{
			"+62 test",
			args{mobilePhoneNumber: "+6281272702504"},
			"081272702504",
			false,
		},
		{
			"0 test",
			args{mobilePhoneNumber: "081272702504"},
			"081272702504",
			false,
		},
		{
			"no prefix test",
			args{mobilePhoneNumber: "81272702504"},
			"081272702504",
			false,
		},
		{
			"invalid number",
			args{mobilePhoneNumber: "8127270250n"},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewMobilePhoneNumber, err := ConvertPhoneNumber(tt.args.mobilePhoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertPhoneNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNewMobilePhoneNumber != tt.wantNewMobilePhoneNumber {
				t.Errorf("ConvertPhoneNumber() gotNewMobilePhoneNumber = %v, want %v", gotNewMobilePhoneNumber, tt.wantNewMobilePhoneNumber)
			}
		})
	}
}
