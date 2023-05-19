package retry

import (
	"fmt"
	"math/rand"
	"time"
)

// Retry is a helper to retry a function under the specific conditions.
type Retry struct {
	shouldRetry func(error) bool
	maxAttempt  int // max attemp
	initDelay   int // ms
	maxDelay    int // ms
}

// ErrMaxAttemptExceeded wraps the original error when the max retry attempt exceeded.
type ErrMaxAttemptExceeded struct {
	Err error
}

func (e *ErrMaxAttemptExceeded) Error() string {
	return fmt.Sprintf("exceed max retry attempts. Original error: %v", e.Err.Error())
}

func (e *ErrMaxAttemptExceeded) Unwrap() error {
	return e.Err
}

// New creates a "Retry"
// shouldRetry is a function to decide if a function should retry.
// maxAttemp specifies the max attempts.
// delay is the delay between retries. The unit is ms.
func New(shouldRetry func(error) bool, maxAttempt int, initDelay int, maxDelay int) Retry {
	return Retry{
		shouldRetry: shouldRetry,
		maxAttempt:  maxAttempt,
		initDelay:   initDelay,
		maxDelay:    maxDelay,
	}
}

// Do calls the input function and check the result.
// ErrMaxAttemptExceeded returns when maxAttamp exceeded.
func (r Retry) Do(f func() error) error {
	if r.maxAttempt <= 0 {
		panic("maxAttemp must be greater than 0")
	}
	maxAttempt := r.maxAttempt
	delay := r.initDelay
	var lastErr error
	for i := 0; i < maxAttempt; i++ {
		lastErr = f()
		if lastErr == nil {
			return nil
		}
		if r.shouldRetry(lastErr) {
			realDelay := int(float32(delay) * rand.Float32())
			time.Sleep(time.Duration(realDelay) * time.Millisecond)
			delay = delay * 2
			if delay > r.maxDelay {
				delay = r.maxDelay
			}
			continue
		}
		return lastErr
	}

	return &ErrMaxAttemptExceeded{
		Err: lastErr,
	}
}

func Retry2[R any](r Retry, f func() (R, error)) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f()
		return e
	})
	return result, err
}

func Retry2Func1[R, P1 any](r Retry, f func(P1) (R, error), p1 P1) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1)
		return e
	})
	return result, err
}

func Retry2Func2[R, P1, P2 any](r Retry, f func(P1, P2) (R, error), p1 P1, p2 P2) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2)
		return e
	})
	return result, err
}

func Retry2Func3[R, P1, P2, P3 any](r Retry, f func(P1, P2, P3) (R, error), p1 P1, p2 P2, p3 P3) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2, p3)
		return e
	})
	return result, err
}

func Retry2Func4[R, P1, P2, P3, P4 any](r Retry, f func(P1, P2, P3, P4) (R, error), p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2, p3, p4)
		return e
	})
	return result, err
}

func Retry2Func5[R, P1, P2, P3, P4, P5 any](r Retry, f func(P1, P2, P3, P4, P5) (R, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2, p3, p4, p5)
		return e
	})
	return result, err
}

func Retry2Func6[R, P1, P2, P3, P4, P5, P6 any](r Retry, f func(P1, P2, P3, P4, P5, P6) (R, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2, p3, p4, p5, p6)
		return e
	})
	return result, err
}

func Retry2Func7[R, P1, P2, P3, P4, P5, P6, P7 any](r Retry, f func(P1, P2, P3, P4, P5, P6, P7) (R, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2, p3, p4, p5, p6, p7)
		return e
	})
	return result, err
}

func Retry2Func8[R, P1, P2, P3, P4, P5, P6, P7, P8 any](r Retry, f func(P1, P2, P3, P4, P5, P6, P7, P8) (R, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
	var result R
	err := r.Do(func() error {
		var e error
		result, e = f(p1, p2, p3, p4, p5, p6, p7, p8)
		return e
	})
	return result, err
}

func Retry3[R1, R2 any](r Retry, f func() (R1, R2, error)) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f()
		return e
	})
	return result1, result2, err
}

func Retry3Func1[R1, R2, P1 any](r Retry, f func(P1) (R1, R2, error), p1 P1) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1)
		return e
	})
	return result1, result2, err
}

func Retry3Func2[R1, R2, P1, P2 any](r Retry, f func(P1, P2) (R1, R2, error), p1 P1, p2 P2) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2)
		return e
	})
	return result1, result2, err
}

func Retry3Func3[R1, R2, P1, P2, P3 any](r Retry, f func(P1, P2, P3) (R1, R2, error), p1 P1, p2 P2, p3 P3) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2, p3)
		return e
	})
	return result1, result2, err
}

func Retry3Func4[R1, R2, P1, P2, P3, P4 any](r Retry, f func(P1, P2, P3, P4) (R1, R2, error), p1 P1, p2 P2, p3 P3, p4 P4) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2, p3, p4)
		return e
	})
	return result1, result2, err
}

func Retry3Func5[R1, R2, P1, P2, P3, P4, P5 any](r Retry, f func(P1, P2, P3, P4, P5) (R1, R2, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2, p3, p4, p5)
		return e
	})
	return result1, result2, err
}

func Retry3Func6[R1, R2, P1, P2, P3, P4, P5, P6 any](r Retry, f func(P1, P2, P3, P4, P5, P6) (R1, R2, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2, p3, p4, p5, p6)
		return e
	})
	return result1, result2, err
}

func Retry3Func7[R1, R2, P1, P2, P3, P4, P5, P6, P7 any](r Retry, f func(P1, P2, P3, P4, P5, P6, P7) (R1, R2, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2, p3, p4, p5, p6, p7)
		return e
	})
	return result1, result2, err
}

func Retry3Func8[R1, R2, P1, P2, P3, P4, P5, P6, P7, P8 any](r Retry, f func(P1, P2, P3, P4, P5, P6, P7, P8) (R1, R2, error), p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R1, R2, error) {
	var result1 R1
	var result2 R2
	err := r.Do(func() error {
		var e error
		result1, result2, e = f(p1, p2, p3, p4, p5, p6, p7, p8)
		return e
	})
	return result1, result2, err
}
