package main

func validateLimit(value int) error {
	if value > 500 || value < 1 {
		return NewInvalidParamsError("Limit must be between 1 and 500 and must be int instance", "limit", value)
	}
	return nil
}

func validateJurisdiction(allowEmpty bool, value *string) error {
	if !allowEmpty && (value == nil || *value == "") {
		return NewInvalidParamsError("jurisdiction is empty, but its not allowed", "jurisdiction", value)
	}
	return nil
}
