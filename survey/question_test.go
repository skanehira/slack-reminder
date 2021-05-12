package survey

import "testing"

func TestValidateHour(t *testing.T) {
	t.Run("valid hour", func(t *testing.T) {
		tests := []struct {
			in string
		}{
			{
				in: "10:00",
			},
			{
				in: "01:00",
			},
		}

		for _, tt := range tests {
			if err := ValidateHour(tt.in); err != nil {
				t.Fatalf("unexpected result. want=nil, got=%s", err)
			}
		}
	})

	t.Run("invalid hour", func(t *testing.T) {
		tests := []struct {
			in   string
			want string
		}{
			{
				in:   "1",
				want: "invalid hour: 1",
			},
			{
				in:   "01:",
				want: "invalid hour: 01:",
			},
			{
				in:   ":1",
				want: "invalid hour: :1",
			},
			{
				in:   ":10",
				want: "invalid hour: :10",
			},
			{
				in:   "10:999",
				want: "invalid hour: 10:999",
			},
		}

		for _, tt := range tests {
			got := ValidateHour(tt.in)
			if got == nil || tt.want != got.Error() {
				t.Fatalf("unexpected result. want=%s, got=%s", tt.want, got)
			}
		}
	})
}

func TestValidateDate(t *testing.T) {
	t.Run("valid date", func(t *testing.T) {
		tests := []struct {
			in string
		}{
			{
				in: "2020-01-01",
			},
			{
				in: "1992-10-29",
			},
		}

		for _, tt := range tests {
			if err := ValidateDate(tt.in); err != nil {
				t.Fatalf("unexpected result. want=nil, got=%s", err)
			}
		}
	})

	t.Run("invalid date", func(t *testing.T) {
		tests := []struct {
			in   string
			want string
		}{
			{
				in:   "202-01-01",
				want: "invalid date: 202-01-01",
			},
			{
				in:   "2020-1-29",
				want: "invalid date: 2020-1-29",
			},
			{
				in:   "2020-01-1",
				want: "invalid date: 2020-01-1",
			},
			{
				in:   "202-01-011",
				want: "invalid date: 202-01-011",
			},
		}

		for _, tt := range tests {
			got := ValidateDate(tt.in)
			if got == nil || tt.want != got.Error() {
				t.Fatalf("unexpected result. want=%s, got=%s", tt.want, got)
			}
		}
	})
}
