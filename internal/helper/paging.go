package helper

import "github.com/aryayunanta-ralali/shorty/internal/consts"

// LimitDefaultValue set default value limit
func LimitDefaultValue(origin int64) int64 {
	if origin < 1 {
		return consts.PagingLimitDefaultValue
	}

	if origin > consts.PagingMaxLimit {
		return consts.PagingMaxLimit
	}

	return origin
}

// PageDefaultValue set default value page
func PageDefaultValue(origin int64) int64 {
	if origin < 1 {
		return consts.PagingPageDefaultValue
	}

	return origin
}
