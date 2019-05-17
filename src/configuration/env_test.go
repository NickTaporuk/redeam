package configuration

import (
	"os"
	"testing"
)

// init
func initTests(fakeEnv map[string]string, data ...string ) {

}

// nolint
func TestInitEnv(t *testing.T) {

	var envData = make(map[string]string)
	var env1 = make(map[string]string)
	env1[EnvNameDatabaseType] = "test"

	type args struct {
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		env map[string]string
	}{
		{
			name: "Test 1 positive",
			args: args{
				data: envData,
			},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			for key, val := range tt.env {
				err := os.Setenv(key, val)
				if err != nil {
					t.Fatal(err)
				}
			}

			if err := InitEnv(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("InitEnv() error = %v, wantErr %v", err, tt.wantErr)
			}

			os.Clearenv()
		})

	}
}
