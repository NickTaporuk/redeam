package db

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestDatabaseConfig_String(t *testing.T) {
	type fields struct {
		SslMode          string
		DatabaseHost     string
		DatabaseType     string
		DatabaseUser     string
		DatabaseName     string
		DatabasePassword string
		DatabasePort     string
		Migrate          bool
		Seeds            bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbCnf := DatabaseConfig{
				SslMode:          tt.fields.SslMode,
				DatabaseHost:     tt.fields.DatabaseHost,
				DatabaseType:     tt.fields.DatabaseType,
				DatabaseUser:     tt.fields.DatabaseUser,
				DatabaseName:     tt.fields.DatabaseName,
				DatabasePassword: tt.fields.DatabasePassword,
				DatabasePort:     tt.fields.DatabasePort,
				Migrate:          tt.fields.Migrate,
				Seeds:            tt.fields.Seeds,
			}
			if got := dbCnf.String(); got != tt.want {
				t.Errorf("DatabaseConfig.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConnection(t *testing.T) {
	type args struct {
		dbCnf *DatabaseConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConnection(tt.args.dbCnf)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
