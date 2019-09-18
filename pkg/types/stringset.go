package types

// StringSet is a basic set implementation for strings.
// This is used a lot so it deserves its own type.
// Other types of sets are used throughout the code but do not have
// their own typedef.
// String sets and <type>sets should be used throughout the code when applicable,
// they are a lot more flexible than slices and provide easy lookup.
type StringSet map[string]struct{}

// MapStringSet is a Map of StringSets.
type MapStringSet map[string]StringSet

// Add adds a new value to the Map.
func (mss MapStringSet) Add(n string, v string) {
	_, ok := mss[n]
	if !ok {
		mss[n] = make(StringSet)
	}
	mss[n].Set(v)
}

// Set sets key in StringSet.
func (set StringSet) Set(v string) {
	set[v] = struct{}{}
}

// Get returns true if the key exists in the set.
func (set StringSet) Get(v string) bool {
	_, exists := set[v]
	return exists
}

// Remove deletes a key from the set.
func (set StringSet) Remove(v string) {
	delete(set, v)
}

// ToSlice turns all keys into a string slice.
func (set StringSet) ToSlice() []string {
	slice := make([]string, 0, len(set))

	for v := range set {
		slice = append(slice, v)
	}

	return slice
}

// Copy copies a StringSet into a new structure of the same type.
func (set StringSet) Copy() StringSet {
	newSet := make(StringSet)

	for str := range set {
		newSet.Set(str)
	}

	return newSet
}

// SliceToStringSet creates a new StringSet from an input slice
func SliceToStringSet(in []string) StringSet {
	set := make(StringSet)

	for _, v := range in {
		set.Set(v)
	}

	return set
}

// MakeStringSet creates a new StringSet from a set of arguments
func MakeStringSet(in ...string) StringSet {
	return SliceToStringSet(in)
}
