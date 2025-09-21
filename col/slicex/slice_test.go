package slicex_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go4x/goal/col/slicex"
	"github.com/go4x/got"
)

var testSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestNew(t *testing.T) {
	s := slicex.New[int]()
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4)
	// s = append(s, 5)
	fmt.Println(len(s), cap(s))

	s = slicex.NewSize[int](5)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	fmt.Println(len(s), cap(s))
}

func TestRetain(t *testing.T) {
	logger := got.New(t, "Retain")

	ret := slicex.From(testSlice).Retain(func(a int) bool {
		return a > 5
	}).To()
	fmt.Println(testSlice)
	expect := []int{6, 7, 8, 9}
	logger.Require(reflect.DeepEqual(expect, ret), "expect %#v, actual %#v", expect, ret)

	// ret1 := slice.Wrap([]any{"a", 1, 3.14}).Retain(func(a any) bool {
	// 	switch a.(type) {
	// 	case int:
	// 		return a.(int) > 1
	// 	default:
	// 		return true
	// 	}
	// })
	// expect1 := []any{"a", 3.14}
	// logger.Require(reflect.DeepEqual(expect1, ret1), "expect %v, actual %v", expect1, ret1)
}

func TestJoin(t *testing.T) {
	logger := got.New(t, "Join")

	s := slicex.From(testSlice).Join(",")
	logger.Require(s == "1,2,3,4,5,6,7,8,9", "%v join result should be %s", testSlice, s)

	join := slicex.From([]string{"a", "b"}).Join(".")
	logger.Require(join == "a.b", "a join b should be %s", join)

	// join = slice.Wrap([]TestData{"a", 1, "3.14"}).Join(",")
	// logger.Require(join == "a,1,3.14", "%s join %d join %.2f should be %s", "a", 1, 3.14, join)
}

func TestUnion(t *testing.T) {
	var before = testSlice

	logger := got.New(t, "Union")
	sl := []int{1, 2, 3, 4, 5, 6, 10, 11}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	ret := slicex.From(testSlice).Union(sl).To()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestIntersect(t *testing.T) {
	var before = testSlice

	logger := got.New(t, "Intersect")
	sl := []int{1, 2, 3, 10, 11}
	want := []int{1, 2, 3}
	ret := slicex.From(testSlice).Intersect(sl).To()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestRemove(t *testing.T) {
	var before = testSlice

	logger := got.New(t, "Remove")
	sl := []int{1, 2, 3, 10, 11}
	want := []int{4, 5, 6, 7, 8, 9}
	ret := slicex.From(testSlice).Remove(sl).To()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestDiff(t *testing.T) {
	var before = testSlice

	logger := got.New(t, "Diff")
	sl := []int{1, 2, 3, 10, 11}
	want := []int{4, 5, 6, 7, 8, 9, 10, 11}
	ret := slicex.From(testSlice).Diff(sl).To()
	logger.Require(reflect.DeepEqual(testSlice, before), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result is correct")
}

func TestDelete(t *testing.T) {
	raw := testSlice

	logger := got.New(t, "Delete")
	want := []int{4, 5, 6, 7, 8, 9}
	var ret = slicex.From(testSlice).Delete(1, 2, 3).To()
	logger.Require(reflect.DeepEqual(testSlice, raw), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")

	want = []int{2, 4, 6, 8}
	ret = slicex.From(testSlice).Delete(1, 3, 5, 7, 9).To()
	logger.Require(reflect.DeepEqual(testSlice, raw), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")
}

func TestRemoveDuplicate(t *testing.T) {
	logger := got.New(t, "RemoveDuplicate")

	raw := []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7}
	ret := slicex.From(raw).RemoveDuplicate().To()
	want := []int{1, 2, 3, 4, 5, 6, 7}
	logger.Require(reflect.DeepEqual(raw, []int{1, 1, 2, 2, 3, 4, 5, 5, 6, 7, 7}), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret, want), "result correct")
}

func TestChainInvoke(t *testing.T) {
	logger := got.New(t, "ChainInvoke")

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// not use Raw()
	ret := slicex.From(testSlice).Union([]int{1, 2, 3, 4, 10}).RemoveDuplicate()
	logger.Require(slicex.Equal(testSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}), "raw slice should not be changed")
	logger.Require(slicex.Equal(ret, want), "correct result")

	want = []int{1, 2, 3, 10}
	// use Raw()
	ret1 := slicex.From(testSlice).Union([]int{1, 2, 3, 4, 10}).Intersect([]int{1, 2, 3, 10, 11, 12}).To()
	logger.Require(reflect.DeepEqual(testSlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}), "raw slice should not be changed")
	logger.Require(reflect.DeepEqual(ret1, want), "correct result")
}

func TestSort(t *testing.T) {
	logger := got.New(t, "Sort")
	var ss = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	original := make([]int, len(ss))
	copy(original, ss) // Keep a copy of original

	s := slicex.From(ss)
	sr := s.Sort(func(i, j int) bool {
		return i < j
	})

	// Verify original slice is not modified
	logger.Require(slicex.Equal(ss, original), "original slice should not be modified after Sort")

	// Verify sorted result
	expectedSorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	logger.Require(slicex.Equal(sr.To(), expectedSorted), "sorted result should be correct")

	// Test reverse
	reversed := sr.Reverse()
	expectedReversed := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	logger.Require(slicex.Equal(reversed, expectedReversed), "reversed result should be correct")
}

type user struct {
	name  string
	age   int
	score float32
}

func TestSortObj(t *testing.T) {
	users := []user{
		{name: "user", age: 18, score: 99.5},
		{name: "hello", age: 16, score: 100},
		{name: "lily", age: 17, score: 99.5},
		{name: "abc", age: 17, score: 99.5},
	}

	s := slicex.From(users)
	sr := s.Sort(func(a, b user) bool {
		if a.score != b.score {
			return a.score > b.score
		}
		if a.age != b.age {
			return a.age < b.age
		}
		if a.name != b.name {
			return a.name < b.name
		}
		return false
	})
	fmt.Printf("%v\n", s)
	sr.Reverse()
	fmt.Printf("%v\n", s)
}

// TestEqual tests the Equal function
func TestEqual(t *testing.T) {
	logger := got.New(t, "Equal")

	// Test equal slices
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	logger.Require(slicex.Equal(s1, s2), "equal slices should return true")

	// Test different lengths
	s3 := []int{1, 2, 3}
	logger.Require(!slicex.Equal(s1, s3), "different length slices should return false")

	// Test different elements
	s4 := []int{1, 2, 3, 4, 6}
	logger.Require(!slicex.Equal(s1, s4), "different elements should return false")

	// Test empty slices
	s5 := []int{}
	s6 := []int{}
	logger.Require(slicex.Equal(s5, s6), "empty slices should return true")

	// Test string slices
	s7 := []string{"a", "b", "c"}
	s8 := []string{"a", "b", "c"}
	logger.Require(slicex.Equal(s7, s8), "equal string slices should return true")
}

// TestEqualFunc tests the EqualFunc function
func TestEqualFunc(t *testing.T) {
	logger := got.New(t, "EqualFunc")

	// Test equal slices with custom comparison
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	logger.Require(slicex.EqualFunc(s1, s2, func(a, b int) bool { return a == b }), "equal slices should return true")

	// Test different lengths
	s3 := []int{1, 2, 3}
	logger.Require(!slicex.EqualFunc(s1, s3, func(a, b int) bool { return a == b }), "different length slices should return false")

	// Test different elements
	s4 := []int{1, 2, 3, 4, 6}
	logger.Require(!slicex.EqualFunc(s1, s4, func(a, b int) bool { return a == b }), "different elements should return false")

	// Test custom comparison function
	s5 := []int{1, 2, 3}
	s6 := []int{2, 3, 4}
	logger.Require(slicex.EqualFunc(s5, s6, func(a, b int) bool { return a+1 == b }), "custom comparison should work")

	// Test different types
	s7 := []int{1, 2, 3}
	s8 := []float64{1.0, 2.0, 3.0}
	logger.Require(slicex.EqualFunc(s7, s8, func(a int, b float64) bool { return float64(a) == b }), "different types should work")
}

// TestNewSize tests the NewSize function
func TestNewSize(t *testing.T) {
	logger := got.New(t, "NewSize")

	// Test with positive size
	s := slicex.NewSize[int](5)
	logger.Require(len(s) == 5, "length should be 5")
	logger.Require(cap(s) == 5, "capacity should be 5")

	// Test with zero size
	s2 := slicex.NewSize[int](0)
	logger.Require(len(s2) == 0, "length should be 0")
	logger.Require(cap(s2) == 0, "capacity should be 0")

	// Test with string type
	s3 := slicex.NewSize[string](3)
	logger.Require(len(s3) == 3, "string slice length should be 3")
}

// TestFrom tests the From function
func TestFrom(t *testing.T) {
	logger := got.New(t, "From")

	// Test with valid slice
	s := []int{1, 2, 3, 4, 5}
	result := slicex.From(s)
	logger.Require(slicex.Equal(result, s), "From should return the same slice")

	// Test with empty slice
	s2 := []int{}
	result2 := slicex.From(s2)
	logger.Require(slicex.Equal(result2, s2), "From should work with empty slice")

	// Test with string slice
	s3 := []string{"a", "b", "c"}
	result3 := slicex.From(s3)
	logger.Require(slicex.Equal(result3, s3), "From should work with string slice")
}

// TestTo tests the To method
func TestTo(t *testing.T) {
	logger := got.New(t, "To")

	s := slicex.From([]int{1, 2, 3, 4, 5})
	result := s.To()
	logger.Require(slicex.Equal(result, []int{1, 2, 3, 4, 5}), "To should return the underlying slice")
}

// TestFilter tests the Filter method
func TestFilter(t *testing.T) {
	logger := got.New(t, "Filter")

	s := slicex.From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	result := s.Filter(func(x int) bool {
		return x > 5
	}).To()

	expected := []int{1, 2, 3, 4, 5}
	logger.Require(slicex.Equal(result, expected), "Filter should remove elements matching condition")
}

// TestContain tests the Contain method
func TestContain(t *testing.T) {
	logger := got.New(t, "Contain")

	s := slicex.From([]int{1, 2, 3, 4, 5})

	// Test existing element
	logger.Require(s.Contain(3), "should contain existing element")

	// Test non-existing element
	logger.Require(!s.Contain(6), "should not contain non-existing element")

	// Test empty slice
	s2 := slicex.New[int]()
	logger.Require(!s2.Contain(1), "empty slice should not contain any element")
}

// TestClip tests the Clip method
func TestClip(t *testing.T) {
	logger := got.New(t, "Clip")

	// Create a slice with extra capacity
	s := make([]int, 3, 10)
	s[0], s[1], s[2] = 1, 2, 3

	slicex_s := slicex.From(s)
	result := slicex_s.Clip()

	logger.Require(len(result) == 3, "length should be 3")
	logger.Require(cap(result) == 3, "capacity should be 3")
	logger.Require(slicex.Equal(result, []int{1, 2, 3}), "elements should be preserved")

	// Verify that modifying the original slice doesn't affect the clipped result
	s[0] = 999
	logger.Require(slicex.Equal(result, []int{1, 2, 3}), "clipped result should be independent of original")

	// Test with empty slice
	empty := slicex.New[int]()
	emptyClipped := empty.Clip()
	logger.Require(len(emptyClipped) == 0, "empty slice clip should have length 0")
	logger.Require(cap(emptyClipped) == 0, "empty slice clip should have capacity 0")
}

// TestNewSortableSlice tests the NewSortableSlice function
func TestNewSortableSlice(t *testing.T) {
	logger := got.New(t, "NewSortableSlice")

	s := slicex.From([]int{3, 1, 4, 1, 5})
	less := func(x, y int) bool { return x < y }

	sortable := slicex.NewSortableSlice(s, less)

	logger.Require(sortable.Len() == 5, "length should be 5")
	logger.Require(sortable.Less(1, 0), "1 should be less than 3")
	logger.Require(!sortable.Less(0, 1), "3 should not be less than 1")
}

// TestSortableSliceLen tests the Len method
func TestSortableSliceLen(t *testing.T) {
	logger := got.New(t, "SortableSliceLen")

	s := slicex.From([]int{1, 2, 3})
	less := func(x, y int) bool { return x < y }
	sortable := slicex.NewSortableSlice(s, less)

	logger.Require(sortable.Len() == 3, "length should be 3")
}

// TestSortableSliceLess tests the Less method
func TestSortableSliceLess(t *testing.T) {
	logger := got.New(t, "SortableSliceLess")

	s := slicex.From([]int{3, 1, 4})
	less := func(x, y int) bool { return x < y }
	sortable := slicex.NewSortableSlice(s, less)

	logger.Require(sortable.Less(1, 0), "1 should be less than 3")
	logger.Require(!sortable.Less(0, 1), "3 should not be less than 1")
}

// TestSortableSliceSwap tests the Swap method
func TestSortableSliceSwap(t *testing.T) {
	logger := got.New(t, "SortableSliceSwap")

	s := slicex.From([]int{1, 2, 3})
	less := func(x, y int) bool { return x < y }
	sortable := slicex.NewSortableSlice(s, less)

	// Swap elements at index 0 and 2
	sortable.Swap(0, 2)

	expected := []int{3, 2, 1}
	logger.Require(slicex.Equal(sortable.To(), expected), "elements should be swapped")
}

// TestSortableSliceTo tests the To method of SortableSlice
func TestSortableSliceTo(t *testing.T) {
	logger := got.New(t, "SortableSliceTo")

	s := slicex.From([]int{1, 2, 3})
	less := func(x, y int) bool { return x < y }
	sortable := slicex.NewSortableSlice(s, less)

	result := sortable.To()
	expected := []int{1, 2, 3}
	logger.Require(slicex.Equal(result, expected), "To should return the underlying slice")
}

// TestEdgeCases tests various edge cases
func TestEdgeCases(t *testing.T) {
	logger := got.New(t, "EdgeCases")

	// Test empty slice operations
	empty := slicex.New[int]()

	// Test Retain on empty slice
	result := empty.Retain(func(x int) bool { return x > 0 })
	logger.Require(len(result) == 0, "retain on empty slice should return empty")

	// Test Filter on empty slice
	result = empty.Filter(func(x int) bool { return x > 0 })
	logger.Require(len(result) == 0, "filter on empty slice should return empty")

	// Test Join on empty slice
	joinResult := empty.Join(",")
	logger.Require(joinResult == "", "join on empty slice should return empty string")

	// Test Union with empty slice
	nonEmpty := slicex.From([]int{1, 2, 3})
	unionResult := empty.Union(nonEmpty.To())
	logger.Require(slicex.Equal(unionResult, nonEmpty), "union with empty slice should return non-empty")

	// Test Intersect with empty slice
	intersectResult := empty.Intersect(nonEmpty.To())
	logger.Require(len(intersectResult) == 0, "intersect with empty slice should return empty")

	// Test Diff with empty slice
	diffResult := empty.Diff(nonEmpty.To())
	logger.Require(slicex.Equal(diffResult, nonEmpty), "diff with empty slice should return non-empty")

	// Test RemoveDuplicate on empty slice
	removeDupResult := empty.RemoveDuplicate()
	logger.Require(len(removeDupResult) == 0, "remove duplicate on empty slice should return empty")

	// Test Delete on empty slice
	deleteResult := empty.Delete(1, 2, 3)
	logger.Require(len(deleteResult) == 0, "delete on empty slice should return empty")

	// Test Clip on empty slice
	clipResult := empty.Clip()
	logger.Require(len(clipResult) == 0, "clip on empty slice should return empty")
}

// TestComplexOperations tests complex operations
func TestComplexOperations(t *testing.T) {
	logger := got.New(t, "ComplexOperations")

	// Test chaining multiple operations
	s := slicex.From([]int{1, 2, 2, 3, 3, 4, 5, 5, 6})

	// Chain: Retain -> RemoveDuplicate -> Sort
	result := s.Retain(func(x int) bool { return x > 2 }).
		RemoveDuplicate().
		Sort(func(x, y int) bool { return x < y })

	expected := []int{3, 4, 5, 6}
	logger.Require(slicex.Equal(result.To(), expected), "chained operations should work correctly")

	// Test Reverse after Sort
	reversed := result.Reverse()
	expectedReversed := []int{6, 5, 4, 3}
	logger.Require(slicex.Equal(reversed, expectedReversed), "reverse should work after sort")
}

// TestStringOperations tests string-specific operations
func TestStringOperations(t *testing.T) {
	logger := got.New(t, "StringOperations")

	// Test Join with strings
	s := slicex.From([]string{"hello", "world", "test"})
	joinResult := s.Join(" ")
	expected := "hello world test"
	logger.Require(joinResult == expected, "join should work with strings")

	// Test Join with empty separator
	joinResult2 := s.Join("")
	expected2 := "helloworldtest"
	logger.Require(joinResult2 == expected2, "join with empty separator should work")

	// Test Join with single element
	single := slicex.From([]string{"single"})
	joinResult3 := single.Join(",")
	expected3 := "single"
	logger.Require(joinResult3 == expected3, "join with single element should work")
}

// TestFloatOperations tests float-specific operations
func TestFloatOperations(t *testing.T) {
	logger := got.New(t, "FloatOperations")

	// Test with float64
	s := slicex.From([]float64{1.1, 2.2, 3.3, 4.4})

	// Test Retain
	result := s.Retain(func(x float64) bool { return x > 2.0 })
	expected := []float64{2.2, 3.3, 4.4}
	logger.Require(slicex.Equal(result, expected), "retain should work with floats")

	// Test Sort
	sorted := s.Sort(func(x, y float64) bool { return x < y })
	expectedSorted := []float64{1.1, 2.2, 3.3, 4.4}
	logger.Require(slicex.Equal(sorted.To(), expectedSorted), "sort should work with floats")
}

// TestBooleanOperations tests boolean-specific operations
func TestBooleanOperations(t *testing.T) {
	logger := got.New(t, "BooleanOperations")

	// Test with bool
	s := slicex.From([]bool{true, false, true, false})

	// Test Retain
	result := s.Retain(func(x bool) bool { return x })
	expected := []bool{true, true}
	logger.Require(slicex.Equal(result, expected), "retain should work with booleans")

	// Test RemoveDuplicate
	noDup := s.RemoveDuplicate()
	expectedNoDup := []bool{true, false}
	logger.Require(slicex.Equal(noDup, expectedNoDup), "remove duplicate should work with booleans")
}

// TestPerformanceOptimizations tests the performance improvements
func TestPerformanceOptimizations(t *testing.T) {
	logger := got.New(t, "PerformanceOptimizations")

	// Test optimized Intersect with larger datasets
	largeSlice1 := make([]int, 1000)
	largeSlice2 := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		largeSlice1[i] = i
		largeSlice2[i] = i + 500 // Overlap from 500-999
	}

	s1 := slicex.From(largeSlice1)
	result := s1.Intersect(largeSlice2)

	// Should contain elements 500-999
	logger.Require(len(result) == 500, "intersect should work with large datasets")

	// Test optimized Delete with multiple elements
	toDelete := make([]int, 100)
	for i := 0; i < 100; i++ {
		toDelete[i] = i
	}

	s2 := slicex.From(largeSlice1)
	deleted := s2.Delete(toDelete...)
	logger.Require(len(deleted) == 900, "delete should work efficiently with many elements")
}

// TestSortInPlace tests the SortInPlace method
func TestSortInPlace(t *testing.T) {
	logger := got.New(t, "SortInPlace")

	var ss = []int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	s := slicex.From(ss)

	// Sort in place
	sorted := s.SortInPlace(func(i, j int) bool {
		return i < j
	})

	// Verify sorted result
	expectedSorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	logger.Require(slicex.Equal(sorted.To(), expectedSorted), "sorted result should be correct")

	// Verify original slice IS modified (this is the key difference)
	logger.Require(slicex.Equal(s.To(), expectedSorted), "original slice should be modified after SortInPlace")

	// Test reverse
	reversed := sorted.Reverse()
	expectedReversed := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	logger.Require(slicex.Equal(reversed, expectedReversed), "reversed result should be correct")
}
