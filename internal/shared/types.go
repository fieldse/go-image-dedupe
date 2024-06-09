// Common types for the project
package shared

type ScanType int

const (
	ScanTypeNone ScanType = iota
	ScanTypeQuick
	ScanTypeMetadata
	ScanTypeFileName
	ScanTypeFileNameAndSize
	ScanTypeMD5Hash
)
