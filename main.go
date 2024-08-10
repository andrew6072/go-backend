package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Employee struct to hold individual employee data
type Employee struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// Response struct to hold the API response
type Response struct {
	Status string     `json:"status"`
	Data   []Employee `json:"data"`
}

type Result struct {
	ID     int
	Result float64
}

func worker(w int, jobs <-chan *Employee, results chan<- *Result) {
	for j := range jobs {
		if j.EmployeeAge == 0 {
			fmt.Printf("Worker %d skipped employee %d due to zero age\n", w, j.ID)
			continue
		}
		res := float64(j.EmployeeSalary) / float64(j.EmployeeAge)
		fmt.Printf("Worker %d processed employee %d\n", w, j.ID)
		results <- &Result{ID: j.ID, Result: res}
	}
}

func main() {
	resp, err := http.Get("https://dummy.restapiexample.com/api/v1/employees") // returns (*http.Response, error)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get a valid response: %v", resp.Status)
	}

	var body []byte
	buffer := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatalf("Failed to read response body: %v", err)
		}
		if n == 0 {
			break
		}
		body = append(body, buffer[:n]...)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// fmt.Printf("%s", body)

	fmt.Printf("Status: %s\n", response.Status)

	numJobs := len(response.Data)
	jobs := make(chan *Employee, numJobs)
	results := make(chan *Result, numJobs)
	const numWorkers = 3

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for _, employee := range response.Data {
		emp := employee
		jobs <- &emp
	}
	close(jobs)

	//Wait for all workers to finish
	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Printf("Employee ID: %d, Salary/Age: %.2f\n", res.ID, res.Result)
	}
	close(results)
}
