package configuration

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestCheckEnvVar(t *testing.T) {
	const (
		key  = "test"
		key1 = "test1"
	)
	var data = make(map[string]string)
	var err error

	data[key] = key

	err = os.Setenv(key, key)
	assert.NoError(t, err)

	type args struct {
		key  string
		data map[string]string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1 positive",
			args: args{
				key:  key,
				data: data,
			},
		},
		{
			name: "Test 2 negative",
			args: args{
				key:  key1,
				data: data,
			},
			wantErr: true,
		},
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

	var databaseConfig = &DatabaseConfig{
		Migrate: false,
		Seeds:   false,
	}

	var cnfDb = make(map[string]string)

	cnfDb[EnvNameDatabaseMigrate] = "false"
	cnfDb[EnvNameSeeds] = "false"

	type args struct {
		cnfData map[string]string
	}

	tests := []struct {
		name    string
		args    args
		want    *DatabaseConfig
		wantErr bool
	}{
		{
			name: "Test 1 positive",
			args: args{
				cnfData: cnfDb,
			},
			want: databaseConfig,
		},
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
	var cnf = &DatabaseConfig{}

	tests := []struct {
		name  string
		dbCnf *DatabaseConfig
		want  string
	}{
		{
			name:  "Test 1 positive",
			dbCnf: cnf,
			want:  "host= port= user= dbname= password= sslmode=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dbCnf.String(); got != tt.want {
				t.Errorf("DatabaseConfig.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
