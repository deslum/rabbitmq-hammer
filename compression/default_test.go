package compression

import (
	"reflect"
	"testing"
)

func TestDefault_Decode(t *testing.T) {
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
				data: []byte{},
			},
			want:    []byte(""),
			wantErr: false,
		},
		{
			name: "Case 2: Decode string",
			args: args{
				data: []byte("12345ABC"),
			},
			want:    []byte("12345ABC"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Default{}
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

func TestDefault_Encode(t *testing.T) {
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
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "Case 2: Encode nil",
			args: args{
				data: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Case 3: Encode bytes string",
			args: args{
				data: []byte("12345ABC"),
			},
			want:    []byte("12345ABC"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Default{}
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