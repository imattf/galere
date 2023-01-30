package migrations

// Use fs to insure certain files needed at runtime
// are included in builds

import "embed"

//go:embed *.sql
var FS embed.FS
