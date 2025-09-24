package random_test

import (
	"testing"

	"github.com/go4x/goal/random"
	"github.com/go4x/got"
)

func TestInt(t *testing.T) {
	tr := got.New(t, "test Int")
	tr.Case("give 1, should get 0")
	r := random.Int(1)
	tr.Require(r == 0, "result should be 0")

	for _, i := range []int{-1, 0} {
		tr.Case("give %d, should panic", i)
		func() {
			defer func() {
				err := recover()
				tr.Require(err != nil, "should panic")
			}()
			random.Int(i)
		}()
	}

	for _, i := range []int{2, 10, 99, 999, 9999} {
		tr.Case("give %d, should get a random number which < %d", i, i)
		tr.Require(random.Int(i) < i, "result should < %d", i)
	}
}

func TestBetween(t *testing.T) {
	tr := got.New(t, "test Between")
	tr.Case("give 0,0, get 0")
	r := random.Between(0, 0)
	tr.Require(r == 0, "should get 0")

	cases := [][]int{{-1, -1}, {-1, 0}, {0, -1}}
	for _, c := range cases {
		tr.Case("give invalid param: %#v, will panic", c)
		func() {
			defer func() {
				err := recover()
				tr.Require(err != nil, "should panic")
			}()
			random.Between(c[0], c[1])
		}()
	}

	cases = [][]int{{0, 1}, {1, 100}, {100, 9999}}
	for _, c := range cases {
		tr.Case("give param: %#v, will get correct result", c)
		r := random.Between(c[0], c[1])
		tr.Require(r >= c[0] && r < c[1], "result should >= %d < %d", c[0], c[1])
	}
}

func TestFloat64(t *testing.T) {
	tr := got.New(t, "test Float64")
	tr.Case("generate 1000 random floats")
	for i := 0; i < 1000; i++ {
		f := random.Float64()
		tr.Require(f >= 0.0 && f < 1.0, "float should be in range [0, 1), got %f", f)
	}
}

func TestFloat64Between(t *testing.T) {
	tr := got.New(t, "test Float64Between")

	// Test invalid parameters
	tr.Case("test invalid parameters")
	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.Float64Between(5.0, 3.0)
	}()

	// Test valid parameters
	tr.Case("test valid parameters")
	for i := 0; i < 1000; i++ {
		f := random.Float64Between(1.5, 3.7)
		tr.Require(f >= 1.5 && f < 3.7, "float should be in range [1.5, 3.7), got %f", f)
	}
}

func TestBool(t *testing.T) {
	tr := got.New(t, "test Bool")
	tr.Case("generate 1000 random booleans")
	trueCount := 0
	for i := 0; i < 1000; i++ {
		if random.Bool() {
			trueCount++
		}
	}
	// Should be roughly 50%
	tr.Require(trueCount > 400 && trueCount < 600, "true count should be around 500, got %d", trueCount)
}

func TestChoice(t *testing.T) {
	tr := got.New(t, "test Choice")

	// Test empty slice
	tr.Case("test empty slice")
	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.Choice([]string{})
	}()

	// Test valid slice
	tr.Case("test valid slice")
	slice := []string{"A", "B", "C", "D", "E"}
	for i := 0; i < 100; i++ {
		choice := random.Choice(slice)
		found := false
		for _, s := range slice {
			if s == choice {
				found = true
				break
			}
		}
		tr.Require(found, "choice should be from the slice, got %s", choice)
	}
}

func TestShuffle(t *testing.T) {
	tr := got.New(t, "test Shuffle")
	tr.Case("test shuffle preserves length and elements")

	original := []int{1, 2, 3, 4, 5}
	shuffled := make([]int, len(original))
	copy(shuffled, original)

	random.Shuffle(shuffled)

	tr.Require(len(shuffled) == len(original), "length should be preserved")

	// Check that all elements are still present
	for _, elem := range original {
		found := false
		for _, shuffledElem := range shuffled {
			if elem == shuffledElem {
				found = true
				break
			}
		}
		tr.Require(found, "element %d should still be present", elem)
	}
}

func TestSample(t *testing.T) {
	tr := got.New(t, "test Sample")
	tr.Case("test sample")

	slice := []string{"A", "B", "C", "D", "E"}
	k := 3

	sampled := random.Sample(slice, k)
	tr.Require(len(sampled) == k, "sample length should be %d, got %d", k, len(sampled))

	// Test invalid parameters
	tr.Case("test invalid parameters")
	tr.Require(random.Sample(slice, 0) == nil, "sample with k=0 should return nil")
	tr.Require(random.Sample(slice, 10) == nil, "sample with k > len should return nil")
}

func TestSelectWeighted(t *testing.T) {
	tr := got.New(t, "test SelectWeighted")
	tr.Case("test weighted choice")

	choices := []random.WeightedChoice[string]{
		{Weight: 1, Value: "A"},
		{Weight: 2, Value: "B"},
		{Weight: 3, Value: "C"},
	}

	for i := 0; i < 100; i++ {
		choice := random.SelectWeighted(choices)
		valid := false
		for _, c := range choices {
			if c.Value == choice {
				valid = true
				break
			}
		}
		tr.Require(valid, "choice should be valid, got %s", choice)
	}

	// Test empty choices
	tr.Case("test empty choices")
	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.SelectWeighted([]random.WeightedChoice[string]{})
	}()
}

func TestPercent(t *testing.T) {
	tr := got.New(t, "test Percent")
	tr.Case("test percent probability")

	// Test edge cases
	tr.Require(!random.Percent(0), "0% should return false")
	tr.Require(random.Percent(100), "100% should return true")

	// Test invalid parameters
	tr.Case("test invalid parameters")
	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.Percent(-1)
	}()

	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.Percent(101)
	}()
}

func TestPercentFloat(t *testing.T) {
	tr := got.New(t, "test PercentFloat")
	tr.Case("test percent float probability")

	// Test edge cases
	tr.Require(!random.PercentFloat(0.0), "0.0 should return false")
	tr.Require(random.PercentFloat(1.0), "1.0 should return true")

	// Test invalid parameters
	tr.Case("test invalid parameters")
	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.PercentFloat(-0.1)
	}()

	func() {
		defer func() {
			err := recover()
			tr.Require(err != nil, "should panic")
		}()
		random.PercentFloat(1.1)
	}()
}

// Fuzz tests
func FuzzInt(f *testing.F) {
	seeds := []int{1, 100, 999}
	for _, s := range seeds {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, max int) {
		if max > 0 {
			i := random.Int(max)
			if i < 0 || i >= max {
				t.Errorf("test failed, result should be >= 0 and < %d, but is %d", max, i)
			}
		}
	})
}

func FuzzBetween(f *testing.F) {
	f.Add(-1, 0)
	f.Add(0, -1)
	f.Add(0, 0)
	f.Add(1, 1)
	f.Add(1, 100)
	f.Add(100, 999)
	f.Fuzz(func(t *testing.T, min int, max int) {
		if min >= 0 && max >= 0 && min <= max {
			i := random.Between(min, max)
			if min == max {
				if i != min {
					t.Errorf("test failed, result should = %d, but is %d", min, i)
				}
			} else if i < min || i >= max {
				t.Errorf("test failed, result should be >= %d and < %d, but is %d", min, max, i)
			}
		}
	})
}
