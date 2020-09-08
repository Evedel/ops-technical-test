package healthcheck

import "testing"

func TestHealthCheckIsOkWhenHealthy(t *testing.T) {

	setHealthy()

	expectedCode := codeOk
	
	receivedCode := Run()

	if receivedCode != expectedCode {
		t.Errorf("healthcheck.Run() received %v, expected %v", receivedCode, expectedCode)
	}
}

func TestHealthCheckIsFailWhenBroken(t *testing.T) {

	setUnhealthy()

	expectedCode := codeFail

	receivedCode := Run()

	if receivedCode != expectedCode {
		t.Errorf("healthcheck.Run() received %v, expected %v", receivedCode, expectedCode)
	}
}
