module github.com/redhat-appstudio/managed-gitops/backend

go 1.16

require (
	github.com/emicklei/go-restful/v3 v3.9.0
	github.com/go-logr/logr v1.2.4
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.27.7
	github.com/redhat-appstudio/managed-gitops/backend-shared v0.0.0
	github.com/stretchr/testify v1.8.1
	golang.org/x/oauth2 v0.5.0
	k8s.io/api v0.27.2
	k8s.io/apimachinery v0.27.2
	k8s.io/client-go v0.27.2
	sigs.k8s.io/controller-runtime v0.15.0

)

replace github.com/redhat-appstudio/managed-gitops/backend-shared => ../backend-shared
