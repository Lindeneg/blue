package assert

import "log"

func A(c bool, m string, args ...any) {
	if c {
		log.Fatalf(m, args...)
	}
}

func VerifyNotReached() {
	log.Fatalln("unwanted path reached")
}
