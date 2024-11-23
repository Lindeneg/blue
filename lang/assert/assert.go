package assert

import "log"

// A calls log.Fatalf with m and args, if c is true
func A(c bool, m string, args ...any) {
	if c {
		log.Fatalf(m, args...)
	}
}

// VerifyNotReached calls log.Fatalln if this function is called
// Useful to verify that a certain path or branch was not reached.
func VerifyNotReached() {
	log.Fatalln("unwanted path reached")
}
