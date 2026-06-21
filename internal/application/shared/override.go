package shared

func Override[T any](base, override *T) *T {
	if override != nil {
		return override
	}

	return base
}
