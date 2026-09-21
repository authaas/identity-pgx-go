//revive:disable:package-comments
package principal

import (
	"github.com/jackc/pgx/v5/pgtype"

	"buf.build/gen/go/authaas/identity/protocolbuffers/go/identity"
)

// Key answers with the stored form of a principal's id, or the type's own
// refusal of a value it cannot hold.
func Key(principal *identity.Principal) (pgtype.UUID, error) {
	var id pgtype.UUID

	err := id.Scan(principal.GetId())

	return id, err
}
