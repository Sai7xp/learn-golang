## nil slice vs nil map: Why nil Slices accept new values, but nil Maps don't

```go
var slice []int // nil slice
slice = append(slice, 1) // append allocates new underlying array, so it works
fmt.Println(slice) // [1]

var m map[string]int // nil map
m["one"] = 1 // 🚨 Runtime panic: assignment to entry in nil map
```

### Appending to a Nil Slice:

**Why it works:**

- Slice consists of 3 components
  1.  Pointer to underlying array
  2.  Length of Slice
  3.  Capacity of Slice

When we declare a slice as `nil`, it doesn't point to any underlying array and both len and cap are zero. However `append` still works.

The `append` function checks the capacity and if it is insufficient or if the slice is `nil`, it allocates a new underlying array with enough capacity to hold the new elements, and then returns a new slice that points to this array. The original `nil` slice is not mutated; instead, a new slice is created and returned.

### Adding to a Nil Map:

**Why it doesn't work:**

- **Maps** are more complex then slices. Map is basically a hash table.
- Before we can store any key-value pairs, internal data structures needs to be initialized first.

When we declare the `nil` map (`var m map[string]int`), map is not initialized and it doesn't have any underlying hash table to store key-value pairs.

So if we try to add a key-value pair to a nil map, it will create a runtime panic. Maps in go needs to be properly initialized first before we use them.

```go
var m map[string]int // nil map

m = make(map[string]int)
m["one"] = 1 // map is now initialized, so we can add values
print(m) // map[one:1]
```
