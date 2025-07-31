package factories

import "github.com/brianvoe/gofakeit/v7"

type PersonFactory struct {
}

// Definition Define the model's default state.
func (f *PersonFactory) Definition() map[string]any {
	return map[string]any{
		"Name":    gofakeit.Name(),
		"Email":   gofakeit.Email(),
		"Phone":   gofakeit.Phone(),
		"Age":     gofakeit.IntRange(10, 100),
		"Address": gofakeit.Address().Address,
	}
}
