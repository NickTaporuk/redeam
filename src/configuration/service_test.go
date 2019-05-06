package configuration

import (
	"reflect"
	"testing"
)

func TestNewServiceConfig(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name string
		args args
		want *ServiceConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServiceConfig(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServiceConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
