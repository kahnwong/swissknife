package cmd

import (
	"reflect"
	"testing"
)

func TestShouldIDeployToday(t *testing.T) {
	response := ShouldIDeployToday()

	if reflect.TypeOf(response.Message).Kind() != reflect.String {
		t.Errorf("ShouldIDeployToday() does not return a string")
	}
}
