package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([][]string, 12)

	// Create skuces to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	} // we count how many words are in each bucket

	// Print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, "\t", len(buckets[i]))
	}
	// Print the words in one of the buckets
	// fmt.Println(buckets[6])
}

// HashBucket :
func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets

	// return len(word) % buckets
}
