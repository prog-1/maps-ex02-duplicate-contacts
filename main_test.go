package main

import (
	"reflect"
	"testing"
)

func TestCreatePhoneBook(t *testing.T) {
	for _, tc := range []struct {
		names   []string
		numbers []string
		want    map[string]string
	}{
		{[]string{"A1", "A2", "A3"}, []string{"+371222", "+371111", "+371333"}, map[string]string{"A1": "+371222", "A2": "+371111", "A3": "+371333"}},
		{[]string{"+371234"}, []string{"A2"}, map[string]string{"+371234": "A2"}},
		{[]string{"A1", "A1"}, []string{"2266", "+3389"}, map[string]string{"A1": "+3389"}},
		{[]string{",!"}, []string{":0="}, map[string]string{",!": ":0="}},
		{[]string{"A1", "A2", "A9", "A1", "A9"}, []string{"11", "22", "99", "88", "99"}, map[string]string{"A1": "88", "A2": "22", "A9": "99"}},
	} {
		got := createPhoneBook(tc.names, tc.numbers)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("createPhoneBook(%v, %v) = %v, want = %v", tc.names, tc.numbers, got, tc.want)
		}
	}
}
