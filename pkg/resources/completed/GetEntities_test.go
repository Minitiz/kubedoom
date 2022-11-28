package completed

import (
	"reflect"
	"testing"
)

func TestMode_GetEntities(t *testing.T) {
	type fields struct {
		Label   string
		Respawn bool
	}
	type args struct {
		MONSTERS []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "test-1",
			fields: fields{},
			args:   args{},
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mode{
				Label:   tt.fields.Label,
				Respawn: tt.fields.Respawn,
			}
			if got := m.GetEntities(tt.args.MONSTERS); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mode.GetEntities() = %v, want %v", got, tt.want)
			}
		})
	}
}
