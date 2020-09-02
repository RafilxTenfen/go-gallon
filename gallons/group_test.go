package gallons_test

import (
	"testing"

	"github.com/RafilxTenfen/go-gallon/gallons"
)

func TestUseBottles(t *testing.T) {
	tests := []struct {
		Group         gallons.Group
		ExpectedRest  int
		ExpectedError error
	}{
		{
			Group:         gallons.CreateBottleGroup(3500, 1500, 2000),
			ExpectedRest:  0,
			ExpectedError: nil,
		},
		{
			Group:         gallons.CreateBottleGroup(5000, 1500, 2000, 1500, 1000),
			ExpectedRest:  0,
			ExpectedError: nil,
		},
		{
			Group:         gallons.CreateBottleGroup(7000, 1000, 3000, 4500, 1500, 3500),
			ExpectedRest:  0,
			ExpectedError: nil,
		},
		{
			Group:         gallons.CreateBottleGroup(5000, 1000, 3000, 4500, 1500),
			ExpectedRest:  500,
			ExpectedError: nil,
		},
	}

	for i, test := range tests {
		receivedRest, bottles, receivedErr := test.Group.UseBottles()
		if test.ExpectedRest != receivedRest {
			t.Errorf("Error on test %d \nExpected: '%d' \nReceived: '%d' \nBottles: '%+v'", i, test.ExpectedRest, receivedRest, bottles)
		}

		if test.ExpectedError != receivedErr {
			t.Errorf("Error on test %d \nExpected: '%+v' \nReceived: '%+v' \nBottles: '%+v'", i, test.ExpectedError, receivedErr, bottles)
		}
	}
}
