package main

import (
	"fmt"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	for _, tc := range []struct {
		name    string
		names   []string
		numbers []string
		want    map[string]string
	}{
		{"simple", []string{"Anna"}, []string{"+37125896457"}, map[string]string{"Anna": "+37125896457"}},
		{"3 numbers", []string{"Andry", "Kate", "Emili"}, []string{"+371262992451", "+37120555697", "+37152000690"}, map[string]string{"Andry": "+371262992451", "Kate": "+37120555697", "Emili": "+37152000690"}},
		{"empty", []string{}, []string{}, map[string]string{0}},
		{"2 same numbers", []string{"Anna", "Sam"}, []string{"+371262992451", "+371262992451"}, map[string]string{"Something wrong,you entered same numbers or names"}},
		{"2 same names", []string{"Anna", "Anna"}, []string{"+371262992451", "+37152000690"}, map[string]string{"Something wrong,you entered same numbers or names"}},
	} {
		if got := createPhoneBook(tc.names, tc.numbers); got != tc.want {
			fmt.Printf("FAIL %s: got %v, want %v", tc.names, tc.numbers, got, tc.want)

		}
	}
}
