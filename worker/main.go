package main

import (
	"log"
	"simple-workflow/app/activity"
	"simple-workflow/app/common"
	"simple-workflow/app/workflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client.", err)
	}
	defer c.Close()

	w := worker.New(c, common.SimpleWorkflowTaskQueueName, worker.Options{})

	// This worker hosts both Workflow and Activity functions.
	w.RegisterWorkflow(workflow.SimpleWorkflow)
	w.RegisterActivity(activity.GenerateRandomNumber)
	w.RegisterActivity(activity.PrintEvenNumber)
	w.RegisterActivity(activity.PrintOddNumber)

	// Start listening to the Task Queue.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to cmd Worker", err)
	}
}
