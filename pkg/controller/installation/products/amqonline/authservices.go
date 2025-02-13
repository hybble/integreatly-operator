package amqonline

import (
	"github.com/enmasseproject/enmasse/pkg/apis/admin/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDefaultAuthServices(ns string) []*v1beta1.AuthenticationService {
	return []*v1beta1.AuthenticationService{
		{
			ObjectMeta: v1.ObjectMeta{
				Name:      "standard-authservice",
				Namespace: ns,
			},
			Spec: v1beta1.AuthenticationServiceSpec{
				Type: "standard",
			},
		},
		{
			ObjectMeta: v1.ObjectMeta{
				Name:      "none-authservice",
				Namespace: ns,
			},
			Spec: v1beta1.AuthenticationServiceSpec{
				Type: "none",
			},
		},
	}

}
