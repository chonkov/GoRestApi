package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type Response struct {
	Output string `json:"output"`
}

func runCommand(command string) (string, error) {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	// Convert the slice of bytes to string and return int
	return string(output), nil
}

func validateCommand(command string, w http.ResponseWriter) bool {
	if strings.Contains(command, "sudo") {
		http.Error(w, "Not allowed to use this command", http.StatusForbidden)
		return false
	}
	return true
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	var command string

	if r.Method == http.MethodPost {
		// Accept the command via JSON body
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		type CommandBody struct {
			Command string `json:"command"`
		}

		var body CommandBody
		if err := decoder.Decode(&body); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		/*res :=*/
		validateCommand(body.Command, w)
		//if res {
		//	http.Error(w, "Not allowed to use this command", http.StatusForbidden)
		//	return
		//}

		command = body.Command
	} else {
		http.Error(w, "Invalid method", http.StatusBadRequest)
	}

	if command == "" {
		http.Error(w, "Missing command", http.StatusBadRequest)
		return
	}

	output, err := runCommand(command)
	if err != nil {
		http.Error(w, "Error executing command", http.StatusInternalServerError)
		return
	}

	resp := Response{
		Output: output,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func main() {
	port := "9090"
	http.HandleFunc("/api/cmd", handleCommand)

	fmt.Printf("Server started at : %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
