package depsinj

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

type GitLabClient interface {
	GetProject(projectID interface{}, opt *gitlab.GetProjectOptions, option ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error)
	//Other projects
}
type ReleaseServiceGitLabClient struct {
	Client *gitlab.Client
}

func (client *ReleaseServiceGitLabClient) GetProject(projectID interface{}, opt *gitlab.GetProjectOptions, opts ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error) {
	fmt.Println("Calling actual project API...")
	return client.Client.Projects.GetProject(projectID, opt, opts...)
}

type MockGitLabClient struct {

}

func (client *MockGitLabClient) GetProject(projectID interface{}, opt *gitlab.GetProjectOptions, opts ...gitlab.RequestOptionFunc) (*gitlab.Project, *gitlab.Response, error) {
	return &gitlab.Project{Name: "X Project"}, nil, nil
}




type UserService struct {
	logger Logger
}

type ReleaseService struct {
	Client GitLabClient
}

func NewReleaseService(client GitLabClient) *ReleaseService {
	fmt.Println("Creating Relase service with ", client)
	return &ReleaseService{Client: client}
}

func NewUserService(logger Logger) *UserService {
	fmt.Println("Creating new UserService")
	return &UserService{logger: logger}
}

func (service *UserService) DoSomething() {
	service.logger.Log("This is my logging in...in User service DoSomething() method")
}