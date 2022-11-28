package ingress

import "testing"

func TestMode_Format(t *testing.T) {
	type fields struct {
		Label   string
		Respawn bool
	}
	type args struct {
		entity string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{},
			args: args{
				entity: "todd-staging/todd-nginx",
			},
			want: "todd-nginx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Mode{
				Label:   tt.fields.Label,
				Respawn: tt.fields.Respawn,
			}
			if got := m.Format(tt.args.entity); got != tt.want {
				t.Errorf("Mode.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
