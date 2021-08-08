# HelloPrototypePattern

> [Source](https://golangbyexample.com/prototype-pattern-go/)

## Prerequisite

The object needs to implement `inode` to enable clone, the implementation itself is clone operation.

```
type inode interface {
	print(string)
	clone() inode
}
```

## How `clone()` works

It works as recursive , it will exhaust all folders and clone each file.

```
func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
	    // recursive clone
		aCopy := i.clone()
		tempChildren = append(tempChildren, aCopy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

```

## Flow of `folder3.clone()`
```
folder1 := &folder{
    children: []inode{file1},
    name:     "Folder1",
}
folder2 := &folder{
    children: []inode{folder1, file2, file3},
    name:     "Folder2",
}
folder3 := &folder{
    children: []inode{folder2},
    name:     "Folder3",
}
folder3.clone()

# First call 
1. folder3 change name -> "Folder3_clone"
2. folder2.clone()
# In folder2.clone()
1. folder2 change name -> "Folder2_clone"
2. folder1.clone()
# In folder1.clone()
1. folder1 change name -> "Folder1_clone"
2. file1.clone() -> "File1_clone"
# Back to folder2.clone()
1. file2.clone() -> "File2_clone"
2. file3.clone() -> "File3_clone"
```
