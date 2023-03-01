package reconciliation

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Reconcilable[T any] interface {
	client.Object
	DeepCopy() *T
	DeepCopyInto(*T)
	*T
}

func ReconcileObject[T any, U Reconcilable[T]](ctx context.Context, kClient client.Client, desiredObject T) error {
	desiredObjectName := types.NamespacedName{
		Name:      (&desiredObject).GetName(),
		Namespace: (&desiredObject).GetNamespace(),
	}
	currentObject := new(T)
	err := kClient.Get(ctx, desiredObjectName, &currentObject)
	if err != nil {
		return err
	}
	(&desiredObject).DeepCopyInto(currentObject)
	return nil
}
