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
			name:    "number without name",
			names:   []string{""},
			numbers: []string{"+37125152348"},
			want:    map[string]string{"": "+37125152348"},
		},
		{
			name:    "name without number",
			names:   []string{"name"},
			numbers: []string{""},
			want:    map[string]string{"name": ""},
		},
		{
			name:    "name and number",
			names:   []string{"name"},
			numbers: []string{"+37125152348"},
			want:    map[string]string{"name": "+37125152348"},
		},
		{
			name:    "two contacts",
			names:   []string{"name", "secondname"},
			numbers: []string{"+37125152348", "+37125152349"},
			want:    map[string]string{"name": "+37125152348", "secondname": "+37125152349"},
		},
		{
			name:    "two same numbers and names",
			names:   []string{"name", "name"},
			numbers: []string{"+37125152348", "+37125152348"},
			want:    map[string]string{"name": "+37125152348"},
		},
		{
			name:    "names with same number",
			names:   []string{"name1", "name2", "name3"},
			numbers: []string{"+37125152348", "+37125152348", "+37125152348"},
			want:    map[string]string{"name1": "+37125152348", "name2": "+37125152348", "name3": "+37125152348"},
		},
		{
			name:    "numbers with same names",
			names:   []string{"name", "name", "name"},
			numbers: []string{"+37125152348", "+37125152349", "+37125152347", "+37125152346"},
			want:    map[string]string{"name": "+37125152348,+37125152349,+37125152347,+37125152346"},
		},
	}
	for _, a := range tests {
		t.Run(a.name, func(t *testing.T) {
			if got := createPhoneBook(a.names, a.numbers); !reflect.DeepEqual(got, a.want) {
				t.Errorf("createPhoneBook() = %v, want %v", got, a.want)
			}
		})
	}
}
