package main

import (
	order "go-saga-workflow/internal"
	wf "go-saga-workflow/pkg"
	"log"
)

func main() {
	workflow := wf.NewWorkflow()
	order.ProcessOrder(workflow)

	err := workflow.Execute()

	if err != nil {
		log.Printf("saga execution failed: %+v\n", err)
	} else {
		log.Printf("saga executed successfully\n")
	}
}
