package monsters

import (
	"reflect"
	"testing"
)

func TestMonsters_Add(t *testing.T) {
	type args struct {
		monsterName Monster
	}
	tests := []struct {
		name string
		m    *Monsters
		args args
		want *Monsters
	}{
		{
			name: "test-1",
			m:    &Monsters{"1"},
			args: args{"2"},
			want: &Monsters{"1", "2"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Add(tt.args.monsterName)
		})
		if !reflect.DeepEqual(tt.m, tt.want) {
			t.Errorf("monster.Add() = %v, want %v", tt.m, tt.want)
		}
	}
}

func TestMonsters_Delete(t *testing.T) {
	type args struct {
		monsterName Monster
	}
	tests := []struct {
		name string
		m    *Monsters
		args args
		want *Monsters
	}{
		{
			name: "test-1",
			m:    &Monsters{"1", "2"},
			args: args{"2"},
			want: &Monsters{"1"},
		},
		{
			name: "test-2",
			m:    &Monsters{"1", "2"},
			args: args{"1"},
			want: &Monsters{"2"},
		},
		{
			name: "test-3",
			m:    &Monsters{"1", "2", "3"},
			args: args{"2"},
			want: &Monsters{"1", "3"},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Delete(tt.args.monsterName)
		})
		if !reflect.DeepEqual(tt.m, tt.want) {
			t.Errorf("monster.Add() = %v, want %v", tt.m, tt.want)
		}
	}
}
