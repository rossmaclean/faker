package healthcore

import (
	healthcoremodel "faker/internal/health/core/model"
	"reflect"
	"testing"
)

func TestGetHealth(t *testing.T) {
	tests := []struct {
		name    string
		want    healthcoremodel.HealthResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHealth()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHealth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHealth() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getDatabaseHealth(t *testing.T) {
	tests := []struct {
		name string
		want healthcoremodel.HealthStatus
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDatabaseHealth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDatabaseHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
