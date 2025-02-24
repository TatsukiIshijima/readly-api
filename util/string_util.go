package util

func ToStringOrNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
