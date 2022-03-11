package nics

import (
	"github.com/chnsz/golangsdk"
)

// POST /v1/{project_id}/cloudservers/{server_id}/nics
func createURL(client *golangsdk.ServiceClient, serverId string) string {
	return client.ServiceURL("cloudservers", serverId, "nics")
}

// POST /v1/{project_id}/cloudservers/{server_id}/nics/delete
func deleteURL(client *golangsdk.ServiceClient, serverId string) string {
	return client.ServiceURL("cloudservers", serverId, "nics/delete")
}

// GET /v1/{project_id}/cloudservers/{server_id}/os-interface
func getURL(client *golangsdk.ServiceClient, serverId string) string {
	return client.ServiceURL("cloudservers", serverId, "os-interface")
}
