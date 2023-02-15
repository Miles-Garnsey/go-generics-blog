package reconciliation

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconcilable[T any] interface {
	client.Object
	DeepCopy() T
	DeepCopyInto(T)
}

func ReconcileObject[U Reconcilable[U]](ctx context.Context, kClient client.Client, desiredObject U) error {
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
