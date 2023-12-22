//go:build js
// +build js

package tea

// listenForResize is a no-op on the web.
func (p *Program) listenForResize(done chan struct{}) {
	close(done)
}
