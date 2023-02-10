package set

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	type args[T comparable] struct {
		values []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			args: args[int]{
				values: []int{},
			},
			want: make(Set[int]),
		},
		{
			name: "single value",
			args: args[int]{
				values: []int{1},
			},
			want: Set[int]{1: struct{}{}},
		},
		{
			name: "multiple values with duplicates",
			args: args[int]{
				values: []int{1, 2, 3, 4, 3, 1},
			},
			want: Set[int]{1: struct{}{}, 2: struct{}{}, 3: struct{}{}, 4: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSet(tt.args.values...)

			assert.Equalf(t, tt.want, got, "NewSet() = %v, want %v", got, tt.want)
		})
	}
}

func TestSet_Add(t *testing.T) {
	type args[T comparable] struct {
		value []T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "add value to empty set",
			s:    NewSet[int](),
			args: args[int]{
				value: []int{1},
			},
		},
		{
			name: "add value set",
			s:    NewSet[int](2, 3, 4),
			args: args[int]{
				value: []int{1},
			},
		},
		{
			name: "add value set already in set",
			s:    NewSet[int](2, 3, 4),
			args: args[int]{
				value: []int{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			tt.s.Add(tt.args.value...)

			// then
			assert.Truef(t, tt.s.ContainsAll(tt.args.value...), "%v not in set after Add(%v)", tt.args.value, tt.args.value)
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			s:    NewSet[int](),
			args: args[int]{
				value: 1,
			},
			want: false,
		},
		{
			name: "set contains 1",
			s:    NewSet[int](1, 2, 3, 4, 5),
			args: args[int]{
				value: 1,
			},
			want: true,
		},
		{
			name: "set does not contains 2",
			s:    NewSet[int](1, 3, 4, 5),
			args: args[int]{
				value: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Contains(tt.args.value)
			assert.Equalf(t, tt.want, got, "Contains() = %v, want %v", got, tt.want)
		})
	}
}

func TestSet_ContainsAll(t *testing.T) {
	type args[T comparable] struct {
		values []T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			s:    NewSet[int](),
			args: args[int]{
				values: []int{1, 2, 3, 4, 5},
			},
			want: false,
		},
		{
			name: "set contains element",
			s:    NewSet[int](1, 2, 3, 4, 5),
			args: args[int]{
				values: []int{2, 3, 4},
			},
			want: true,
		},
		{
			name: "set does not contain all elements",
			s:    NewSet[int](1, 3, 4, 5),
			args: args[int]{
				values: []int{1, 3, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ContainsAll(tt.args.values...)
			assert.Equalf(t, tt.want, got, "ContainsAll() = %v, want %v", got, tt.want)
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args[T comparable] struct {
		value T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "delete none existing element from empty set",
			s:    NewSet[int](),
			args: args[int]{
				value: 1,
			},
			want: false,
		},
		{
			name: "delete existing element from set",
			s:    NewSet[int](1, 2, 3, 4, 5),
			args: args[int]{
				value: 2,
			},
			want: true,
		},
		{
			name: "delete non existing element from set",
			s:    NewSet[int](1, 3, 4, 5),
			args: args[int]{
				value: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Remove(tt.args.value)
			assert.Equalf(t, tt.want, got, "Remove() = %v, want %v", got, tt.want)

			assert.Falsef(t, tt.s.Contains(tt.args.value), "Removed value %v still in set", tt.args.value)
		})
	}
}

func TestSet_RemoveAll(t *testing.T) {
	type args[T comparable] struct {
		values []T
	}
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		args args[T]
		want Set[T]
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			s:    NewSet[int](),
			args: args[int]{
				values: []int{1, 2, 3, 4, 5},
			},
			want: NewSet[int](),
		},
		{
			name: "remove all elements",
			s:    NewSet[int](1, 2, 3, 4, 5, 6),
			args: args[int]{
				values: []int{1, 2, 3, 5},
			},
			want: NewSet[int](4, 6),
		},
		{
			name: "remove some elements",
			s:    NewSet[int](1, 2, 3, 4, 6),
			args: args[int]{
				values: []int{1, 3, 5},
			},
			want: NewSet[int](2, 4, 6),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.RemoveAll(tt.args.values...)

			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Updateset = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Slice(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			s:    NewSet[int](),
			want: []int{},
		},
		{
			name: "single element",
			s:    NewSet[int](1),
			want: []int{1},
		},
		{
			name: "multiple elements",
			s:    NewSet[int](1, 2, 3, 4, 5),
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Slice()
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestSet_Len(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    Set[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "empty set",
			s:    NewSet[int](),
			want: 0,
		},
		{
			name: "single element",
			s:    NewSet[int](1),
			want: 1,
		},
		{
			name: "multiple elements",
			s:    NewSet[int](1, 2, 3, 4, 5),
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := len(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
