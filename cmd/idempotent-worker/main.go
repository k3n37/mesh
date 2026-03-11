package main

import "fmt"

type Job struct {
	ID      string
	Payload string
}

func main() {
	processed := map[string]bool{}
	jobs := []Job{
		{ID: "job-1", Payload: "send-email"},
		{ID: "job-1", Payload: "send-email"},
		{ID: "job-2", Payload: "refresh-cache"},
	}

	for _, job := range jobs {
		if processed[job.ID] {
			fmt.Printf("skipping duplicate job %s\n", job.ID)
			continue
		}

		processed[job.ID] = true
		fmt.Printf("processing job %s -> %s\n", job.ID, job.Payload)
	}
}
