package storage

import (
	"encoding/json"
	"testing"
)

func TestGetAppMetadata(t *testing.T) {
	expectedBytes, _ := json.Marshal(AppMetadata{[]AppMetadataEntry{
		AppMetadataEntry{
			Version:       "1.0",
			Description:   "pre-interview technical test",
			LastComminSHA: "abc57858585"}}})
	receivedBytes, _ := json.Marshal(GetAppMetadata())

	expected := string(expectedBytes)
	received := string(receivedBytes)
	if received != expected {
		t.Errorf("storage.GetAppMetadata() received %v, expected %v", received, expected)
	}
}
