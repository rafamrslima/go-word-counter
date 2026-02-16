package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
)

type wordCount struct {
	word  string
	count int
}

func countWords(filename string) (map[string]int, int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	counts := make(map[string]int)
	total := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		counts[word]++
		total++
	}
	return counts, total, scanner.Err()
}

func main() {
	top := flag.Int("top", 10, "number of top words to display")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "usage: wcount -top N file1.txt file2.txt ...")
		os.Exit(1)
	}

	type result struct {
		counts map[string]int
		total  int
		err    error
		file   string
	}

	results := make(chan result, len(files))
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			counts, total, err := countWords(f)
			results <- result{counts, total, err, f}
		}(file)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	merged := make(map[string]int)
	totalWords := 0

	for r := range results {
		if r.err != nil {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", r.file, r.err)
			continue
		}
		totalWords += r.total
		for w, c := range r.counts {
			merged[w] += c
		}
	}

	sorted := make([]wordCount, 0, len(merged))
	for w, c := range merged {
		sorted = append(sorted, wordCount{w, c})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].count > sorted[j].count
	})

	n := *top
	if n > len(sorted) {
		n = len(sorted)
	}

	fmt.Printf("total_files: %d\n", len(files))
	fmt.Printf("total_words: %d\n", totalWords)
	fmt.Printf("unique_words: %d\n", len(merged))
	fmt.Printf("\ntop %d:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s  %d\n", i+1, sorted[i].word, sorted[i].count)
	}
}
