package main

import "testing"

func TestPhonebook(t *testing.T) {
	tests := []struct {
		name  string
		numbers,names []string
		want 	map[string]string
	}{
		{
			names: []string{"Alina","Deniss B","Antons","Alina","Antons "},
			numbers: []string{"+37126017505","+37127785804","+37123622588","+37126505719","+37128852154"},
			want: map[string]string{"Alina": "+37126017505","Antons": "+37123622588", "Deniss.B": "+37127785804"}
			},
			{
				names:   []string{"Alina", "Alina", "Alina", "Alina", "Alina"},
			numbers: []string{"+37126017505", "+37126017505", "+37126017505", "+37126017505", "+37126017505"},
			want:    map[string]string{"Alina": "+37126017505"},
			},
			{
				names:   nil,
			numbers: nil,
			want:    map[string]string{},
		},
		for _, a :=range resrts {
			t.Run(a.names, func(t *testing.T)) {
				if got := createPhoneBook(a.number, a.names); !equal(got, a.want){
					t.Errorf("createPhoneBook() = %v, want %v", got, a.want)
				}
			}

		}
		{
			}

		}
	}
	
}