package main

import (
	"fmt"
	depsinj "go-playground/deps_inj"
	"github.com/xanzy/go-gitlab"
)

type Vertex struct {
	X int
	Y int
}

// Pointer receivers
// Go does not have classes, we define methods on types
// Method is a function with special 'receiver' argument
func (r *Vertex) DoubleIt() {
	r.X *= 2
	r.Y *= 2
}

func main() {
	fmt.Println("Hello Shailendra Singh Yadav")

	//closure testing
	//f := fibbonacci()

	fmt.Println(fibbonacci()())
	//fmt.Println(fibbonacci()()) Wont work more than one time

	f := fibbonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	//struct
	v := Vertex{}
	fmt.Println(v)

	vPtr := &Vertex{}
	vPtr.X = 100
	vPtr.Y = 200
	fmt.Printf("Value of vPtr %v\n", *vPtr)

	var v1 = &Vertex{}
	v1.X = 44
	v1.Y = 55
	fmt.Println("V1 pointer", v1)
	fmt.Println("V1 value", *v1)

	var v2 = &Vertex{}
	v2.X = 10
	v2.Y = 20
	v2.DoubleIt()
	fmt.Println("After doubling it", *v2)

	//testGitLab()
	//testGitLabReleaseService()
	testGitLabReleaseServiceWithMockClient()
}

func testGitLabReleaseServiceWithMockClient() {
	client := &depsinj.MockGitLabClient{}
	//releaseServiceMockGitLabClient := &depsinj.ReleaseServiceGitLabClient{Client: client}
	releaseService := depsinj.NewReleaseService(client)
	project, _, err := releaseService.Client.GetProject(12345, nil)
	if err != nil {
		fmt.Println("Error while calling getproject", err)
	}

	fmt.Println("Got response from release service", project.Name)
}

func testGitLabReleaseService() {
	client, err := gitlab.NewClient("glpat-tysW-HXURiRkDJUSEXA-")
	if err != nil {
		fmt.Println("Error ", err)
	}
	releaseServiceGitLabClient := &depsinj.ReleaseServiceGitLabClient{Client: client}

	releaseService := depsinj.NewReleaseService(releaseServiceGitLabClient)

	project, _, err := releaseService.Client.GetProject(48275024, nil)
	if err != nil {
		fmt.Println("Error while calling getproject", err)
	}

	fmt.Println("Got response from release service", project.Name)
	
}

func testGitLab() {
	fmt.Println("Running...")
	logger := depsinj.DefaultLogger{}
	userService := depsinj.NewUserService(logger)

	userService.DoSomething()

	client, err := gitlab.NewClient("glpat-tysW-HXURiRkDJUSEXA-")
	if err != nil {
		fmt.Println("Error ", err)
	}

	users, _, err := client.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err != nil {
		fmt.Println("Error occurred while listing users", err)
	}
	for index, a := range users {
		fmt.Println(index, a.Name)
	}

}

func fibbonacci() func() int {
	last := 1
	secondLast := 0

	return func() int {
		current := last + secondLast
		secondLast = last
		last = current
		return last
	}
}
