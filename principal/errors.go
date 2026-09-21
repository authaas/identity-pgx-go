//revive:disable:package-comments
package principal

import "errors"

// ErrInvalidKey reports a stored key that does not read back as text, which
// a NOT NULL uuid column never yields.
var ErrInvalidKey = errors.New("invalid principal key")
