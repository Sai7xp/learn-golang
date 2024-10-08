## Isolating Go Slices: How to create separate slices from an array safely

```go
arr := [...]int{1, 2, 3, 4, 5}
fmt.Println(len(arr), arr) // 5 [1 2 3 4 5]

// slicing the above array
s1 := arr[2:5] // index 2 (inclusive) to index 5 (exclusive)
fmt.Println(len(s1), cap(s1), s1) // 3 3 [3 4 5]

s2 := arr[1:4]
fmt.Println(len(s2), cap(s2), s2) // 3 4 [2 3 4]

s1[0] = 99 // this will modify both the slices and array as well

fmt.Println("arr:", len(arr), arr) // arr: 5 [1 2 99 4 5] 🤯
fmt.Println("s1:", len(s1), cap(s1), s1) // s1: 3 3 [99 4 5]
fmt.Println("s2:", len(s2), cap(s2), s2) // s2: 3 4 [2 99 4] 🤯
```

> Capacity of a slice will be calculated from starting point of slice (index 2 for `s1`) to the end of the underlying array. So capacity will be 3 for s1. and 4 for s2 (starting from 1 till the end of the array)

**Modifying `s1`**

```go
    s1[0] = 99
```

- Since both `s1` and `s2` pointing to the same underlying array, modifying s1 affects arr and s2 slice as well.
- Slice is a view of an underlying array. When we create a slice from an array, data will not be copied. It will just create a reference to array. So modifying the array via slice will be visible to all other slices that share the same underlying array.

## Then how can we create separate slices using `copy` func

```go
arr := [...]int{1, 2, 3, 4, 5}

s1 := make([]int, 3)
copy(s1, arr[2:5]) // s1 is now a copy of arr[2:5]

s2 := make([]int, 3)
copy(s2, arr[1:4])

s1[0] = 99

fmt.Println("Original array:", arr) // [1 2 3 4 5] 😌 array is safe now
fmt.Println("s1:", s1)              // [99 4 5]
fmt.Println("s2:", s2)              // [2 3 4]
```

- When slice is created using `make` with specified size, it will be initialized with zero values.
- `copy` function will copy data from array into the slice. So new slice will have their own independent memory and not linked with the original array anymore.
