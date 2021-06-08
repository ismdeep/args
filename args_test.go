package args

import (
	"os"
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	os.Args = append(os.Args, "--help")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{"--append"},
			want: false,
		},
		{
			name: "Test 2",
			args: args{"--help"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.key); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	os.Args = append(os.Args, "-n")
	os.Args = append(os.Args, "100")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{"-n"},
			want: "100",
		},
		{
			name: "Test 2",
			args: args{"-l"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.key); got != tt.want {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	os.Args = append(os.Args, "-n")
	os.Args = append(os.Args, "100")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{"-n"},
			want: 100,
		},
		{
			name: "Test 2",
			args: args{"-l"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.key); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueList(t *testing.T) {
	os.Args = append(os.Args, "-d")
	os.Args = append(os.Args, "dir1")
	os.Args = append(os.Args, "-d")
	os.Args = append(os.Args, "dir2")

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestGetValueList-001",
			args: args{
				key: "-d",
			},
			want: []string{"dir1", "dir2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueList(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueList() = %v, want %v", got, tt.want)
			}
		})
	}
}
