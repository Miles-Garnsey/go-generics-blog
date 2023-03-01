package reconciliation

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
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
		Name:      U(&desiredObject).GetName(),
		Namespace: U(&desiredObject).GetNamespace(),
	}
	currentObject := new(T)
	err := kClient.Get(ctx, desiredObjectName, U(currentObject))
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	U(&desiredObject).DeepCopyInto(currentObject)
	return nil
}
