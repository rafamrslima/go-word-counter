# Word Count Tool (wcount)

A fast, concurrent word counter written in Go that processes multiple text files and displays the most frequently occurring words.

## Features

- 🚀 **Concurrent Processing**: Uses goroutines to process multiple files in parallel
- 📊 **Top Words Ranking**: Displays the N most frequent words across all files
- 🔤 **Case Insensitive**: Treats "Go", "GO", and "go" as the same word
- 📈 **Statistics**: Shows total files processed, total words, and unique words
- ⚡ **Fast and Efficient**: Built with Go's powerful concurrency primitives

## Installation

### Prerequisites
- Go 1.25.0 or higher

## Usage

### Basic Usage

```bash
go run main.go file1.txt file2.txt file3.txt
```

Or using the compiled binary:

```bash
./wcount file1.txt file2.txt file3.txt
```

### Display Top N Words

Use the `-top` flag to specify how many top words to display (default: 10):

```bash
go run main.go -top 5 sample1.txt sample2.txt
```

```bash
./wcount -top 20 *.txt
```

## Output Example

```
total_files: 2
total_words: 150
unique_words: 75

top 10:
1. the  15
2. and  12
3. go  8
4. is  7
5. to  6
6. a  5
7. of  5
8. in  4
9. for  4
10. with  3
```


## Testing

Run the test suite:

```bash
go test
```

Run tests with verbose output:

```bash
go test -v
```

## Performance

The tool uses Go's concurrent programming features to process multiple files simultaneously, making it efficient for processing large numbers of files. Each file is read and processed in its own goroutine, and results are aggregated using channels.

## Error Handling

The tool gracefully handles errors:
- Missing files are reported to stderr and processing continues with remaining files
- File read errors don't stop the processing of other files
- Usage information is displayed when no files are provided