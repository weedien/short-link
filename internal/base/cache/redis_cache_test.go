package cache

import "testing"

type student struct {
	name  string
	class string
}

func Test_isNilOrEmpty(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Nil value",
			args:    args{v: nil},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Empty string",
			args:    args{v: ""},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Non-empty string",
			args:    args{v: "test"},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Empty struct",
			args:    args{v: student{}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Empty struct pointer",
			args:    args{v: &student{}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Non-empty struct",
			args:    args{v: student{name: "test", class: "test"}},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Non-empty struct pointer",
			args:    args{v: &student{name: "test", class: "test"}},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Empty slice",
			args:    args{v: []int{}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Non-empty slice",
			args:    args{v: []int{1, 2, 3}},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Empty map",
			args:    args{v: map[string]int{}},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Non-empty map",
			args:    args{v: map[string]int{"key": 1}},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Nil pointer",
			args:    args{v: (*int)(nil)},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Non-nil pointer",
			args:    args{v: new(int)},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Unsupported type",
			args:    args{v: 123},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isNilOrEmpty(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("isNilOrEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isNilOrEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}
