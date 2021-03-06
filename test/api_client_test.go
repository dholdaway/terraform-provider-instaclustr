package test

import (
    "fmt"
	"io/ioutil"
	"testing"

	"github.com/instaclustr/terraform-provider-instaclustr/instaclustr"
)

func TestAPIClientRead(t *testing.T) {
    id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
    client := SetupMock(t, id, fmt.Sprintf(`{"id":"%s"}`, id), 202)
	cluster, err := client.ReadCluster(id)
	if err != nil {
		t.Fatalf("Failed to read cluster %s: %s", id, err)
	}
	if cluster.ID != id {
		t.Fatalf("Cluster expected %s but got %s", id, cluster.ID)
	}
}

func TestAPIClientReadNull(t *testing.T) {
	id := "Invalid_ID"
	client := SetupMock(t, id, "", 404)
	var _, err = client.ReadCluster(id)
	if err == nil {
		t.Fatalf("Read a cluster expected error but got null")
	}
}

func TestAPIClientDelete(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	client := SetupMock(t, id, "", 202)
	var err = client.DeleteCluster(id)
	if err != nil {
		t.Fatalf("Failed to delete cluster %s: %s", id, err)
	}
}

func TestAPIClientDeleteNull(t *testing.T) {
	id := "Invalid_ID"
	client := SetupMock(t, id, "", 404)
	var err = client.DeleteCluster(id)
	if err == nil {
		t.Fatalf("Delete a cluster expected error but got null")
	}
}

func TestAPIClientCreate(t *testing.T) {
	filename := "data/valid_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "extended/", `{"id":"should-be-uuid"}`, 202)
	id, err := client.CreateCluster(jsonStr)
	if err != nil {
		t.Fatalf("Failed to create cluster: %s", err)
	}
	if id == "" {
		t.Fatalf("Failed to fetch cluster id")
	}
}

func TestAPIClientCreateInvalid(t *testing.T) {
	filename := "data/invalid_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}
	client := SetupMock(t, "extended/", ``, 401)
	_, err = client.CreateCluster(jsonStr)
	if err == nil {
		t.Fatalf("Create a cluster expected error but got null")
	}
}

func TestAPIClientCreateNoNetwork(t *testing.T) {
	filename := "data/valid_create.json"
	jsonStr, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("Failed to load %s: %s", filename, err)
	}

	client := new(instaclustr.APIClient)
	client.InitClient("http://127.0.0.1:5000", "Trisolaris", "DoNotAnswer!")
	_, err = client.CreateCluster(jsonStr)
	if err == nil {
		t.Fatalf("Create a cluster expected error but got null")
	}
}

func TestAPIClientReadNoNetwork(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	client := new(instaclustr.APIClient)
	client.InitClient("http://127.0.0.1:5000", "Trisolaris", "DoNotAnswer!")
	_, err := client.ReadCluster(id)
	if err == nil {
		t.Fatalf("Read a cluster expected error but got null")
	}
}

func TestAPIClientDeleteNoNetwork(t *testing.T) {
	id := "77b5a4e1-c422-4a78-b551-d8fa5c42ad95"
	client := new(instaclustr.APIClient)
	client.InitClient("http://127.0.0.1:5000", "Trisolaris", "DoNotAnswer!")
	err := client.DeleteCluster(id)
	if err == nil {
		t.Fatalf("Delete a cluster expected error but got null")
	}
}
