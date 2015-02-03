// +build noos

package delay

import "rtos"

func millisec(ms int) {
	if ms == 0 {
		return
	}
	rtos.SleepUntil(rtos.Uptime() + uint64(ms)*1e6)
}
