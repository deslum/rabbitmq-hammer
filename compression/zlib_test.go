package compression

import (
	"reflect"
	"testing"
)

func TestZLib_Decode(t *testing.T) {
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
				data: []byte{120, 156, 1, 0, 0, 255, 255, 0, 0, 0, 1},
			},
			want:    []byte(""),
			wantErr: false,
		},
		{
			name: "Case 2: Decode string",
			args: args{
				data: []byte{120, 156, 50, 52, 50, 54, 49, 117, 116, 114, 6, 4, 0, 0, 255, 255, 7, 130, 1, 198},
			},
			want:    []byte("12345ABC"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ZLib{}
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

func TestZLib_Encode(t *testing.T) {
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
			want:    []byte{120, 156, 1, 0, 0, 255, 255, 0, 0, 0, 1},
			wantErr: false,
		},
		{
			name: "Case 2: Encode nil",
			args: args{
				data: nil,
			},
			want:    []byte{120, 156, 1, 0, 0, 255, 255, 0, 0, 0, 1},
			wantErr: false,
		},
		{
			name: "Case 3: Encode bytes string",
			args: args{
				data: []byte("12345ABC"),
			},
			want:    []byte{120, 156, 50, 52, 50, 54, 49, 117, 116, 114, 6, 4, 0, 0, 255, 255, 7, 130, 1, 198},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ZLib{}
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
