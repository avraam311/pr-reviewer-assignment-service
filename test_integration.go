package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)

	fmt.Println("Starting integration tests...")

	fmt.Println("\n1. Testing POST /api/v1/team/add")
	team := map[string]interface{}{
		"team_name": "test_team",
		"members": []map[string]interface{}{
			{"user_id": "user1", "username": "User One", "is_active": true},
		},
	}
	data, _ := json.Marshal(team)
	resp, err := http.Post("http://localhost:8080/api/v1/team/add", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d, Response: %s\n", resp.StatusCode, string(body))

	fmt.Println("\n2. Testing GET /api/v1/team/get/test_team")
	resp, err = http.Get("http://localhost:8080/api/v1/team/get/test_team")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d, Response: %s\n", resp.StatusCode, string(body))

	fmt.Println("\n3. Testing POST /api/v1/users/setIsActive")
	user := map[string]interface{}{
		"user_id":   "user1",
		"is_active": true,
	}
	data, _ = json.Marshal(user)
	resp, err = http.Post("http://localhost:8080/api/v1/users/setIsActive", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d, Response: %s\n", resp.StatusCode, string(body))

	fmt.Println("\n4. Testing POST /api/v1/pullRequest/create")
	pr := map[string]interface{}{
		"pull_request_id":   "pr1",
		"pull_request_name": "Test PR",
		"author_id":         "user1",
	}
	data, _ = json.Marshal(pr)
	resp, err = http.Post("http://localhost:8080/api/v1/pullRequest/create", "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d, Response: %s\n", resp.StatusCode, string(body))

	fmt.Println("\n5. Testing GET /api/v1/users/getReview/user1")
	resp, err = http.Get("http://localhost:8080/api/v1/users/getReview/user1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d, Response: %s\n", resp.StatusCode, string(body))

	fmt.Println("\n6. Testing GET /api/v1/statistics/getPRsForUser/user1")
	resp, err = http.Get("http://localhost:8080/api/v1/statistics/getPRsForUser/user1")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("Status: %d, Response: %s\n", resp.StatusCode, string(body))

	fmt.Println("\nIntegration tests completed.")
}
