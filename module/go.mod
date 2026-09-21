module github.com/authaas/identity-pgx-go

go 1.27.1

require (
	buf.build/gen/go/authaas/identity/protocolbuffers/go v1.36.12-20260920174019-cb9ef3c3aa4c.2
	github.com/jackc/pgx/v5 v5.11.0
)

require google.golang.org/protobuf v1.36.12 // indirect
