package survey

import "testing"

func newOnetime() Onetime {
	o := Onetime{
		Date: "2021-05-12",
		Hour: "15:00",
	}
	return o
}

func TestOnetime(t *testing.T) {
	want := "at 15:00 on 2021-05-12"
	got := newOnetime().String()

	if want != got {
		t.Fatalf("unexpected result. want=%s, got=%s", want, got)
	}
}

func TestAnswer(t *testing.T) {
	ans := Answer{
		Destination: "me",
		Message:     "test",
		When:        newOnetime(),
	}

	want := `/remind me "test" at 15:00 on 2021-05-12`
	got := ans.String()
	if want != got {
		t.Fatalf("unexpected result. want=%s, got=%s", want, got)
	}
}
