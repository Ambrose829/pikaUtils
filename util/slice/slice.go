package slice

import "strings"

func SliceToString(list []string) (res string) {
	for _, s := range list {
		if s != "" {
			if res != "" {
				res += ","
			}
			res += s
		}
	}
	return
}

func IsIntInSlice(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func IsInt64InSlice(slice []int64, target int64) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func IsStringInSlice(slice []string, target string) bool {
	for _, element := range slice {
		if element == target {
			return true
		}
	}
	return false
}

func TrimSliceSpace(list []string) (res []string) {
	for _, s := range list {
		if strings.TrimSpace(s) != "" {
			res = append(res, s)
		}
	}
	return
}

