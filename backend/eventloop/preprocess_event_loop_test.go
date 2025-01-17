package eventloop

import (
	"context"
	"fmt"
	"testing"

	sharedutil "github.com/redhat-appstudio/managed-gitops/backend-shared/util"
	managedgitopsv1alpha1 "github.com/redhat-appstudio/managed-gitops/backend/apis/managed-gitops/v1alpha1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	fake "sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestPreprocessEventLoop(t *testing.T) {

	ctx := context.Background()

	scheme, argocdNamespace, kubesystemNamespace, workspace := genericTestSetup(t)

	gitopsDepl := &managedgitopsv1alpha1.GitOpsDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-gitops-depl",
			Namespace: workspace.Name,
			UID:       uuid.NewUUID(),
		},
	}

	informer := sharedutil.ListEventReceiver{}

	k8sClientOuter := fake.NewClientBuilder().WithScheme(scheme).WithObjects(gitopsDepl, workspace, argocdNamespace, kubesystemNamespace).Build()
	k8sClient := &sharedutil.ProxyClient{
		InnerClient: k8sClientOuter,
		Informer:    &informer,
	}

	fakeEventLoopChannel := make(chan eventLoopEvent)
	fakeEventLoop := controllerEventLoop{
		eventLoopInputChannel: fakeEventLoopChannel,
	}

	channel := make(chan eventLoopEvent)

	go preprocessEventLoopRouter(channel, &fakeEventLoop)

	event := eventLoopEvent{
		eventType: DeploymentModified,
		request: reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: gitopsDepl.Namespace,
				Name:      gitopsDepl.Name,
			},
		},
		reqResource: managedgitopsv1alpha1.GitOpsDeploymentTypeName,
		client:      k8sClient,
		workspaceID: getWorkspaceIDFromNamespaceID(*workspace),
	}

	fmt.Println("sent event")
	channel <- event

	fmt.Println("waiting for response")
	fmt.Printf("response: %v\n", <-fakeEventLoop.eventLoopInputChannel)

	if err := k8sClient.Delete(ctx, gitopsDepl); err != nil {
		assert.Nil(t, err)
	}

	event = eventLoopEvent{
		eventType: DeploymentModified,
		request: reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: gitopsDepl.Namespace,
				Name:      gitopsDepl.Name,
			},
		},
		reqResource: managedgitopsv1alpha1.GitOpsDeploymentTypeName,
		client:      k8sClient,
		workspaceID: getWorkspaceIDFromNamespaceID(*workspace),
	}

	fmt.Println("sent event")
	channel <- event
	fmt.Println("waiting for response")
	fmt.Printf("response: %v\n", <-fakeEventLoop.eventLoopInputChannel)

}
