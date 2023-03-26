package gset

import (
	"reflect"
	"testing"
)

func TestSet_Clone(t *testing.T) {
	type T int
	tests := []struct {
		name string
		s    Set[T]
		want Set[T]
	}{
		{name: "1. test common", s: FromArray([]T{1, 2, 3}), want: FromArray([]T{1, 2, 3})},
		{name: "2. test empty - nil input", s: nil, want: FromArray([]T{})},
		{name: "3. test empty - empty value", s: FromArray([]T{}), want: FromArray([]T{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Delete(t *testing.T) {
	type T int
	type args struct {
		v T
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want Set[T]
	}{
		{name: "1. test common case 1", args: args{v: 1}, s: FromArray([]T{1}), want: FromArray([]T{})},
		{name: "2. test common case 2", args: args{v: 1}, s: FromArray([]T{}), want: FromArray([]T{})},
		{name: "3. test common case 3", args: args{v: 1}, s: FromArray([]T{1, 2}), want: FromArray([]T{2})},
		{name: "4. test common case 4", args: args{v: 1}, s: FromArray([]T{2}), want: FromArray([]T{2})},
		{name: "4. test common case 4", args: args{v: 1}, s: FromArray([]T{2}), want: FromArray([]T{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Delete(tt.args.v)
			if !reflect.DeepEqual(tt.want, tt.s) {
				t.Errorf("Delete(%v) = %v, want = %v", tt.args.v, tt.s, tt.want)
			}
		})
	}
}

func TestSet_DeleteFrom(t *testing.T) {
	type T int
	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want Set[T]
	}{
		{name: "1. test common - all empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: FromArray([]T{})},
		{name: "2. delete empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{1}), want: FromArray([]T{1})},
		{name: "3. delete exists", args: args{set: FromArray([]T{1})}, s: FromArray([]T{1}), want: FromArray([]T{})},
		{name: "4. delete not exists", args: args{set: FromArray([]T{2})}, s: FromArray([]T{1}), want: FromArray([]T{1})},
		{name: "5. delete supper", args: args{set: FromArray([]T{1, 2, 3})}, s: FromArray([]T{1, 2}), want: FromArray([]T{})},
		{name: "6. delete sub", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{3})},
		{name: "7. delete mix", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1, 3}), want: FromArray([]T{3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.DeleteFrom(tt.args.set)
			if !reflect.DeepEqual(tt.want, tt.s) {
				t.Errorf("Delete(%v) = %v, want = %v", tt.args.set, tt.s, tt.want)
			}
		})
	}
}

func TestSet_Has(t *testing.T) {
	type T int
	type args struct {
		v T
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want bool
	}{
		{name: "1. test common - exists", args: args{v: 1}, s: FromArray([]T{1, 2, 3}), want: true},
		{name: "2. test common - not exists", args: args{v: 4}, s: FromArray([]T{1, 2, 3}), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Has(tt.args.v); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_MergeFrom(t *testing.T) {
	type T int
	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want Set[T]
	}{
		{name: "1. test all empty value", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: FromArray([]T{})},
		{name: "2. merge from empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{1, 2, 3})},
		{name: "3. merge to empty", args: args{set: FromArray([]T{1, 2, 3})}, s: FromArray([]T{}), want: FromArray([]T{1, 2, 3})},
		{name: "4. merge exists", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{1, 2, 3})},
		{name: "5. merge not exists", args: args{set: FromArray([]T{4, 5})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{1, 2, 3, 4, 5})},
		{name: "6. merge mix", args: args{set: FromArray([]T{1, 2, 4, 5})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{1, 2, 3, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.MergeFrom(tt.args.set)
		})

		if !reflect.DeepEqual(tt.want, tt.s) {
			t.Errorf("MergeFrom(%v) = %v, want = %v", tt.args.set, tt.s, tt.want)
		}
	}
}

func TestSet_PushValue(t *testing.T) {
	type T int
	type args struct {
		v T
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want Set[T]
	}{
		{name: "1. test push not exists", args: args{v: 1}, s: FromArray([]T{}), want: FromArray([]T{1})},
		{name: "2. test push exists", args: args{v: 1}, s: FromArray([]T{1}), want: FromArray([]T{1})},
		{name: "3. test push to empty", args: args{v: 1}, s: FromArray([]T{}), want: FromArray([]T{1})},
		{name: "4. test push common", args: args{v: 1}, s: FromArray([]T{2, 3}), want: FromArray([]T{1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.v)
			if !reflect.DeepEqual(tt.want, tt.s) {
				t.Errorf("Push(%v) = %v, want = %v", tt.args.v, tt.s, tt.want)
			}
		})
	}
}

func TestFromArray(t *testing.T) {
	type T int
	type args struct {
		array []T
	}
	tests := []struct {
		name string
		args args
		want Set[T]
	}{
		{name: "1. common case", args: args{array: []T{1, 2, 3}}, want: Set[T]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}},
		{name: "2. empty input - empty array", args: args{array: []T{}}, want: Set[T]{}},
		{name: "3. empty input - nil input", args: args{array: nil}, want: Set[T]{}},
		{name: "4. dump key", args: args{array: []T{1, 2, 1}}, want: Set[T]{1: struct{}{}, 2: struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_SafeEquals(t *testing.T) {
	type T int
	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want bool
	}{
		{name: "1. test empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: true},
		{name: "2. test same length different", args: args{set: FromArray([]T{1})}, s: FromArray([]T{2}), want: false},
		{name: "3. test different length", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1}), want: false},
		{name: "4. test same length with many element", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{3, 4}), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SafeEquals(tt.args.set); got != tt.want {
				t.Errorf("SafeEquals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Equals(t *testing.T) {
	type T int
	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want bool
	}{
		{name: "1. test empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: true},
		{name: "2. test same length different", args: args{set: FromArray([]T{1})}, s: FromArray([]T{2}), want: false},
		{name: "3. test different length", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1}), want: false},
		{name: "4. test same length with many element", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{3, 4}), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equals(tt.args.set); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSupperOf(t *testing.T) {

	type T int

	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want bool
	}{
		{name: "1. test nil of nil", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: true},
		{name: "2. test nil of set", args: args{set: FromArray([]T{})}, s: FromArray([]T{1, 2, 3}), want: true},
		{name: "3. test equals", args: args{set: FromArray([]T{1, 2, 3})}, s: FromArray([]T{1, 2, 3}), want: true},
		{name: "4. test sub of set", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1, 2, 3}), want: true},
		{name: "5. test supper of set", args: args{set: FromArray([]T{1, 2, 3, 4})}, s: FromArray([]T{1, 2, 3}), want: false},
		{name: "6. test different set", args: args{set: FromArray([]T{1, 2, 3, 4})}, s: FromArray([]T{6}), want: false},
		{name: "7. test some to nil", args: args{set: FromArray([]T{1, 2, 3, 4})}, s: FromArray([]T{}), want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSupperOf(tt.args.set); got != tt.want {
				t.Errorf("IsSupperOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsSubOf(t *testing.T) {
	type T int

	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want bool
	}{
		{name: "1. test nil of nil", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: true},
		{name: "2. test nil of set", args: args{set: FromArray([]T{})}, s: FromArray([]T{1, 2, 3}), want: false},
		{name: "3. test equals", args: args{set: FromArray([]T{1, 2, 3})}, s: FromArray([]T{1, 2, 3}), want: true},
		{name: "4. test sub of set", args: args{set: FromArray([]T{1, 2})}, s: FromArray([]T{1, 2, 3}), want: false},
		{name: "5. test supper of set", args: args{set: FromArray([]T{1, 2, 3, 4})}, s: FromArray([]T{1, 2, 3}), want: true},
		{name: "6. test different set", args: args{set: FromArray([]T{1, 2, 3, 4})}, s: FromArray([]T{6}), want: false},
		{name: "7. test some to nil", args: args{set: FromArray([]T{1, 2, 3, 4})}, s: FromArray([]T{}), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSubOf(tt.args.set); got != tt.want {
				t.Errorf("IsSubOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_GetSub(t *testing.T) {
	type T int
	type args struct {
		set Set[T]
	}
	tests := []struct {
		name string
		s    Set[T]
		args args
		want Set[T]
	}{
		{name: "1. test empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{}), want: FromArray([]T{})},
		{name: "2. test s empty", args: args{set: FromArray([]T{1, 2, 3})}, s: FromArray([]T{}), want: FromArray([]T{})},
		{name: "3. test set empty", args: args{set: FromArray([]T{})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{})},
		{name: "4. test (a contain b) a sub from b", args: args{set: FromArray([]T{1})}, s: FromArray([]T{1, 2, 3}), want: FromArray([]T{1})},
		{name: "5. test (a contain b) b sub from a", args: args{set: FromArray([]T{1, 2, 3})}, s: FromArray([]T{1}), want: FromArray([]T{1})},
		{name: "6. test sub nil", args: args{set: nil}, s: FromArray([]T{1}), want: FromArray([]T{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetMix(tt.args.set); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMix(t *testing.T) {

	type T int

	type args struct {
		sets []Set[T]
	}
	tests := []struct {
		name string
		args args
		want Set[T]
	}{
		{
			name: "1. test common case 1",
			args: args{
				sets: []Set[T]{
					FromArray([]T{}),
					FromArray([]T{}),
					FromArray([]T{}),
				},
			},
			want: FromArray([]T{}),
		},
		{
			name: "2. test common case 2",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1}),
					FromArray([]T{2}),
					FromArray([]T{3}),
				},
			},
			want: FromArray([]T{}),
		},
		{
			name: "3. test common case 3",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1, 2, 3}),
					FromArray([]T{1, 2}),
					FromArray([]T{1}),
				},
			},
			want: FromArray([]T{1}),
		},
		{
			name: "4. test common case 4",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1, 2}),
					FromArray([]T{2, 3}),
					FromArray([]T{1, 3}),
				},
			},
			want: FromArray([]T{}),
		},
		{
			name: "5. test common case 5",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1, 2}),
					nil,
					FromArray([]T{1, 3}),
				},
			},
			want: FromArray([]T{1}),
		},
		{
			name: "6. test mix of nil",
			args: args{
				sets: nil,
			},
			want: FromArray([]T{}),
		},
		{
			name: "6. test mix of nil",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1}),
				},
			},
			want: FromArray([]T{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMix(tt.args.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMerge(t *testing.T) {
	type T int
	type args struct {
		sets []Set[T]
	}
	tests := []struct {
		name string
		args args
		want Set[T]
	}{
		{
			name: "1. test nil",
			args: args{
				sets: nil,
			},
			want: FromArray([]T{}),
		},
		{
			name: "2. test one element",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1, 2, 3}),
				},
			},
			want: FromArray([]T{1, 2, 3}),
		},
		{
			name: "3. test one nil element",
			args: args{
				sets: []Set[T]{
					nil,
				},
			},
			want: FromArray([]T{}),
		},
		{
			name: "4. test common case 1",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1, 2, 3}),
					FromArray([]T{1, 2, 3}),
				},
			},
			want: FromArray([]T{1, 2, 3}),
		},
		{
			name: "5. test common case 2",
			args: args{
				sets: []Set[T]{
					FromArray([]T{1, 2, 3}),
					FromArray([]T{3, 4, 5}),
				},
			},
			want: FromArray([]T{1, 2, 3, 4, 5}),
		},
		{
			name: "5. test common case 2",
			args: args{
				sets: []Set[T]{
					nil,
					FromArray([]T{3, 4, 5}),
				},
			},
			want: FromArray([]T{3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMerge(tt.args.sets...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromMapKey(t *testing.T) {
	type T int
	type v int
	type args struct {
		ma map[T]v
	}
	tests := []struct {
		name string
		args args
		want Set[T]
	}{
		{name: "1. test common", args: args{ma: map[T]v{1: 1}}, want: FromArray([]T{1})},
		{name: "2. test nil", args: args{ma: nil}, want: FromArray([]T{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromMapKey(tt.args.ma); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromMapKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ToArray(t *testing.T) {
	type T int
	tests := []struct {
		name string
		s    Set[T]
		want []T
	}{
		{name: "1. test common", s: FromArray([]T{1, 2, 3}), want: []T{1, 2, 3}},
		{name: "2. test nil", s: FromArray([]T{}), want: []T{}},
		{name: "3. test nil", s: FromArray([]T{1, 2, 3, 1}), want: []T{1, 2, 3}},
		{name: "4. test nil", s: nil, want: []T{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToArray()
			aSet := FromArray(got)
			bSet := FromArray(tt.want)
			if !aSet.Equals(bSet) {
				t.Errorf("ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
