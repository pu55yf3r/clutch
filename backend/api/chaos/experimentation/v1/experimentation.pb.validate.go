// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: chaos/experimentation/v1/experimentation.proto

package experimentationv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _experimentation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Experiment with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Experiment) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExperimentValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExperimentValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExperimentValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ExperimentValidationError is the validation error returned by
// Experiment.Validate if the designated constraints aren't met.
type ExperimentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExperimentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExperimentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExperimentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExperimentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExperimentValidationError) ErrorName() string { return "ExperimentValidationError" }

// Error satisfies the builtin error interface
func (e ExperimentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExperiment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExperimentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExperimentValidationError{}

// Validate checks the field values on CreateExperimentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateExperimentRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateExperimentRequestValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStartTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateExperimentRequestValidationError{
				field:  "StartTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEndTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateExperimentRequestValidationError{
				field:  "EndTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateExperimentRequestValidationError is the validation error returned by
// CreateExperimentRequest.Validate if the designated constraints aren't met.
type CreateExperimentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateExperimentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateExperimentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateExperimentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateExperimentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateExperimentRequestValidationError) ErrorName() string {
	return "CreateExperimentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateExperimentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateExperimentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateExperimentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateExperimentRequestValidationError{}

// Validate checks the field values on CreateExperimentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateExperimentResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetExperiment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateExperimentResponseValidationError{
				field:  "Experiment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateExperimentResponseValidationError is the validation error returned by
// CreateExperimentResponse.Validate if the designated constraints aren't met.
type CreateExperimentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateExperimentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateExperimentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateExperimentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateExperimentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateExperimentResponseValidationError) ErrorName() string {
	return "CreateExperimentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateExperimentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateExperimentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateExperimentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateExperimentResponseValidationError{}

// Validate checks the field values on GetExperimentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExperimentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetExperimentsRequestValidationError is the validation error returned by
// GetExperimentsRequest.Validate if the designated constraints aren't met.
type GetExperimentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExperimentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExperimentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExperimentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExperimentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExperimentsRequestValidationError) ErrorName() string {
	return "GetExperimentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExperimentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExperimentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExperimentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExperimentsRequestValidationError{}

// Validate checks the field values on GetExperimentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExperimentsResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetExperiments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetExperimentsResponseValidationError{
					field:  fmt.Sprintf("Experiments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetExperimentsResponseValidationError is the validation error returned by
// GetExperimentsResponse.Validate if the designated constraints aren't met.
type GetExperimentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExperimentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExperimentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExperimentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExperimentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExperimentsResponseValidationError) ErrorName() string {
	return "GetExperimentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExperimentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExperimentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExperimentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExperimentsResponseValidationError{}

// Validate checks the field values on GetExperimentRunDetailsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExperimentRunDetailsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetExperimentRunDetailsRequestValidationError is the validation error
// returned by GetExperimentRunDetailsRequest.Validate if the designated
// constraints aren't met.
type GetExperimentRunDetailsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExperimentRunDetailsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExperimentRunDetailsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExperimentRunDetailsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExperimentRunDetailsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExperimentRunDetailsRequestValidationError) ErrorName() string {
	return "GetExperimentRunDetailsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExperimentRunDetailsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExperimentRunDetailsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExperimentRunDetailsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExperimentRunDetailsRequestValidationError{}

// Validate checks the field values on GetExperimentRunDetailsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetExperimentRunDetailsResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRunDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetExperimentRunDetailsResponseValidationError{
				field:  "RunDetails",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetExperimentRunDetailsResponseValidationError is the validation error
// returned by GetExperimentRunDetailsResponse.Validate if the designated
// constraints aren't met.
type GetExperimentRunDetailsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExperimentRunDetailsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExperimentRunDetailsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExperimentRunDetailsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExperimentRunDetailsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExperimentRunDetailsResponseValidationError) ErrorName() string {
	return "GetExperimentRunDetailsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExperimentRunDetailsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExperimentRunDetailsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExperimentRunDetailsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExperimentRunDetailsResponseValidationError{}

// Validate checks the field values on StopExperimentsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StopExperimentsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StopExperimentsRequestValidationError is the validation error returned by
// StopExperimentsRequest.Validate if the designated constraints aren't met.
type StopExperimentsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopExperimentsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopExperimentsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopExperimentsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopExperimentsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopExperimentsRequestValidationError) ErrorName() string {
	return "StopExperimentsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StopExperimentsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopExperimentsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopExperimentsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopExperimentsRequestValidationError{}

// Validate checks the field values on StopExperimentsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StopExperimentsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StopExperimentsResponseValidationError is the validation error returned by
// StopExperimentsResponse.Validate if the designated constraints aren't met.
type StopExperimentsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StopExperimentsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StopExperimentsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StopExperimentsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StopExperimentsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StopExperimentsResponseValidationError) ErrorName() string {
	return "StopExperimentsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e StopExperimentsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStopExperimentsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StopExperimentsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StopExperimentsResponseValidationError{}
