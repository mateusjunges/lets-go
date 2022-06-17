package forms

type errors map[string][]string

// Add adds an error message for the given key
func (e errors) Add(key string, message string) {
	e[key] = append(e[key], message)
}

// Get returns the first error message for the given key
func (e errors) Get(key string) string {
	errorString := e[key]

	if len(errorString) == 0 {
		return ""
	}

	return errorString[0]
}
