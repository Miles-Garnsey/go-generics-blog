package reconciliation

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func ReconcileObject_Test(t *testing.T) {
	desiredCm := corev1.ConfigMap{}
	ctx := context.Background()
	fakeClient := fake.NewFakeClient()
	result := ReconcileObject(ctx, fakeClient, desiredCm)
	assert.NoError(t, result)
}
