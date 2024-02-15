package assets

import (
	"embed"
)

//go:embed *
var EmbeddedFiles embed.FS
