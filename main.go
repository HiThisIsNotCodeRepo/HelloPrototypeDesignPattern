package main

import "fmt"

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
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
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("\t")
	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("\t")
	fmt.Println("\nPrinting hierarchy for Folder3")
	folder3.print("\t")
	cloneFolder = folder3.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder3")
	cloneFolder.print("\t")
}
