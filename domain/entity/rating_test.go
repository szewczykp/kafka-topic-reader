package entity

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRating(t *testing.T) {
	expectedId := "someId"
	expectedRecipeId := "someRecipeId"
	expectedValue := int8(5)

	jsonString := fmt.Sprintf(
		"{\"Id\": \"%s\", \"RecipeId\": \"%s\", \"Value\": %d}",
		expectedId,
		expectedRecipeId,
		expectedValue,
	)

	rating := &Rating{}

	err := json.Unmarshal([]byte(jsonString), rating)

	assert.Nil(t, err)
	assert.Equal(t, expectedId, rating.Id)
	assert.Equal(t, expectedRecipeId, rating.RecipeId)
	assert.Equal(t, expectedValue, rating.Value)
}

func TestRating_WithInvalidType(t *testing.T) {
	expectedId := "someId"
	expectedRecipeId := 1918
	expectedValue := int8(5)

	jsonString := fmt.Sprintf(
		"{\"Id\": \"%s\", \"RecipeId\": %d, \"Value\": %d}",
		expectedId,
		expectedRecipeId,
		expectedValue,
	)

	rating := &Rating{}

	err := json.Unmarshal([]byte(jsonString), rating)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "RecipeId")
}
