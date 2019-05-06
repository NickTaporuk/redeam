package configuration

import "testing"

func TestInitEnv(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitEnv(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("InitEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
