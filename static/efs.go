package static

import "embed"

//go:embed "js"
var JS embed.FS

//go:embed "css"
var CSS embed.FS

// //go:embed "images"
// var Images embed.FS
