// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/protocol.proto

package envoy_config_core_v3

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

// Validate checks the field values on TcpProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TcpProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// TcpProtocolOptionsValidationError is the validation error returned by
// TcpProtocolOptions.Validate if the designated constraints aren't met.
type TcpProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpProtocolOptionsValidationError) ErrorName() string {
	return "TcpProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e TcpProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpProtocolOptionsValidationError{}

// Validate checks the field values on UpstreamHttpProtocolOptions with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpstreamHttpProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AutoSni

	// no validation rules for AutoSanValidation

	return nil
}

// UpstreamHttpProtocolOptionsValidationError is the validation error returned
// by UpstreamHttpProtocolOptions.Validate if the designated constraints
// aren't met.
type UpstreamHttpProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamHttpProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamHttpProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamHttpProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamHttpProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamHttpProtocolOptionsValidationError) ErrorName() string {
	return "UpstreamHttpProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamHttpProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamHttpProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamHttpProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamHttpProtocolOptionsValidationError{}

// Validate checks the field values on HttpProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HttpProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxConnectionDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "MaxConnectionDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetMaxHeadersCount(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return HttpProtocolOptionsValidationError{
				field:  "MaxHeadersCount",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if v, ok := interface{}(m.GetMaxStreamDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "MaxStreamDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HeadersWithUnderscoresAction

	return nil
}

// HttpProtocolOptionsValidationError is the validation error returned by
// HttpProtocolOptions.Validate if the designated constraints aren't met.
type HttpProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptionsValidationError) ErrorName() string {
	return "HttpProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptionsValidationError{}

// Validate checks the field values on Http1ProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Http1ProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAllowAbsoluteUrl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http1ProtocolOptionsValidationError{
				field:  "AllowAbsoluteUrl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AcceptHttp_10

	// no validation rules for DefaultHostForHttp_10

	if v, ok := interface{}(m.GetHeaderKeyFormat()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http1ProtocolOptionsValidationError{
				field:  "HeaderKeyFormat",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EnableTrailers

	// no validation rules for AllowChunkedLength

	if v, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http1ProtocolOptionsValidationError{
				field:  "OverrideStreamErrorOnInvalidHttpMessage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Http1ProtocolOptionsValidationError is the validation error returned by
// Http1ProtocolOptions.Validate if the designated constraints aren't met.
type Http1ProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Http1ProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Http1ProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Http1ProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Http1ProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Http1ProtocolOptionsValidationError) ErrorName() string {
	return "Http1ProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e Http1ProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttp1ProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Http1ProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Http1ProtocolOptionsValidationError{}

// Validate checks the field values on KeepaliveSettings with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *KeepaliveSettings) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetInterval() == nil {
		return KeepaliveSettingsValidationError{
			field:  "Interval",
			reason: "value is required",
		}
	}

	if d := m.GetInterval(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return KeepaliveSettingsValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

		if dur < gte {
			return KeepaliveSettingsValidationError{
				field:  "Interval",
				reason: "value must be greater than or equal to 1ms",
			}
		}

	}

	if m.GetTimeout() == nil {
		return KeepaliveSettingsValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
	}

	if d := m.GetTimeout(); d != nil {
		dur, err := ptypes.Duration(d)
		if err != nil {
			return KeepaliveSettingsValidationError{
				field:  "Timeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

		if dur < gte {
			return KeepaliveSettingsValidationError{
				field:  "Timeout",
				reason: "value must be greater than or equal to 1ms",
			}
		}

	}

	if v, ok := interface{}(m.GetIntervalJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeepaliveSettingsValidationError{
				field:  "IntervalJitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// KeepaliveSettingsValidationError is the validation error returned by
// KeepaliveSettings.Validate if the designated constraints aren't met.
type KeepaliveSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeepaliveSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeepaliveSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeepaliveSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeepaliveSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeepaliveSettingsValidationError) ErrorName() string {
	return "KeepaliveSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e KeepaliveSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeepaliveSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeepaliveSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeepaliveSettingsValidationError{}

// Validate checks the field values on Http2ProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Http2ProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHpackTableSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http2ProtocolOptionsValidationError{
				field:  "HpackTableSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetMaxConcurrentStreams(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 2147483647 {
			return Http2ProtocolOptionsValidationError{
				field:  "MaxConcurrentStreams",
				reason: "value must be inside range [1, 2147483647]",
			}
		}

	}

	if wrapper := m.GetInitialStreamWindowSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 65535 || val > 2147483647 {
			return Http2ProtocolOptionsValidationError{
				field:  "InitialStreamWindowSize",
				reason: "value must be inside range [65535, 2147483647]",
			}
		}

	}

	if wrapper := m.GetInitialConnectionWindowSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 65535 || val > 2147483647 {
			return Http2ProtocolOptionsValidationError{
				field:  "InitialConnectionWindowSize",
				reason: "value must be inside range [65535, 2147483647]",
			}
		}

	}

	// no validation rules for AllowConnect

	// no validation rules for AllowMetadata

	if wrapper := m.GetMaxOutboundFrames(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return Http2ProtocolOptionsValidationError{
				field:  "MaxOutboundFrames",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if wrapper := m.GetMaxOutboundControlFrames(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return Http2ProtocolOptionsValidationError{
				field:  "MaxOutboundControlFrames",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if v, ok := interface{}(m.GetMaxConsecutiveInboundFramesWithEmptyPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http2ProtocolOptionsValidationError{
				field:  "MaxConsecutiveInboundFramesWithEmptyPayload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxInboundPriorityFramesPerStream()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http2ProtocolOptionsValidationError{
				field:  "MaxInboundPriorityFramesPerStream",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetMaxInboundWindowUpdateFramesPerDataFrameSent(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return Http2ProtocolOptionsValidationError{
				field:  "MaxInboundWindowUpdateFramesPerDataFrameSent",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	// no validation rules for StreamErrorOnInvalidHttpMessaging

	if v, ok := interface{}(m.GetOverrideStreamErrorOnInvalidHttpMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http2ProtocolOptionsValidationError{
				field:  "OverrideStreamErrorOnInvalidHttpMessage",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCustomSettingsParameters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Http2ProtocolOptionsValidationError{
					field:  fmt.Sprintf("CustomSettingsParameters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetConnectionKeepalive()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http2ProtocolOptionsValidationError{
				field:  "ConnectionKeepalive",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Http2ProtocolOptionsValidationError is the validation error returned by
// Http2ProtocolOptions.Validate if the designated constraints aren't met.
type Http2ProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Http2ProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Http2ProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Http2ProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Http2ProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Http2ProtocolOptionsValidationError) ErrorName() string {
	return "Http2ProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e Http2ProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttp2ProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Http2ProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Http2ProtocolOptionsValidationError{}

// Validate checks the field values on GrpcProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GrpcProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcProtocolOptionsValidationError{
				field:  "Http2ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GrpcProtocolOptionsValidationError is the validation error returned by
// GrpcProtocolOptions.Validate if the designated constraints aren't met.
type GrpcProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcProtocolOptionsValidationError) ErrorName() string {
	return "GrpcProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcProtocolOptionsValidationError{}

// Validate checks the field values on Http1ProtocolOptions_HeaderKeyFormat
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *Http1ProtocolOptions_HeaderKeyFormat) Validate() error {
	if m == nil {
		return nil
	}

	switch m.HeaderFormat.(type) {

	case *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords_:

		if v, ok := interface{}(m.GetProperCaseWords()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Http1ProtocolOptions_HeaderKeyFormatValidationError{
					field:  "ProperCaseWords",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Http1ProtocolOptions_HeaderKeyFormat_Custom_:

		if v, ok := interface{}(m.GetCustom()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Http1ProtocolOptions_HeaderKeyFormatValidationError{
					field:  "Custom",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return Http1ProtocolOptions_HeaderKeyFormatValidationError{
			field:  "HeaderFormat",
			reason: "value is required",
		}

	}

	return nil
}

// Http1ProtocolOptions_HeaderKeyFormatValidationError is the validation error
// returned by Http1ProtocolOptions_HeaderKeyFormat.Validate if the designated
// constraints aren't met.
type Http1ProtocolOptions_HeaderKeyFormatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Http1ProtocolOptions_HeaderKeyFormatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Http1ProtocolOptions_HeaderKeyFormatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Http1ProtocolOptions_HeaderKeyFormatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Http1ProtocolOptions_HeaderKeyFormatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Http1ProtocolOptions_HeaderKeyFormatValidationError) ErrorName() string {
	return "Http1ProtocolOptions_HeaderKeyFormatValidationError"
}

// Error satisfies the builtin error interface
func (e Http1ProtocolOptions_HeaderKeyFormatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttp1ProtocolOptions_HeaderKeyFormat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Http1ProtocolOptions_HeaderKeyFormatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Http1ProtocolOptions_HeaderKeyFormatValidationError{}

// Validate checks the field values on
// Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError is the
// validation error returned by
// Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.Validate if the
// designated constraints aren't met.
type Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError) ErrorName() string {
	return "Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError"
}

// Error satisfies the builtin error interface
func (e Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttp1ProtocolOptions_HeaderKeyFormat_ProperCaseWords.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWordsValidationError{}

// Validate checks the field values on
// Http1ProtocolOptions_HeaderKeyFormat_Custom with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Http1ProtocolOptions_HeaderKeyFormat_Custom) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Rules

	return nil
}

// Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError is the validation
// error returned by Http1ProtocolOptions_HeaderKeyFormat_Custom.Validate if
// the designated constraints aren't met.
type Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError) ErrorName() string {
	return "Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError"
}

// Error satisfies the builtin error interface
func (e Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttp1ProtocolOptions_HeaderKeyFormat_Custom.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Http1ProtocolOptions_HeaderKeyFormat_CustomValidationError{}

// Validate checks the field values on Http2ProtocolOptions_SettingsParameter
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *Http2ProtocolOptions_SettingsParameter) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetIdentifier(); wrapper != nil {

		if val := wrapper.GetValue(); val < 0 || val > 65535 {
			return Http2ProtocolOptions_SettingsParameterValidationError{
				field:  "Identifier",
				reason: "value must be inside range [0, 65535]",
			}
		}

	} else {
		return Http2ProtocolOptions_SettingsParameterValidationError{
			field:  "Identifier",
			reason: "value is required and must not be nil.",
		}
	}

	if m.GetValue() == nil {
		return Http2ProtocolOptions_SettingsParameterValidationError{
			field:  "Value",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Http2ProtocolOptions_SettingsParameterValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Http2ProtocolOptions_SettingsParameterValidationError is the validation
// error returned by Http2ProtocolOptions_SettingsParameter.Validate if the
// designated constraints aren't met.
type Http2ProtocolOptions_SettingsParameterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Http2ProtocolOptions_SettingsParameterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Http2ProtocolOptions_SettingsParameterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Http2ProtocolOptions_SettingsParameterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Http2ProtocolOptions_SettingsParameterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Http2ProtocolOptions_SettingsParameterValidationError) ErrorName() string {
	return "Http2ProtocolOptions_SettingsParameterValidationError"
}

// Error satisfies the builtin error interface
func (e Http2ProtocolOptions_SettingsParameterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttp2ProtocolOptions_SettingsParameter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Http2ProtocolOptions_SettingsParameterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Http2ProtocolOptions_SettingsParameterValidationError{}
