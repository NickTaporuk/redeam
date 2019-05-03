package app

import "testing"

func TestMain_Run(t *testing.T) {
	tests := []struct {
		name    string
		m       *Main
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Main.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMain_Close(t *testing.T) {
	tests := []struct {
		name string
		m    *Main
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Close()
		})
	}
}
