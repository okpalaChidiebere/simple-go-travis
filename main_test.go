package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FavoriteFood(t *testing.T) {
	testCases := []struct {
		name string
		favoriteFood string

		expectedError error
	}{
		{
			name: "Empty FavoriteFood - should return an error",
			expectedError: errors.New("favoritee food is undefined"),
		},
		{
			name: "Non Empty FavoriteFood - should not return an error",
			favoriteFood: "Ogbono Soup",
		},
	}

	for i := range testCases{
		c := testCases[i]
		t.Run(c.name, func (t *testing.T)  {
			gotError := favoriteFood(c.favoriteFood)

			assert.Equal(t, gotError, c.expectedError)
		})
	}

}