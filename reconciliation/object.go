package reconciliation

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ReconcileObject(ctx context.Context, kClient client.Client, desiredObject client.Object) error {
	desiredObjectName := types.NamespacedName{
		Name:      desiredObject.GetName(),
		Namespace: desiredObject.GetNamespace(),
	}
	var currentObject client.Object

	err := kClient.Get(ctx, desiredObjectName, currentObject)
	if err != nil {
		return err
	}
	desiredObject.DeepCopyInto(currentObject)
	return nil
}
