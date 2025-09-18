package slicex

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
