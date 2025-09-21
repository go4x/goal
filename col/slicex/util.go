package slicex

// basic functions

// Equal reports whether two slices are equal: the same length and all
// elements equal. If the lengths are different, Equal returns false.
// Otherwise, the elements are compared in increasing index order, and the
// comparison stops at the first unequal pair.
// Floating point NaNs are not considered equal.
//
// Parameters:
//   - s1: The first slice to compare
//   - s2: The second slice to compare
//
// Returns:
//   - bool: true if slices are equal, false otherwise
//
// Example:
//
//	s1 := []int{1, 2, 3}
//	s2 := []int{1, 2, 3}
//	s3 := []int{1, 2, 4}
//	fmt.Println(slicex.Equal(s1, s2)) // true
//	fmt.Println(slicex.Equal(s1, s3)) // false
func Equal[E comparable](s1, s2 []E) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// EqualFunc reports whether two slices are equal using a comparison
// function on each pair of elements. If the lengths are different,
// EqualFunc returns false. Otherwise, the elements are compared in
// increasing index order, and the comparison stops at the first index
// for which eq returns false.
//
// Parameters:
//   - s1: The first slice to compare
//   - s2: The second slice to compare
//   - eq: The comparison function that takes elements from both slices
//
// Returns:
//   - bool: true if slices are equal according to the comparison function
//
// Example:
//
//	s1 := []int{1, 2, 3}
//	s2 := []float64{1.0, 2.0, 3.0}
//	equal := slicex.EqualFunc(s1, s2, func(a int, b float64) bool {
//	    return float64(a) == b
//	})
//	fmt.Println(equal) // true
func EqualFunc[E1, E2 any](s1 []E1, s2 []E2, eq func(E1, E2) bool) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v1 := range s1 {
		v2 := s2[i]
		if !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Each iterates over each element in the slice and applies the given function to it.
// This function is useful for performing side effects on each element without collecting results.
//
// Parameters:
//   - ts: The input slice of type T
//   - f: The function to apply to each element
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	slicex.Each(numbers, func(x int) {
//	    fmt.Println(x * 2) // Prints: 2, 4, 6, 8, 10
//	})
func Each[T any](ts []T, f func(t T)) {
	for _, t := range ts {
		f(t)
	}
}

// Eachv transforms each element in the slice using the given function and returns a new slice.
// This function is similar to Each but collects the results into a new slice.
// It pre-allocates the result slice for better performance.
//
// Parameters:
//   - ts: The input slice of type T
//   - f: The transformation function that takes a T and returns a V
//
// Returns:
//   - []V: A new slice containing the transformed elements
//
// Example:
//
//	numbers := []int{1, 2, 3, 4, 5}
//	strings := slicex.Eachv(numbers, func(x int) string {
//	    return fmt.Sprintf("num_%d", x)
//	})
//	// Result: []string{"num_1", "num_2", "num_3", "num_4", "num_5"}
func Eachv[T any, V any](ts []T, f func(t T) V) []V {
	var vs = make([]V, len(ts))
	for i, t := range ts {
		vs[i] = f(t)
	}
	return vs
}

// Map creates a map from a slice using the given key extraction function.
// Each element in the slice becomes a value in the map, with the key determined by the function.
// If multiple elements produce the same key, the last one will overwrite the previous ones.
//
// Parameters:
//   - vs: The input slice of type V
//   - f: The key extraction function that takes a V and returns a K
//
// Returns:
//   - map[K]V: A map where keys are determined by the function and values are the original elements
//
// Example:
//
//	people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
//	nameMap := slicex.Map(people, func(p Person) string {
//	    return p.Name
//	})
//	// Result: map[string]Person{"Alice": {Name: "Alice", Age: 30}, "Bob": {Name: "Bob", Age: 25}}
func Map[K comparable, V any](vs []V, f func(v V) K) map[K]V {
	var m = make(map[K]V, len(vs))
	for _, v := range vs {
		m[f(v)] = v
	}
	return m
}

// Mapv creates a map from a slice using the given key-value extraction function.
// This function allows you to extract both key and value from each element, providing more flexibility than Map.
// If multiple elements produce the same key, the last one will overwrite the previous ones.
//
// Parameters:
//   - ts: The input slice of type T
//   - f: The key-value extraction function that takes a T and returns (K, V)
//
// Returns:
//   - map[K]V: A map where keys and values are determined by the function
//
// Example:
//
//	people := []Person{{Name: "Alice", Age: 30}, {Name: "Bob", Age: 25}}
//	ageMap := slicex.Mapv(people, func(p Person) (string, int) {
//	    return p.Name, p.Age
//	})
//	// Result: map[string]int{"Alice": 30, "Bob": 25}
func Mapv[K comparable, V any, T any](ts []T, f func(t T) (K, V)) map[K]V {
	var m = make(map[K]V, len(ts))
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}
