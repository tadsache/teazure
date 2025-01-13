package azure

import (
	"context"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/pipelines"
	"log"
	"os"
)

// GetPipelinesForProject retrieves pipelines for a given project Id
func GetPipelinesForProject(projectId string) *[]pipelines.Pipeline {
	organizationUrl := os.Getenv("organizationUrl")
	personalAccessToken := os.Getenv("personalAccessToken")

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()

	// why the f is pipelines not returning errr??
	pipelineClient := pipelines.NewClient(ctx, connection)

	args := pipelines.ListPipelinesArgs{
		Project: &projectId,
	}
	response, err := pipelineClient.ListPipelines(ctx, args)
	if err != nil {
		log.Fatalf("Failed to fetch pipelines for project %s: %v", projectId, err)
	}

	return response
}
