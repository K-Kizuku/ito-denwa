package typeconvert

func ToPtr[T any](v T) *T {
	return &v
}

func StringSliceToInterfaceSlice(slice []string) []any {
	if len(slice) == 0 {
		return nil
	}
	interfaceSlice := make([]any, len(slice))
	for i, v := range slice {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
