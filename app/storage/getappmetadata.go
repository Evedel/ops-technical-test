package storage

// AppMetadataEntry is single element of the App metadata
type AppMetadataEntry struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastComminSHA string `json:"lastcommitsha"`
}

// AppMetadata is the basic information about this application
type AppMetadata struct {
	Entries []AppMetadataEntry `json:"myapplication"`
}

// GetAppMetadata queries and constructs basic information about this application
func GetAppMetadata() AppMetadata {
	return AppMetadata{[]AppMetadataEntry{
		AppMetadataEntry{
			Version:       "1.0",
			Description:   "pre-interview technical test",
			LastComminSHA: "abc57858585"}}}
}
