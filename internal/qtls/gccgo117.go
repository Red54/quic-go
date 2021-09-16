// +build go1.17
// +build gccgo

package qtls

import (
	_ "unsafe"
)

//go:linkname cipherSuiteTLS13ByID github_0com_1marten_x2dseemann_1qtls_x2dgo1_x2d17.cipherSuiteTLS13ByID
func cipherSuiteTLS13ByID(id uint16) *cipherSuiteTLS13
