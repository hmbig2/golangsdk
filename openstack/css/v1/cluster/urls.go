package cluster

import "github.com/chnsz/golangsdk"

// createClusterURL /v1.0/{project_id}/clusters
func createClusterURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("clusters")
}

// queryingClusterDetailURL GET /v1.0/{project_id}/clusters/{cluster_id}
func queryingClusterDetailURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId)
}

// deleteClusterURL DELETE  /v1.0/{project_id}/clusters/{cluster_id}
func deleteClusterURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId)
}

// listClustersDetailsURL /v1.0/{project_id}/clusters
func listClustersDetailsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("clusters")
}

// extendInstanceStorageURL /v1.0/{project_id}/clusters/{cluster_id}/role_extend
func extendInstanceStorageURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "role_extend")
}

// listFlavorsURL /v1.0/{project_id}/es-flavors
func listFlavorsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("es-flavors")
}

// restartClusterURL /v1.0/{project_id}/clusters/{cluster_id}/restart
func restartClusterURL(c *golangsdk.ServiceClient, clusterId string) string {
	return c.ServiceURL("clusters", clusterId, "restart")
}
