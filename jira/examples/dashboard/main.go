package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

/*
----------- Set an environment variable in git bash -----------
export HOST="https://ctreminiom.atlassian.net/"
export MAIL="MAIL_ADDRESS"
export TOKEN="TOKEN_API"

Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
*/

var (
	host  = os.Getenv("HOST")
	mail  = os.Getenv("MAIL")
	token = os.Getenv("TOKEN")
)

func getDashboardByID() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	dashboard, response, err := jiraCloud.Dashboard.Get(context.Background(), "10001")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Printf("Dashboard Name: %v", dashboard.Name)
	log.Printf("Dashboard ID: %v", dashboard.ID)
	log.Printf("Dashboard View: %v", dashboard.View)

}

func searchDashboard() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	searchOptions := jira.DashboardSearchOptionsScheme{
		DashboardName:       "Bug",
		GroupPermissionName: "administrators",
		OrderBy:             "description",
		Expand:              []string{"description", "favourite", "sharePermissions"},
	}

	dashboards, response, err := jiraCloud.Dashboard.Search(context.Background(), &searchOptions, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, dashboard := range dashboards.Values {
		log.Printf("Dashboard Name: %v", dashboard.Name)
		log.Printf("Dashboard ID: %v", dashboard.ID)
		log.Printf("Dashboard View: %v", dashboard.View)
	}

}

func createDashboard() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	var sharePermissions []jira.SharePermissionScheme

	projectPermission := &jira.SharePermissionScheme{
		Type: "project",
		Project: &jira.SharePermissionProjectScheme{
			ID: "10000",
		},
	}

	groupPermission := &jira.SharePermissionScheme{
		Type:  "group",
		Group: &jira.SharePermissionGroupScheme{Name: "jira-administrators"},
	}

	sharePermissions = append(sharePermissions, *projectPermission, *groupPermission)

	dashboard, response, err := jiraCloud.Dashboard.Create(context.Background(), "Team Tracking", "", &sharePermissions)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Printf("Dashboard Name: %v", dashboard.Name)
	log.Printf("Dashboard ID: %v", dashboard.ID)
	log.Printf("Dashboard View: %v", dashboard.View)

}

func getDashboards() (err error) {

	log.Println("------------- getDashboards -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	dashboards, response, err := atlassian.Dashboard.Gets(context.Background(), 0, 50, "")

	if err != nil {

		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}

		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(len(dashboards.Dashboards))

	return
}

func main() {

	/*
		if err := getDashboards(); err != nil {
			log.Fatal(err)
		}
	*/

	//createDashboard()
	searchDashboard()
	getDashboardByID()
}
