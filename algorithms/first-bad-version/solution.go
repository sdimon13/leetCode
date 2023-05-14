package first_bad_version

// VersionControl is a struct which has the API to identify whether a version is bad
type VersionControl struct {
	badVersion int
}

// NewVersionControl initializes a VersionControl with a first bad version
func NewVersionControl(n int) *VersionControl {
	return &VersionControl{badVersion: n}
}

func (v *VersionControl) firstBadVersion(n int, isBadVersion func(int) bool) int {
	left, right := 1, n

	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
