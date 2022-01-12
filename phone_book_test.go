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
			name:     "empty",
			names:    []string{},
			phonenum: []string{},
			want:     map[string]string{},
		},
		{
			name:     "example",
			names:    []string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			phonenum: []string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			want:     map[string]string{"Alina": "+37126505719", "Antons": "+37128852154", "Deniss B": "+37127785804"},
		},