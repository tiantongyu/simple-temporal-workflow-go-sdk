package workflow

import (
	"log"
	"simple-workflow/app/activity"
	"simple-workflow/app/common"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func SimpleWorkflow(ctx workflow.Context) error {
	ctx = initContext(ctx)

	var randomNumber int
	err := workflow.ExecuteActivity(ctx, activity.GenerateRandomNumber).Get(ctx, &randomNumber)
	if err != nil {
		log.Printf("Activity failed due to error %s", err)
		return err
	}
	if randomNumber%2 == 0 {
		err := workflow.ExecuteActivity(ctx, activity.PrintEvenNumber, randomNumber).Get(ctx, nil)
		if err != nil {
			return err
		}
	} else {
		err := workflow.ExecuteActivity(ctx, activity.PrintOddNumber, randomNumber).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// No error
	return nil
}

func initContext(ctx workflow.Context) workflow.Context {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     common.ExponentialBackoffCoefficient,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        common.MaxiumRetryAttempts,
		NonRetryableErrorTypes: []string{"NonRetryableError"},
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	return ctx
}
