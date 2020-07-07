// Code generated by rendertemplates. DO NOT EDIT.

package releases

// List of currently supported releases
var (
	Release115 = mustParse("1.15")
	Release114 = mustParse("1.14")
	Release113 = mustParse("1.13")
	Release112 = mustParse("1.12")
)

// GetAllKymaReleases returns all supported kyma release branches
func GetAllKymaReleases() []*SupportedRelease {
	return []*SupportedRelease{
		Release114,
		Release113,
		Release112,
	}
}

// GetNextKymaRelease returns the version of kyma currently under development
func GetNextKymaRelease() *SupportedRelease {
	return Release115
}
