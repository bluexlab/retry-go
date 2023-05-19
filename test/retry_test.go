package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bluexlab/retry-go"
	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	needRetry := errors.New("ALSKDJFALKDSJF")
	realError := errors.New("DON'T RETRY")
	shouldRetry := func(e error) bool {
		return e == needRetry
	}

	r := retry.New(shouldRetry, 10, 10, 1000)

	count := 0
	err := r.Do(func() error {
		count = count + 1
		return needRetry
	})
	assert.Equal(t, 10, count)
	assert.IsType(t, &retry.ErrMaxAttemptExceeded{}, err)
	assert.Equal(t, needRetry, err.(*retry.ErrMaxAttemptExceeded).Err)

	count = 0
	err = r.Do(func() error {
		count = count + 1
		return realError
	})
	assert.Equal(t, 1, count)
	assert.Equal(t, realError, err)

	count = 0
	err = r.Do(func() error {
		count = count + 1
		return nil
	})
	assert.Equal(t, 1, count)
	assert.NoError(t, err)

	count = 0
	okAfter2 := func(input string, to string) (string, error) {
		count = count + 1
		if count == 2 {
			return fmt.Sprintf("%s %s", input, to), nil
		}
		return "", needRetry
	}
	result, err := retry.Retry2Func2(r, okAfter2, "hello", "world")
	assert.NoError(t, err)
	assert.Equal(t, 2, count)
	assert.Equal(t, "hello world", result)
}
