//revive:disable:package-comments
package principal

import (
	"errors"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"buf.build/gen/go/authaas/identity/protocolbuffers/go/identity"
)

const canonical = "01234567-89ab-4def-8123-456789abcdef"

func TestKey(t *testing.T) {
	t.Run("holds a canonical id", func(t *testing.T) {
		id, err := Key(identity.Principal_builder{Id: canonical}.Build())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !id.Valid {
			t.Error("expected a valid key")
		}
	})

	t.Run("refuses what it cannot hold", func(t *testing.T) {
		if _, err := Key(identity.Principal_builder{Id: "not-a-uuid"}.Build()); err == nil {
			t.Fatal("expected error")
		}
	})
}

func TestFromKey(t *testing.T) {
	t.Run("answers the canonical id", func(t *testing.T) {
		id, err := Key(identity.Principal_builder{Id: canonical}.Build())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		principal, err := FromKey(id)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if got := principal.GetId(); got != canonical {
			t.Errorf("id = %q, want %q", got, canonical)
		}
	})

	t.Run("refuses a key with no value", func(t *testing.T) {
		if _, err := FromKey(pgtype.UUID{}); !errors.Is(err, ErrInvalidKey) {
			t.Fatalf("error = %v, want %v", err, ErrInvalidKey)
		}
	})
}
