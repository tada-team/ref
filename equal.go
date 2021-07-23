package ref

import "time"

func IntEqual(v1, v2 *int) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}
	return *v1 == *v2
}

func Uint16Equal(v1, v2 *uint16) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}
	return *v1 == *v2
}

func TimeEqual(v1, v2 *time.Time) bool {
	if v1 == nil && v2 == nil {
		return true
	}
	if v1 == nil || v2 == nil {
		return false
	}
	return v1.Equal(*v2)
}
