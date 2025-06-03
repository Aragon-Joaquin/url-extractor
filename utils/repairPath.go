package utils

import (
	"strings"
)

func RepairPath(url string, result string) string {
	builder := result

	//! eliminate queries if exist
	queryRemover := strings.Index(builder, "?")
	if queryRemover >= 0 {
		builder = builder[:queryRemover]
	}

	//! add domain if doesn't exists
	for idx, val := range URLProtocols {

		if strings.HasPrefix(result, val) {
			break
		}

		if idx == len(URLProtocols)-1 {
			builder = url + builder
		}
	}

	//! delete www.
	if strings.Contains(builder, WWWPrefix) {
		builder = strings.Replace(builder, WWWPrefix, "", 1)
	}

	return builder
}
