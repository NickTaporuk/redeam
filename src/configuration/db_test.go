package configuration

import (
	"reflect"
	"testing"
)

func TestCheckEnvVar(t *testing.T) {
	type args struct {
		key  string
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
			if err := CheckEnvVar(tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("CheckEnvVar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDatabaseConfig(t *testing.T) {
	type args struct {
		cnfData map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *DatabaseConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDatabaseConfig(tt.args.cnfData)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDatabaseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseConfig_String(t *testing.T) {
	tests := []struct {
		name  string
		dbCnf *DatabaseConfig
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dbCnf.String(); got != tt.want {
				t.Errorf("DatabaseConfig.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
