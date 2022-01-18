package main

import (
	"reflect"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	tests := []struct {
		name           string
		names, numbers []string
		want           map[string]string
	}{
		{
			name:    "empty",
			names:   []string{},
			numbers: []string{},
			want:    map[string]string{},
		},
		{
			name:    "name without a phone number",
			names:   []string{"Aleksej"},
			numbers: []string{"0"},
			want:    map[string]string{"Aleksej": "0"},
		},
		{
			name:    "number without a name",
			names:   []string{""},
			numbers: []string{"+37125542501"},
			want:    map[string]string{"": "+37125542501"},
		},
		{
			name:    "name and number",
			names:   []string{"Aleksej"},
			numbers: []string{"+37125542501"},
			want:    map[string]string{"Aleksej": "+37125542501"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := CreatePhoneBook(tc.names, tc.numbers); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, tc.want)
			}
		})
	}
}
