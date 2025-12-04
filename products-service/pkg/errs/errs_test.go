package errs

import (
	"errors"
	"testing"
)

// TestErrorsImplementHTTPError tests that all custom errors implement HTTPError interface.
func TestErrorsImplementHTTPError(t *testing.T) {
	types := []HTTPError{
		// Custom errors
		Empty{},
		FieldsValidation{},

		// Client errors
		BadRequest{},
		Unauthorized{},
		PaymentRequired{},
		Forbidden{},
		NotFound{},
		MethodNotAllowed{},
		NotAcceptable{},
		RequestTimeout{},
		Conflict{},
		Gone{},
		PayloadTooLarge{},
		UnsupportedMediaType{},
		TooManyRequests{},

		// Server errors
		Internal{},
		NotImplemented{},
		BadGateway{},
		ServiceUnavailable{},
		GatewayTimeout{},
		HTTPVersionNotSupported{},
	}

	for _, e := range types {
		var httpErr HTTPError
		if !errors.As(e, &httpErr) {
			t.Errorf("type %T does not implement HTTPError", e)
		}
	}
}

// TestInternalSrv_Error tests the Error method of the Internal struct.
func TestInternalSrv_Error(t *testing.T) {
	type fields struct {
		Cause string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "correct error",
			fields: fields{Cause: "test"},
			want:   "internal_server_error" + divider + "test",
		},

		{
			name:   "empty error",
			fields: fields{Cause: ""},
			want:   "internal_server_error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Internal{
				Cause: tt.fields.Cause,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFieldsValidation_Error tests the Error method of the FieldsValidation struct.
func TestFieldsValidation_Error(t *testing.T) {
	type fields struct {
		FieldErrors []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "[success] correct error",
			fields: fields{
				FieldErrors: []string{
					"cyrillic_min::min_length_is::70",
					"polish_min::min_length_is::111",
					"german_min::min_length_is::120",
				},
			},
			want: "validation_error" + divider +
				"cyrillic_min::min_length_is::70,polish_min::min_length_is::111,german_min::min_length_is::120",
		},

		{
			name:   "[success] empty error with initial slice",
			fields: fields{FieldErrors: []string{}},
			want:   "validation_error",
		},

		{
			name:   "[success] empty error with nil slice",
			fields: fields{FieldErrors: nil},
			want:   "validation_error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := FieldsValidation{
				Errors: tt.fields.FieldErrors,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("FieldsValidation.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
