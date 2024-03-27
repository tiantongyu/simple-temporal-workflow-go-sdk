# Simple Temporal Workflow In Go

This is a simple workflow written with Temporal Go SDK

## Project Structure
`/cmd`

Contains main application for starting the simple workflow

`/worker`

Contains main application for starting the worker

`/workflow`

Contains the simple workflow implementation and test

`/activity`

Contains the activity implementations for the simple workflow

`/common`

Contains shared constants and etc


## Running the application

### Step 0: Start Temporal Server
For Mac M1/M2
```bash
docker run -p 7233:7233 -p 8233:8233 tiantongdocker/temporalite:0.3.0-arm64
```
For Mac Intel and Linux
```bash
docker run -p 7233:7233 -p 8233:8233 tiantongdocker/temporalite:0.3.0
```
This will start a temporalite server listening on local port 7233 and a serving the temporal UI on port 8233

### Step 1: Start the worker

The temporal worker will register with temporal server and waiting on the simple workflow task queue

```bash
go run worker/main.go
```

### Step 2: Run the Workflow

```bash
go run cmd/main.go
```

### Step 3: Check workflow execution

You should be able to see the workflow execution logs on the worker console. 
Alternatively, you can go to http://localhost:8233 to check the workflow execution from UI
