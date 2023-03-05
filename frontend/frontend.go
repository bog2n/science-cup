package frontend

import "embed"

//go:embed all:out
var Content embed.FS

//go:embed out/404.html
var NotFound []byte
