package logging

import (
	"testing"

	"github.com/ljpx/test"
)

func TestDummyLoggerAssertLoggedFailure(t *testing.T) {
	// Arrange.
	logger := NewDummyLogger()
	recorder := test.NewRecorder()

	// Act.
	logger.Printf("Hello, World!")
	logger.AssertLogged(recorder, "Hello, Worlds!")

	// Assert.
	test.That(t, recorder.DidFail).IsTrue()
	test.That(t, recorder.HelperCallCount).IsEqualTo(1)
}

func TestDummyLoggerAssertLoggedSuccess(t *testing.T) {
	// Arrange.
	logger := NewDummyLogger()
	recorder := test.NewRecorder()

	// Act.
	logger.Printf("Hello, World!")
	logger.AssertLogged(recorder, "Hello, World!")

	// Assert.
	test.That(t, recorder.DidFail).IsFalse()
	test.That(t, recorder.HelperCallCount).IsEqualTo(1)
}
