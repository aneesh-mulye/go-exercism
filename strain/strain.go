package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](coll []T, pred func(T) bool) []T {
	res := make([]T, 0)
	for _, elem := range coll {
		if pred(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func Discard[T any](coll []T, pred func(T) bool) []T {
	revPred := func(t T) bool {
		return !pred(t)
	}
	return Keep(coll, revPred)
}
