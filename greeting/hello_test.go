package greeting

import "testing"

func TestSayGreeting(t *testing.T) {

	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Make sure to input name argument",
		},
		{
			result: "Hello Nabil!",
			items:  []string{"Nabil"},
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Wanted %s [%v], but got %s", st.result, st.items, s)
		}
	}

}
