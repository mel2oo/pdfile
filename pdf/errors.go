package pdf

import (
	"errors"
)

// errors
var ErrEncryptionError = errors.New("missing required encryption info")
var ErrEncryptionPasswordError = errors.New("incorrect password")
var ErrEncryptionUnsupported = errors.New("unsupported encryption")
var ErrEndOfArray = errors.New("end of array")
var ErrEndOfDictionary = errors.New("end of dictionary")
var ErrEndOfHexString = errors.New("end of hex string")
var ErrEndOfString = errors.New("end of string")
var ErrReadError = errors.New("read failed")

// format errors and abnormalities
var InvalidDictionaryKeyType = "invalid dictionary key type"
var InvalidHexStringChar = "invalid hex string character"
var InvalidNameEscapeChar = "invalid name escape character"
var InvalidOctal = "invalid octal in string"
var MissingDictionaryValue = "missing dictionary value"
var UnclosedArray = "unclosed array"
var UnclosedDictionary = "unclosed dictionary"
var UnclosedHexString = "unclosed hex string"
var UnclosedStream = "unclosed stream"
var UnclosedString = "unclosed string"
var UnclosedStringEscape = "unclosed escape in string"
var UnclosedStringOctal = "unclosed octal in string"
var UnnecessaryEscapeName = "unnecessary espace sequence in name"
var UnnecessaryEscapeString = "unnecessary espace sequence in string"
