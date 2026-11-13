package utils

import "regexp"

// sensitiveJSONPattern matches JSON string values of well-known credential
// keys so they can be masked before they reach logs.
var sensitiveJSONPattern = regexp.MustCompile(`(?i)("(?:access_key|access_key_id|secret_key|secret_access_key|password|passwd|api_key|apikey|token)"\s*:\s*)"[^"]*"`)

// RedactSensitive masks the values of credential-like JSON fields in the
// given payload. Used to keep secrets (e.g. Stook bucket access keys) out of
// debug logs.
func RedactSensitive(payload string) string {
	return sensitiveJSONPattern.ReplaceAllString(payload, `${1}"***"`)
}
