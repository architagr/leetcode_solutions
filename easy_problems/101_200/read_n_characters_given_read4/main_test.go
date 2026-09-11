package readncharactersgivenread4

import "testing"

// read4For hands back a read4 over file, delivering up to four bytes per
// call and advancing a file pointer, exactly as the statement describes.
func read4For(file string) func([]byte) int {
	pos := 0
	return func(buf4 []byte) int {
		n := copy(buf4, file[pos:min(pos+4, len(file))])
		pos += n
		return n
	}
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestRead(t *testing.T) {
	tests := []struct {
		name string
		file string
		n    int
		want string
	}{
		{name: "example 1", file: "abc", n: 4, want: "abc"},
		{name: "example 2", file: "abcde", n: 5, want: "abcde"},
		{name: "example 3", file: "abcdABCD1234", n: 12, want: "abcdABCD1234"},
		{name: "fewer than the file holds", file: "abcdefghijk", n: 5, want: "abcde"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			read := solution(read4For(tt.file))
			buf := make([]byte, tt.n)
			got := read(buf, tt.n)
			if got != len(tt.want) {
				t.Fatalf("read() = %v, want %v", got, len(tt.want))
			}
			if string(buf[:got]) != tt.want {
				t.Errorf("buf = %q, want %q", string(buf[:got]), tt.want)
			}
		})
	}
}
