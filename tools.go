package main

func normalizeLimit(limitRef *int) int {
	var limit int
	if limitRef == nil {
		limit = 5
	} else {
		limit = *limitRef
	}
	return limit
}
