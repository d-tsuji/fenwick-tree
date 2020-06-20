// This library has been validated by the following online judges.
// http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=DSL_2_B
package main

import (
	"bufio"
	"fmt"
	"os"

	// We can't go mod download on the online judges,
	// so you'll need to embed the library actually.
	"github.com/d-tsuji/fenwick"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, q int
	fmt.Fscan(r, &n, &q)

	tree := fenwick.New(n)

	for i := 0; i < q; i++ {
		var com, x, y int
		fmt.Fscan(r, &com, &x, &y)

		if com == 0 {
			tree.Add(x-1, y)
		} else if com == 1 {
			fmt.Fprintln(w, tree.Query(x-1, y-1))
		}
	}
}
