package healthcheck

var codeOk = 200
var codeFail = 500

var isHealthy = true

// Run healthcheck routines
func Run() int {
	if isHealthy {
		return codeOk
	}
	return codeFail
}

func setHealthy() {
	isHealthy = true
}

func setUnhealthy() {
	isHealthy = false
}
