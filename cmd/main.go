package cmd

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"log"
	"simple-workflow/app/common"
	"simple-workflow/app/workflow"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}

	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        fmt.Sprintf("simple-workflow-%s", uuid.New()),
		TaskQueue: common.SimpleWorkflowTaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, workflow.SimpleWorkflow)
	if err != nil {
		log.Fatalln("Unable to cmd the Workflow:", err)
	}

	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
}
