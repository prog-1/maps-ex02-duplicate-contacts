package main

func TestCreatePhoneBook(t *testing.T){
	for _, tc := []struct {
		size uint
		names []string
		numbers []string
		want map[string]string
	}{
		{
			0,
			[]string{},
			[]string{},
			map[string]string{}
		},
		{
			5,
			[]string{"Alina", "Deniss B", "Antons", "Alina", "Antons"},
			[]string{"+37126017505", "+37127785804", "+37123622588", "+37126505719", "+37128852154"},
			map[string]string{"Alina": "+37126505719", "Antons": "+37128852154", "Deniss B": "+37127785804"},
		},
		{
			4,
			[]string{"Alina", "Alina", "Alina", "Alina"},
			[]string{"+37126017505", "+371705", "+371234234237505", "+37126017505"},
			map[string]string{"Alina": "+37126017505"},
		},
		{
			3,
			[]string{"Tihons", "Arina", "Egors"},
			[]string{"+37134988264", "+37192744934", "+3712281337"},
			map[string]string{"Tihons": "+37134988264", "Arina": "+37192744934", "Egors": "+3712281337"},
		},
	}{
		got := CreatePhoneBook(tc.names, tc.numbers)
		if !equal(got, tc.want) {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
func equal(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}