// Code generated by mockery v1.0.0. DO NOT EDIT.

package ls

import (
	context "context"

	annotations "github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/ingress/annotations"

	elbv2 "github.com/aws/aws-sdk-go/service/elbv2"

	mock "github.com/stretchr/testify/mock"

	tg "github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/alb/tg"

	v1beta1 "k8s.io/api/extensions/v1beta1"
)

// MockRulesController is an autogenerated mock type for the RulesController type
type MockRulesController struct {
	mock.Mock
}

// Reconcile provides a mock function with given fields: ctx, listener, ingress, ingressAnnos, tgGroup
func (_m *MockRulesController) Reconcile(ctx context.Context, listener *elbv2.Listener, ingress *v1beta1.Ingress, ingressAnnos *annotations.Ingress, tgGroup tg.TargetGroupGroup) error {
	ret := _m.Called(ctx, listener, ingress, ingressAnnos, tgGroup)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *elbv2.Listener, *v1beta1.Ingress, *annotations.Ingress, tg.TargetGroupGroup) error); ok {
		r0 = rf(ctx, listener, ingress, ingressAnnos, tgGroup)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
