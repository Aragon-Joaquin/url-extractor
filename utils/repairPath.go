package utils

import (
	"strings"
)

func RepairPath(currentUrl string, result string) string {
	builder := result

	//! add domain if doesn't exists
	for idx, val := range URLProtocols {

		if strings.HasPrefix(result, val) {
			break
		}

		if idx == len(URLProtocols)-1 {
			builder = currentUrl + builder
		}
	}

	//! delete www.
	if strings.Contains(builder, WWWPrefix) {
		builder = strings.Replace(builder, WWWPrefix, "", 1)
	}

	//! eliminate queries if exist
	queryRemover := strings.Index(builder, "?")
	if queryRemover >= 0 {
		builder = builder[:queryRemover]
	}

	//! eliminate self referencing
	secondRemover := strings.Index(builder, "#")
	if secondRemover >= 0 {
		builder = builder[:secondRemover]
	}

	return builder
}
