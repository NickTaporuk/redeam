package seeds

import (
	"reflect"
	"testing"

	"github.com/NickTaporuk/redeam/src/models"
)

//nolint
func TestSeeds(t *testing.T) {
	tests := []struct {
		name string
		want models.RedeamModels
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seeds(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Seeds() = %v, want %v", got, tt.want)
			}
		})
	}
}
