package assert

import (
	"reflect"
	"testing"
)

func Equal[T any](tb testing.TB, expected, actual T) {
	tb.Helper()
	if !reflect.DeepEqual(expected, actual) {
		tb.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func NotEqual[T any](tb testing.TB, expected, actual T) {
	tb.Helper()
	if reflect.DeepEqual(expected, actual) {
		tb.Errorf("Expected not %v, but got %v", expected, actual)
	}
}

func True(tb testing.TB, value bool) {
	tb.Helper()
	if !value {
		tb.Errorf("Expected true, but got false")
	}
}
func False(tb testing.TB, value bool) {
	tb.Helper()
	if value {
		tb.Errorf("Expected false, but got true")
	}
}

func NoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Errorf("Expected no error, but got %v", err)
	}
}

func Error(tb testing.TB, err error) {
	tb.Helper()
	if err == nil {
		tb.Errorf("Expected error, but got nil")
	}
}

func NotNil(tb testing.TB, value any) {
	tb.Helper()
	if value == nil {
		tb.Errorf("Expected not nil, but got nil")
	}
}

func Nil(tb testing.TB, value any) {
	tb.Helper()
	if value != nil {
		tb.Errorf("Expected nil, but got %v", value)
	}
}

func Slice[T any](tb testing.TB, expected, found []T) {
	tb.Helper()
	if expected == nil {
		Nil(tb, found)
		return
	}

	NotNil(tb, found)
	for i, value := range expected {
		Equal(tb, value, found[i])
	}

	if len(expected) != len(found) {
		for i, value := range expected {
			if i >= len(found) {
				tb.Errorf("found more than expected, got %v", value)
			}
		}
		for i, value := range found {
			if i >= len(expected) {
				tb.Errorf("expected more items than found, expected %v", value)
			}
		}
	}
}
