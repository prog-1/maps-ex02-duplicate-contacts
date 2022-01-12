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
			name:    "single letter",
			names:   []string{"a"},
			numbers: []string{"0"},
			want:    map[string]string{"a": "0"},
		},
		{
			name:    "no name, only phone number",
			names:   []string{""},
			numbers: []string{"+37167812439"},
			want:    map[string]string{"": "+37167812439"},
		},
		{
			name:    "no phone numbers, only names",
			names:   []string{"Iļja", "Glebs"},
			numbers: []string{"", ""},
			want:    map[string]string{"Glebs": "", "Iļja": ""},
		},
		{
			name:    "one contact",
			names:   []string{"Denis M"},
			numbers: []string{"+37187847375"},
			want:    map[string]string{"Denis M": "+37187847375"},
		},

		{
			name:    "example from the task",
			names:   []string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			numbers: []string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			want:    map[string]string{"Alina": "+37126505719", "Antons": "+37128852154", "Deniss B": "+37127785804"},
		},

		{
			name:    "unique names with the same phone number",
			names:   []string{"Aleksejs", "Anastasija", "Vladislavs", "Gļebs"},
			numbers: []string{"+37128167023", "+37128167023", "+37128167023", "+37128167023"},
			want:    map[string]string{"Aleksejs": "+37128167023", "Anastasija": "+37128167023", "Gļebs": "+37128167023", "Vladislavs": "+37128167023"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createPhoneBook(tt.names, tt.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, tt.want)
			}
		})
	}
}
