// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/on_demand/v2/on_demand.proto

package envoy_config_filter_http_on_demand_v2

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

// Validate checks the field values on OnDemand with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OnDemand) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// OnDemandValidationError is the validation error returned by
// OnDemand.Validate if the designated constraints aren't met.
type OnDemandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OnDemandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OnDemandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OnDemandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OnDemandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OnDemandValidationError) ErrorName() string { return "OnDemandValidationError" }

// Error satisfies the builtin error interface
func (e OnDemandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOnDemand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OnDemandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OnDemandValidationError{}
