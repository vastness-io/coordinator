package util

func MergeStringSlices(stringSlices ...[]string) []string {
	length := len(stringSlices)

	if length == 1 {
		return stringSlices[0]
	} else if length > 1 {

		var mergedSlice []string

		for _, slice := range stringSlices {
			mergedSlice = append(mergedSlice, slice...)
		}

		return mergedSlice
	}

	return nil
}
