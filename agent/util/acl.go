package util

import "errors"

func GetBit(bit string, position int) (bool, error) {
	if len(bit) > position && position > -1 {
		if bit[position:position+1] == "1" {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("out of range")
}
