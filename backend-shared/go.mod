module github.com/redhat-appstudio/managed-gitops/backend-shared

go 1.16

require (
	github.com/emicklei/go-restful/v3 v3.9.0
	github.com/go-logr/logr v1.2.3
	github.com/go-pg/pg/extra/pgdebug v0.2.0
	github.com/go-pg/pg/v10 v10.10.6
	github.com/google/uuid v1.3.0
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/stretchr/testify v1.8.0
	k8s.io/api v0.26.1
	k8s.io/apimachinery v0.26.1
	k8s.io/client-go v0.26.1
	sigs.k8s.io/controller-runtime v0.14.5
)
