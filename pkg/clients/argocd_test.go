package clients

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/crossplane-contrib/provider-argocd/apis/v1alpha1"
)

var (
	testAdress = "no.where.example:8443"
)

func Test_resolveServerAddress(t *testing.T) {
	type args struct {
		c  client.Client
		pc v1alpha1.ProviderConfigSpec
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "No address set, error",
			args: args{
				c:  nil,
				pc: v1alpha1.ProviderConfigSpec{},
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, errNoAddrSet)
			},
		},
		{
			name: "ServerAdr set, returns serverAddr",
			args: args{
				c: nil,
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddr: &testAdress,
				},
			},
			want:    testAdress,
			wantErr: assert.NoError,
		},
		{
			name: "Type None set, error",
			args: args{
				c: nil,
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceNone,
					},
				},
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, errNoneSourceType)
			},
		},
		{
			name: "Type Anything set, error",
			args: args{
				c: nil,
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: "Anything",
					},
				},
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, fmt.Sprintf(errServerAdressTypeNotSupport, "Anything"))
			},
		},
		{
			name: "Type Secret set, returns from secret",
			args: args{
				c: fake.NewClientBuilder().WithObjects(&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testsecret",
						Namespace: "testns",
					},
					Data: map[string][]byte{
						"testkey": []byte(testAdress),
					},
				}).Build(),
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceSecret,
						SourceSelector: v1alpha1.SourceSelector{
							SourceReference: v1alpha1.SourceReference{
								Name:      "testsecret",
								Namespace: "testns",
							},
							Key: "testkey",
						},
					},
				},
			},
			want:    testAdress,
			wantErr: assert.NoError,
		},
		{
			name: "Type ConfigMap set, returns from configmap",
			args: args{
				c: fake.NewClientBuilder().WithObjects(&v1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testsecret",
						Namespace: "testns",
					},
					Data: map[string]string{
						"testkey": testAdress,
					},
				}).Build(),
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceConfigMap,
						SourceSelector: v1alpha1.SourceSelector{
							SourceReference: v1alpha1.SourceReference{
								Name:      "testsecret",
								Namespace: "testns",
							},
							Key: "testkey",
						},
					},
				},
			},
			want:    testAdress,
			wantErr: assert.NoError,
		},
		{
			name: "Type ConfigMap set, Config map does not exist, returns error",
			args: args{
				c: fake.NewClientBuilder().WithObjects().Build(),
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceConfigMap,
						SourceSelector: v1alpha1.SourceSelector{
							SourceReference: v1alpha1.SourceReference{
								Name:      "testsecret",
								Namespace: "testns",
							},
							Key: "testkey",
						},
					},
				},
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "not found")
			},
		},
		{
			name: "Type Secret set, Secret does not exist, returns error",
			args: args{
				c: fake.NewClientBuilder().WithObjects().Build(),
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceSecret,
						SourceSelector: v1alpha1.SourceSelector{
							SourceReference: v1alpha1.SourceReference{
								Name:      "testsecret",
								Namespace: "testns",
							},
							Key: "testkey",
						},
					},
				},
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.ErrorContains(t, err, "not found")
			},
		},
		{
			name: "Type ConfigMap set, key does not exists, returns empty string",
			args: args{
				c: fake.NewClientBuilder().WithObjects(&v1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testsecret",
						Namespace: "testns",
					},
					Data: map[string]string{
						"testkey": testAdress,
					},
				}).Build(),
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceConfigMap,
						SourceSelector: v1alpha1.SourceSelector{
							SourceReference: v1alpha1.SourceReference{
								Name:      "testsecret",
								Namespace: "testns",
							},
							Key: "OtherKey",
						},
					},
				},
			},
			want:    "",
			wantErr: assert.NoError,
		},

		{
			name: "Type Secret set, key does not exist in secret, returns empty string",
			args: args{
				c: fake.NewClientBuilder().WithObjects(&v1.Secret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testsecret",
						Namespace: "testns",
					},
					Data: map[string][]byte{
						"testkey": []byte(testAdress),
					},
				}).Build(),
				pc: v1alpha1.ProviderConfigSpec{
					ServerAddressReference: &v1alpha1.ServerReference{
						Source: v1alpha1.ServerAddressSourceSecret,
						SourceSelector: v1alpha1.SourceSelector{
							SourceReference: v1alpha1.SourceReference{
								Name:      "testsecret",
								Namespace: "testns",
							},
							Key: "otherKey",
						},
					},
				},
			},
			want:    "",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := resolveServerAddress(context.TODO(), tt.args.c, tt.args.pc)
			if !tt.wantErr(t, err, "resolveServerAddress()") {
				return
			}
			if got != tt.want {
				t.Errorf("resolveServerAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
