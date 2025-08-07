package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/prbllm/go-kata/internal/actionpurpose"
	_ "github.com/prbllm/go-kata/internal/route"
	"github.com/prbllm/go-kata/internal/task"
)

func main() {
	taskFlag := flag.String("task", "", "Имя задачи (см. --list)")
	listFlag := flag.Bool("list", false, "Показать все доступные задачи")
	flag.Parse()

	if *listFlag {
		fmt.Println("Доступные задачи:")
		for _, name := range task.All() {
			fmt.Printf("  %s\n", name)
		}
		return
	}

	if *taskFlag == "" {
		fmt.Fprintln(os.Stderr, "нужен --task=<name> (или --list)")
		os.Exit(1)
	}

	runner, ok := task.Get(*taskFlag)
	if !ok {
		fmt.Fprintf(os.Stderr, "задача %q не найдена\n", *taskFlag)
		os.Exit(1)
	}

	if err := runner.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}
