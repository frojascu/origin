package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_Project = map[string]string{
	"":         "Projects are the unit of isolation and collaboration in OpenShift. A project has one or more members, a quota on the resources that the project may consume, and the security controls on the resources in the project. Within a project, members may have different roles - project administrators can set membership, editors can create and manage the resources, and viewers can see but not access running containers. In a normal cluster project administrators are not able to alter their quotas - that is restricted to cluster administrators.\n\nListing or watching projects will return only projects the user has the reader role on.\n\nAn OpenShift project is an alternative representation of a Kubernetes namespace. Projects are exposed as editable to end users while namespaces are not. Direct creation of a project is typically restricted to administrators, while end users should use the requestproject resource.",
	"metadata": "Standard object's metadata.",
	"spec":     "Spec defines the behavior of the Namespace.",
	"status":   "Status describes the current status of a Namespace",
}

func (Project) SwaggerDoc() map[string]string {
	return map_Project
}

var map_ProjectList = map[string]string{
	"":         "ProjectList is a list of Project objects.",
	"metadata": "Standard object's metadata.",
	"items":    "Items is the list of projects",
}

func (ProjectList) SwaggerDoc() map[string]string {
	return map_ProjectList
}

var map_ProjectRequest = map[string]string{
	"":            "ProjecRequest is the set of options necessary to fully qualify a project request",
	"metadata":    "Standard object's metadata.",
	"displayName": "DisplayName is the display name to apply to a project",
	"description": "Description is the description to apply to a project",
}

func (ProjectRequest) SwaggerDoc() map[string]string {
	return map_ProjectRequest
}

var map_ProjectSpec = map[string]string{
	"":           "ProjectSpec describes the attributes on a Project",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage",
}

func (ProjectSpec) SwaggerDoc() map[string]string {
	return map_ProjectSpec
}

var map_ProjectStatus = map[string]string{
	"":      "ProjectStatus is information about the current status of a Project",
	"phase": "Phase is the current lifecycle phase of the project",
}

func (ProjectStatus) SwaggerDoc() map[string]string {
	return map_ProjectStatus
}
