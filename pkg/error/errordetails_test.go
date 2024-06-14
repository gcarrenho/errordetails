package errordetails

import (
	"errors"
	"testing"
)

func TestNewErrorDetails(t *testing.T) {
	baseErr := errors.New("an error occurred")
	customErr := NewErrorDetails(baseErr)

	if customErr.err != baseErr {
		t.Errorf("expected %v, got %v", baseErr, customErr.err)
	}

	/*if customErr.Message == "" {
		t.Errorf("expected title to be initialized, got empty")
	}*/
}

func TestErrorDetails_Str(t *testing.T) {
	baseErr := errors.New("an error occurred")
	errorDetails := NewErrorDetails(baseErr)

	// Add fields using Str method
	errorDetails.Str("OrderItemID", "123").
		Str("PresetID", "preset123").
		Str("SettingsJSON", `{"key":"value"}`)

	// Check if the fields were added correctly
	expectedFields := []field{
		{key: "OrderItemID", val: "123"},
		{key: "PresetID", val: "preset123"},
		{key: "SettingsJSON", val: `{"key":"value"}`},
	}

	if len(errorDetails.fields) != len(expectedFields) {
		t.Fatalf("expected %d fields, got %d", len(expectedFields), len(errorDetails.fields))
	}

	for i, field := range errorDetails.fields {
		if field.key != expectedFields[i].key || field.val != expectedFields[i].val {
			t.Errorf("expected field %d to be %v, got %v", i, expectedFields[i], field)
		}
	}
}

func TestErrorDetails_Int(t *testing.T) {
	baseErr := errors.New("an error occurred")
	errorDetails := NewErrorDetails(baseErr)

	// Add integer fields using Int method
	errorDetails.Int("OrderItemID", 123).
		Int("ErrorCode", 404)

	// Check if the fields were added correctly
	expectedFields := []field{
		{key: "OrderItemID", val: "123"},
		{key: "ErrorCode", val: "404"},
	}

	if len(errorDetails.fields) != len(expectedFields) {
		t.Fatalf("expected %d fields, got %d", len(expectedFields), len(errorDetails.fields))
	}

	for i, field := range errorDetails.fields {
		if field.key != expectedFields[i].key || field.val != expectedFields[i].val {
			t.Errorf("expected field %d to be %v, got %v", i, expectedFields[i], field)
		}
	}
}

func TestErrorDetails_Error(t *testing.T) {
	baseErr := errors.New("an error occurred")
	customErr := NewErrorDetails(baseErr).
		Str("OrderItemID", "123").
		Str("PresetID", "preset123").
		Str("SettingsJSON", `{"key":"value"}`).Msg("message error example")

	expectedMessage := "message error example | OrderItemID: 123 | PresetID: preset123 | SettingsJSON: {\"key\":\"value\"} --> an error occurred"

	if customErr.Error() != expectedMessage {
		t.Errorf("expected %s, got %s", expectedMessage, customErr.Error())
	}
}
