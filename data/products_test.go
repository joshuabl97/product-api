package data

import "testing"

func TestCheckValidation(t *testing.T) {
	// pass in an empty product to check which validations will make the request fail
	p := &Product{
		Name:  "test",
		Price: 420.69,
		SKU:   "sadfioh-oque-aoefib",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
