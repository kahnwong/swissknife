package cmd

import (
	"reflect"
	"testing"
)

func TestShouldIDeployToday(t *testing.T) {
	response, err := ShouldIDeployToday()
	if err != nil {
		t.Errorf("ShouldIDeployToday() error = %v", err)

	} else if reflect.TypeOf(response.Message).Kind() != reflect.String {
		t.Errorf("ShouldIDeployToday() does not return a string")
	}
}
