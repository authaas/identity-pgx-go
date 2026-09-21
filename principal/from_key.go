//revive:disable:package-comments
package principal

import (
	"github.com/jackc/pgx/v5/pgtype"

	"buf.build/gen/go/authaas/identity/protocolbuffers/go/identity"
)

// FromKey answers with the message for a stored key.
//
// pgtype.UUID.Value has no path that returns an error: a key with no value
// answers nil, and encoding sixteen bytes to text cannot fail. The error is
// not consulted, and nil in place of the text is the one failure.
func FromKey(id pgtype.UUID) (*identity.Principal, error) {
	value, _ := id.Value()

	text, ok := value.(string)
	if !ok {
		return nil, ErrInvalidKey
	}

	return identity.Principal_builder{Id: text}.Build(), nil
}
