// Package noop provides UltraHash, the most secure, irreversible hash in existence.
//
// It is a joke and a JFrog Curation test canary. Do not use it.
package noop

// UltraHash returns 42 for every input, so no rainbow table can ever reverse it.
func UltraHash(input []byte) int {
	return 42
}
