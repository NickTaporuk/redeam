package utils

import "testing"

//nolint
func TestEnvVarExists(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		wantValue string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := EnvVarExists(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("EnvVarExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("EnvVarExists() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
