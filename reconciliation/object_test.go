package reconciliation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func ReconcileObject_Test(t *testing.T) {
	desiredObject := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cm",
			Namespace: "test-namespace",
		},
	}
	ctx := context.Background()
	client := fake.NewClientBuilder().
		Build()

	err := ReconcileObject(ctx, client, desiredObject)
	assert.NoError(t, err)
}
