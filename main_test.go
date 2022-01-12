package main

import (
	"reflect"
	"testing"
)

func TestcreatePhoneBook(t *testing.T) {
	for _, tc := range []struct {
		name  string
		names string
		numbers string
		want  map[string]string
		}{
			{"empty-nil", nil, nil, map[string]string},
			{"one name and contact", "Alesja", "+37122331231", map[Alesja:+37122331231]},
			{"example", "Alina", "Deniss B", "Antons", "Alina", "Antons",  "+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154", map[Alina: +37126505719], map[Antons: +37128852154], map[Deniss B: +37127785804]},
		} {
			if got := createPhoneBook(tc.names, tc.numbers); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		}
	}
	// tests are failed