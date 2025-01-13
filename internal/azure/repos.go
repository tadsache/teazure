package azure

import (
	"context"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7"
	"github.com/microsoft/azure-devops-go-api/azuredevops/v7/git"
	"log"
	"os"
)

// GetReposForProject retrieves repositories for a given project Id.
// func GetReposForProject(projectId string) []git.GitRepository {
func GetReposForProject(projectId string) *[]git.GitRepository {
	organizationUrl := os.Getenv("organizationUrl")
	personalAccessToken := os.Getenv("personalAccessToken")

	connection := azuredevops.NewPatConnection(organizationUrl, personalAccessToken)
	ctx := context.Background()

	// Create a client for the Git area
	gitClient, err := git.NewClient(ctx, connection)
	if err != nil {
		log.Fatalf("Failed to create Git client: %v", err)
	}

	// Fetch repositories for the specified project
	repos, err := gitClient.GetRepositories(ctx, git.GetRepositoriesArgs{
		Project: &projectId,
	})
	if err != nil {
		log.Fatalf("Failed to get repositories for project %s: %v", projectId, err)
	}

	return repos
}
