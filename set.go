package gset

// Set is a data structure that can store any number of unique values in any order you so wish. Set’s are different
// from arrays in the sense that they only allow non-repeated, unique values within them.
type Set[T comparable] map[T]struct{}

// FromMapKey return a set from the map key.
func FromMapKey[T comparable, v any](ma map[T]v) Set[T] {
	s := Set[T]{}

	for key := range ma {
		s.Push(key)
	}

	return s
}

// FromArray return a set from the array or slice, and the element must comparable. And if in array, some key is same
// the set only store one.
func FromArray[T comparable](array []T) Set[T] {
	s := Set[T]{}
	for _, key := range array {
		s.Push(key)
	}

	return s
}

// Push will push new value to set.
func (s Set[T]) Push(vs ...T) {
	if vs != nil {
		for _, key := range vs {
			s[key] = struct{}{}
		}
	}
}

// Has return true if Set contain the target value.
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// Delete remove the values from set.
func (s Set[T]) Delete(vs ...T) {
	if vs != nil {
		for _, key := range vs {
			delete(s, key)
		}
	}

}

// MergeFrom will append value from other set to this set. The same key will be ignored.
func (s Set[T]) MergeFrom(set Set[T]) {
	for key := range set {
		s[key] = struct{}{}
	}
}

// DeleteFrom will delete value from other set, and if some key does not in this set, skill this value.
func (s Set[T]) DeleteFrom(set Set[T]) {
	for key := range set {
		delete(s, key)
	}
}

// SafeEquals like Equals, return true if this set value is in other set, and other one value in this set too.
//
// Different from Equals, the SafeEquals always check all element in set. For detail,
// see: https://en.wikipedia.org/wiki/Timing_attack.
func (s Set[T]) SafeEquals(set Set[T]) bool {
	if len(s) != len(set) {
		return false
	}
	res := true
	for key := range s {
		if !set.Has(key) {
			res = false
		}
		// loop all element for run safe
	}

	return res
}

// Equals return true if this set value is in other set, and other one value in this set too.
func (s Set[T]) Equals(set Set[T]) bool {
	if len(s) != len(set) {
		return false
	}

	// if len is same, and if all key in this set can be found in other set, is same.
	for key := range s {
		if !set.Has(key) {
			return false
		}
	}
	return true
}

func (s Set[T]) Clone() Set[T] {
	nSet := Set[T]{}
	for key := range s {
		nSet[key] = struct{}{}
	}

	return nSet
}

// IsSupperOf return true when all element in set can be found in s.
//
// Empty is all Set's sub set.
func (s Set[T]) IsSupperOf(set Set[T]) bool {
	for key := range set {
		if !s.Has(key) {
			return false
		}
	}

	return true
}

// IsSubOf return true when all element in s can be found in set.
//
// Empty is all Set's sub set.
func (s Set[T]) IsSubOf(set Set[T]) bool {
	for key := range s {
		if !set.Has(key) {
			return false
		}
	}

	return true
}

func (s Set[T]) GetMix(set Set[T]) Set[T] {
	if set == nil {
		return s
	}
	res := s.Clone()
	for key := range res {
		if !set.Has(key) {
			res.Delete(key)
		}
	}

	return res
}

func (s Set[T]) ToArray() []T {
	t := make([]T, len(s), len(s))
	if s == nil {
		return t
	}

	index := 0
	for key := range s {
		t[index] = key
		index++
	}

	return t
}

func GetMerge[T comparable](sets ...Set[T]) Set[T] {
	s := FromArray([]T{})

	if len(sets) == 0 {
		return s
	}
	for _, set := range sets {
		s.MergeFrom(set)
	}

	return s
}

func GetMix[T comparable](sets ...Set[T]) Set[T] {
	if len(sets) == 0 {
		return FromArray([]T{})
	}
	base := sets[0].Clone()
	if len(sets) == 1 {
		return base
	}
	if base == nil {
		return FromArray([]T{})
	}

	sets = sets[0:]

	for _, set := range sets {
		base = base.GetMix(set)
	}

	return base
}
