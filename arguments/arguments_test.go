package arguments

import (
	"errors"
	"reflect"
	"testing"
)

func TestValidateArg(t *testing.T) {
	parallelFlagNameInArgs := "sampleflag"
	invalidArgError := errors.New("number of arguments are invalid")

	validationErrorStr := "ValidateArgs(%v,%v); want <%v>, got <%v>"
	tests := []struct {
		input  []string
		output error
	}{
		{[]string{}, invalidArgError},
		{[]string{""}, nil},
		{[]string{"google.com", "fb.com"}, nil},
		{[]string{"-sampleflag", "1"}, invalidArgError},
		{[]string{"-sampleflag", "1", "google.com"}, nil},
		{[]string{"-sampleflag", "1", "google.com", "fb.com"}, nil},
		{[]string{"-sampleflag=1", "google.com", "fb.com"}, nil},
	}

	for _, v := range tests {
		err := validateArgs(v.input, parallelFlagNameInArgs)
		if v.output != nil && err != nil && v.output.Error() != err.Error() { // if both error expected and received not nil
			t.Errorf(validationErrorStr, v.input, parallelFlagNameInArgs, v.output.Error(), err.Error())
		} else if v.output == nil && err != nil { // if expected nil and received not nil
			t.Errorf(validationErrorStr, v.input, parallelFlagNameInArgs, v.output, err)
		} else if v.output != nil && err == nil { // if expected not nil and received nil
			t.Errorf(validationErrorStr, v.input, parallelFlagNameInArgs, v.output, err)
		}
	}
}

func TestGetSitesFromArgs(t *testing.T) {
	parallelFlagNameInArgs := "sampleflag"

	validationErrorStr := "GetSitesFromArgs(%v,%v); want <%v>, got <%v>"
	tests := []struct {
		input  []string
		output []string
	}{
		{[]string{}, []string{}},
		{[]string{""}, nil},
		{[]string{"google.com", "fb.com"}, []string{"google.com", "fb.com"}},
		{[]string{"-sampleflag", "1"}, nil},
		{[]string{"-sampleflag", "1", "google.com"}, []string{"google.com"}},
		{[]string{"-sampleflag", "1", "google.com", "fb.com"}, []string{"google.com", "fb.com"}},
		{[]string{"-sampleflag=1", "google.com", "fb.com"}, []string{"google.com", "fb.com"}},
	}

	for _, v := range tests {
		sites, _ := GetSitesFromArgs(v.input, parallelFlagNameInArgs)
		if v.output != nil && sites != nil && !reflect.DeepEqual(v.output, sites) {
			t.Errorf(validationErrorStr, v.input, parallelFlagNameInArgs, v.output, sites)
		}
	}
}
