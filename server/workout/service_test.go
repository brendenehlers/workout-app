package workout

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWorkout(t *testing.T) {
	service := WorkoutService{}

	q := "test"
	w, err := service.CreateWorkout(context.Background(), q)

	assert.NoError(t, err)
	assert.Equal(t, "Killer Workout 3000", w.Name)
}
