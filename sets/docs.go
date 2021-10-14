/*
Package sets implement sets data struct for generics world
```go

	setStr := sets.NewWithCap[string](10)
	setStr.Add("my")
	art.True(setStr.Has("my"))
	art.Equal(1, setStr.Length())

	setStr.Add("hello")
	art.Equal(2, setStr.Length())
	setStr.Pop("hello")
	art.Equal(1, setStr.Length())
```

*/
package sets
