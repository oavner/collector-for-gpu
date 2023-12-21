package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {

	// Parse command-line arguments
	var namespace string
	flag.StringVar(&namespace, "namespace", "default", "Kubernetes namespace")
	flag.Parse()

	// Connect to the containerd daemon
	client, err := containerd.New("/run/containerd/containerd.sock")
	if err != nil {
		log.Fatal("Error connecting to containerd:", err)
	}
	defer client.Close()

	// Specify the namespace (default in this case)
	// namespace := "default"

	// Get all containers in the specified namespace
	containerSummaries, err := getContainers(client, namespace)
	if err != nil {
		log.Fatal("Error getting containers:", err)
	}

	fmt.Println(containerSummaries)
	// Print information about each container
	//for _, summary := range containerSummaries {
	//	fmt.Printf("Container ID: %s\n", summary.ID)
	//	fmt.Printf("Container Pid: %d\n", summary.Pid)
	//	fmt.Printf("Container Status: %s\n", summary.Status)
	//	fmt.Println("-------------------------------")
	//}
}

func getContainers(client *containerd.Client, namespace string) ([]containerd.Container, error) {
	// List all containers in the specified namespace
	ctx := namespaces.WithNamespace(context.Background(), namespace)
	containers, err := client.Containers(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %v", err)
	}

	// Collect container summaries
	var containerSummaries []containerd.Container

	// Iterate through the containers and get summary information
	for _, c := range containers {
		containerSummaries = append(containerSummaries, c)
	}

	return containerSummaries, nil
}
