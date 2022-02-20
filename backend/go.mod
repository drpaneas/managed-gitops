module github.com/redhat-appstudio/managed-gitops/backend

go 1.16

require (
	github.com/emicklei/go-restful/v3 v3.7.1
	github.com/go-logr/logr v1.2.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/redhat-appstudio/managed-gitops/backend-shared v0.0.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	sigs.k8s.io/controller-runtime v0.11.1

)

replace github.com/redhat-appstudio/managed-gitops/backend-shared => ../backend-shared
