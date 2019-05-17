package configuration

import (
	"reflect"
	"testing"
)

//nolint
func TestNewConfig(t *testing.T) {
	var dbConf = &DatabaseConfig{
		Migrate:          false,
		Seeds:            false,
		SslMode:          "disabled",
		DatabaseHost:     "0.0.0.0",
		DatabaseType:     "postgres",
		DatabasePort:     ":8888",
		DatabasePassword: "redeam",
		DatabaseName:     "redeam",
		DatabaseUser:     "redeam",
	}

	var serviceConf = &ServiceConfig{}

	type args struct {
		databaseConfig *DatabaseConfig
		serviceConfig  *ServiceConfig
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "Test 1 positive",
			args: args{
				databaseConfig: dbConf,
				serviceConfig:  serviceConf,
			},
			want: NewConfig(dbConf, serviceConf),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.databaseConfig, tt.args.serviceConfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
