// Package recovery provides goroutine-safe panic recovery helpers.
//
// Every goroutine launched outside of test code should defer a recovery
// function at its top to prevent panics from propagating and killing the
// entire process — particularly important for scraper-style code where
// upstream HTML/JSON shape changes can trigger nil dereferences or index
// out-of-range panics deep inside callbacks.
package recovery

import (
	"log"
)

// Recover catches a panic in the calling goroutine and logs it with the
// given label as context. It MUST be used as `defer recovery.Recover("name")`.
//
// Example:
//
//	go func() {
//	    defer wg.Done()
//	    defer recovery.Recover("searchMovie")
//	    // ... work that may panic
//	}()
func Recover(label string) {
	if r := recover(); r != nil {
		log.Printf("[recovered panic] %s: %v", label, r)
	}
}
