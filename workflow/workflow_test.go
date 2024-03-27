package workflow

import (
	"go.temporal.io/sdk/temporal"
	"simple-workflow/app/activity"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_GenerateEvenNumber(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	generatedNumber := 2
	env.OnActivity(activity.GenerateRandomNumber, mock.Anything).Return(generatedNumber, nil)
	env.OnActivity(activity.PrintEvenNumber, mock.Anything, generatedNumber).Return(nil)
	env.OnActivity(activity.PrintOddNumber, mock.Anything, generatedNumber).Return(nil)
	env.ExecuteWorkflow(SimpleWorkflow)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	env.AssertCalled(t, "PrintEvenNumber", mock.Anything, generatedNumber)
	env.AssertNotCalled(t, "PrintOddNumber", mock.Anything, mock.Anything)
}

func Test_GenerateOddNumber(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	generatedNumber := 1
	env.OnActivity(activity.GenerateRandomNumber, mock.Anything).Return(generatedNumber, nil)
	env.OnActivity(activity.PrintEvenNumber, mock.Anything, generatedNumber).Return(nil)
	env.OnActivity(activity.PrintOddNumber, mock.Anything, generatedNumber).Return(nil)
	env.ExecuteWorkflow(SimpleWorkflow)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	env.AssertCalled(t, "PrintOddNumber", mock.Anything, generatedNumber)
	env.AssertNotCalled(t, "PrintEvenNumber", mock.Anything, mock.Anything)
}

// Activity should retry on retryable error
func Test_ActivityRetryMaxiumAttempts(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	generatedNumber := 2
	env.OnActivity(activity.GenerateRandomNumber, mock.Anything).Return(generatedNumber, "some error")
	env.ExecuteWorkflow(SimpleWorkflow)
	require.True(t, env.IsWorkflowCompleted())
	require.ErrorContains(t, env.GetWorkflowError(), "some error")
	env.AssertNumberOfCalls(t, "GenerateRandomNumber", 3)
	env.AssertNotCalled(t, "PrintEvenNumber", mock.Anything, mock.Anything)
	env.AssertNotCalled(t, "PrintOddNumber", mock.Anything, mock.Anything)
}

// Workflow should succeed if retry succeed
func Test_ActivitySucceedAfterRetry(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	generatedNumber := 2
	env.OnActivity(activity.GenerateRandomNumber, mock.Anything).Panic("Retryable").Once()
	env.OnActivity(activity.GenerateRandomNumber, mock.Anything).Return(generatedNumber, nil)
	env.OnActivity(activity.PrintEvenNumber, mock.Anything, generatedNumber).Return(nil)
	env.ExecuteWorkflow(SimpleWorkflow)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	env.AssertNumberOfCalls(t, "GenerateRandomNumber", 2)
}

// Activity should not retry on non retryable error
func Test_ActivityNonRetryableError(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	env.OnActivity(activity.GenerateRandomNumber, mock.Anything).Return(2, temporal.NewApplicationError("NonRetryableError", "NonRetryableError"))
	env.ExecuteWorkflow(SimpleWorkflow)
	require.True(t, env.IsWorkflowCompleted())
	require.ErrorContains(t, env.GetWorkflowError(), "NonRetryableError")
	env.AssertNumberOfCalls(t, "GenerateRandomNumber", 1)
	env.AssertNotCalled(t, "PrintEvenNumber", mock.Anything, mock.Anything)
	env.AssertNotCalled(t, "PrintOddNumber", mock.Anything, mock.Anything)
}
