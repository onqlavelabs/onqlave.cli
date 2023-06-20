package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAppErrorRender(t *testing.T) {
	t.Run("Minimal", func(t *testing.T) {
		aerr := NewAppError("here", "message", nil, "", http.StatusTeapot)
		assert.EqualError(t, aerr, "here: message")
	})

	t.Run("Detailed", func(t *testing.T) {
		aerr := NewAppError("here", "message", nil, "details", http.StatusTeapot)
		assert.EqualError(t, aerr, "here: message, details")
	})

	t.Run("Wrapped", func(t *testing.T) {
		rootError := fmt.Errorf("root error")
		aerr := NewAppError("here", "message", nil, "", http.StatusTeapot).Wrap(rootError)
		assert.EqualError(t, aerr, "here: message, root error")
	})

	t.Run("WrappedMultiple", func(t *testing.T) {
		aerr := NewAppError("here", "message", nil, "", http.StatusTeapot).Wrap(fmt.Errorf("my error (%w)", fmt.Errorf("inner error")))
		assert.EqualError(t, aerr, "here: message, my error (inner error)")
	})

	t.Run("DetailedWrappedMultiple", func(t *testing.T) {
		aerr := NewAppError("here", "message", nil, "details", http.StatusTeapot).Wrap(fmt.Errorf("my error (%w)", fmt.Errorf("inner error")))
		assert.EqualError(t, aerr, "here: message, details, my error (inner error)")
	})
}
