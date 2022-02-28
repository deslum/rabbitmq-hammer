package compression

import (
	"reflect"
	"testing"
)

func TestLZ4_Decode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1: Decode empty string",
			args: args{
				data: []byte{4, 34, 77, 24, 100, 112, 185, 0, 0, 0, 0, 5, 93, 204, 2},
			},
			want:    []byte(""),
			wantErr: false,
		},
		{
			name: "Case 2: Decode string",
			args: args{
				data: []byte{4, 34, 77, 24, 100, 112, 185, 8, 0, 0, 128, 49, 50, 51, 52, 53, 65, 66, 67, 0, 0, 0, 0, 206, 104, 128, 3},
			},
			want:    []byte("12345ABC"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LZ4{}
			got, err := o.Decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLZ4_Encode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1: Encode empty",
			args: args{
				data: []byte(""),
			},
			want:    []byte{4, 34, 77, 24, 100, 112, 185, 0, 0, 0, 0, 5, 93, 204, 2},
			wantErr: false,
		},
		{
			name: "Case 2: Encode nil",
			args: args{
				data: nil,
			},
			want:    []byte{4, 34, 77, 24, 100, 112, 185, 0, 0, 0, 0, 5, 93, 204, 2},
			wantErr: false,
		},
		{
			name: "Case 3: Encode bytes string",
			args: args{
				data: []byte("12345ABC"),
			},
			want:    []byte{4, 34, 77, 24, 100, 112, 185, 8, 0, 0, 128, 49, 50, 51, 52, 53, 65, 66, 67, 0, 0, 0, 0, 206, 104, 128, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LZ4{}
			got, err := o.Encode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}
