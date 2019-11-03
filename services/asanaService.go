package services

import (
	"asanaSync/common"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// GetAsanaUserTaskByID Asana
func GetAsanaUserTaskByID() {
	url := "https://app.asana.com/api/1.0/tasks/1142421340174241"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer 0/c4bff9e1cd5dbc4be3c2154b99e04dc5")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// GetAllProjects Asana
func GetAllProjects() {
	url := "https://app.asana.com/api/1.0/projects?limit=30&workspace=1126471612412206"
	payload := strings.NewReader("opt_expand=id")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("Authorization", "Bearer 0/c4bff9e1cd5dbc4be3c2154b99e04dc5")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// GetTasksByProject Asana
func GetTasksByProject() {
	url := "https://app.asana.com/api/1.0/projects/1142153483755225/tasks?opt_expand=assignee,assignee_status,created_at,completed,due_on,due_at,name,notes&limit=10"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer 0/c4bff9e1cd5dbc4be3c2154b99e04dc5")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

// AsanaCreateTask Asana
func AsanaCreateTask(description, name, assignee, dueTo string, projectID int64) {
	if assignee == common.USER_WAN_NOT_FOND_ERROR {
		return
	}
	url := "https://app.asana.com/api/1.0/tasks"

	payload := []byte(`{"data":{"workspace":"1126471612412206", "assignee":"` +
		assignee + `","projects":["` + strconv.FormatInt(projectID, 10) + `"],"parent":"null","notes":"` +
		description + `","name":"` + name + `","due_on":"` + dueTo + `"}}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 0/c4bff9e1cd5dbc4be3c2154b99e04dc5")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	fmt.Println("task saved to asana")
}

// AsanaGetUsers Asana
func AsanaGetUsers() {
	url := "https://app.asana.com/api/1.0/users?opt_expand=name,email&workspace=1126471612412206&limit=40"
	payload := strings.NewReader("opt_expand=email")

	req, _ := http.NewRequest("GET", url, payload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer 0/c4bff9e1cd5dbc4be3c2154b99e04dc5")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
