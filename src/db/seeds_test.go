// nolint
package db

import (
	"testing"

	"github.com/NickTaporuk/redeam/src/models"
	"github.com/jinzhu/gorm"
)

func TestDatabaseMigrateSeeds(t *testing.T) {
	type args struct {
		db  *gorm.DB
		mds models.RedeamModels
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
			if err := DatabaseMigrateSeeds(tt.args.db, tt.args.mds); (err != nil) != tt.wantErr {
				t.Errorf("DatabaseMigrateSeeds() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
