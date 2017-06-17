package main

import "fmt"

func main() {
	n := HashBucket("Go", 12)
	fmt.Println(n)
}

// HashBucket : Returns the bucket number where it will be stored
func HashBucket(word string, buckets int) int {
	// letter := rune(word[0])
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
