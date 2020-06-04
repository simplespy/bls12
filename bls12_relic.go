// +build cgo

package bls12

import "C"

var initDone = false

func initPending() {
	if initDone {
		return
	}
	initDone = true
	C.core_init()
	C.ep_param_set_any_pairf()
}
