package progress

import "errors"

func GetPercentage(parts, total float64) (float64, error) {
	if total == 0 {
		return 0, errors.New("division by zero")
	}

	return parts / total * 100, nil
}
