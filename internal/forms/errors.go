package forms

type errors map[string][]string

func (error errors) Add(field, message string) {
	error[field] = append(error[field], message)
}

func (error errors) Get(field string) string {
	messages := error[field]
	if len(messages) == 0 { return "" }

	return messages[0]
}
