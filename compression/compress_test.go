package compression

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCompressType_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		result  string
		wantErr bool
	}{
		{
			name:    "Case1: None",
			arg:     "None",
			result:  "None",
			wantErr: false,
		},
		{
			name:    "Case2: ZLib",
			arg:     "ZLib",
			result:  "ZLib",
			wantErr: false,
		},
		{
			name:    "Case3: LZ4",
			arg:     "LZ4",
			result:  "LZ4",
			wantErr: false,
		},
		{
			name:    "Case4: Snappy",
			arg:     "Snappy",
			result:  "Snappy",
			wantErr: false,
		},
		{
			name:    "Case4: Not support compress type",
			arg:     "ABCDEF",
			result:  "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compressType, _ := json.Marshal(tt.arg)

			o := new(CompressType)
			err := o.UnmarshalJSON(compressType)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			case err == nil && !reflect.DeepEqual(o.String(), tt.result):
				t.Errorf("UnmarshalJSON() got = %v, want %v", o.String(), tt.result)
			}
		})
	}
}
