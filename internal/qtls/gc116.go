// +build go1.16
// +build !go1.17
// +build gc

package qtls

import (
	_ "unsafe"
)

//go:linkname cipherSuiteTLS13ByID github.com/marten-seemann/qtls-go1-16.cipherSuiteTLS13ByID
func cipherSuiteTLS13ByID(id uint16) *cipherSuiteTLS13
