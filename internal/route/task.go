package route

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/prbllm/go-kata/internal/task"
)

type Task struct{}

func (Task) Name() string {
	return "route"
}

func (Task) Run() error {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testsCount, rowCount, colCount int
	fmt.Fscan(in, &testsCount)
	in.ReadString('\n')

	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &rowCount)
		in.ReadString('\n')

		fmt.Fscan(in, &colCount)
		in.ReadString('\n')

		navigator := NewNavigator(rowCount, colCount)

		for j := 0; j < rowCount; j++ {
			line, _ := in.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			err := navigator.ParseLine(line, j)

			if err != nil {
				fmt.Fprintln(os.Stderr, "parse:", err)
				j++
				for ; j < rowCount; j++ {
					in.ReadString('\n')
				}
				break
			}
		}

		results := navigator.GetResult()
		for _, resultLine := range results {
			fmt.Fprintln(out, resultLine)
		}
	}
	return nil
}

func init() {
	task.Register(Task{})
}
