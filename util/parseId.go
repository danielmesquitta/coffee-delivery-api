package util

import "strconv"

func ParseId(id string) (uint, error) {
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(uintId), nil
}
