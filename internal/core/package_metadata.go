package core

type PackageMetadata struct {
	Name            string   `json:"name"`
	Version         string   `json:"version"`
	Language        string   `json:"language"`
	Dependencies    []string `json:"dependencies,omitempty"`
	MetadataVersion string   `json:"metadata_version"`
}
