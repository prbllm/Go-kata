package actionpurpose

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/prbllm/go-kata/internal/task"
)

type Task struct{}

func (Task) Name() string { return "actionpurpose" }

func (Task) Run() error {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	parser := NewActionPurposeParser()
	var testsCount, rowCount int
	fmt.Fscan(in, &testsCount)
	in.ReadString('\n')

	result := make([]string, 0)

	for i := 0; i < testsCount; i++ {
		fmt.Fscan(in, &rowCount)
		in.ReadString('\n')

		for j := 0; j < rowCount; j++ {
			line, _ := in.ReadString('\n')
			line = strings.TrimRight(line, "\r\n")
			err := parser.ParseLine(line)

			if err != nil {
				fmt.Fprintln(os.Stderr, "parse:", err)
				j++
				for ; j < rowCount; j++ {
					in.ReadString('\n')
				}
				break
			}
		}

		results := parser.GetResult()
		for _, resultLine := range results {
			result = append(result, resultLine+" is "+parser.lastAction+".")
		}
		parser.CleanData()
	}

	for _, resultLine := range result {
		fmt.Fprintln(out, resultLine)
	}
	return nil
}

func init() {
	task.Register(Task{})
}
