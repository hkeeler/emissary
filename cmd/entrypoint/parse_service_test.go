package entrypoint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var serviceTests = []struct {
	Module            moduleResolver
	Service           string
	Namespace         string
	ExpectedService   string
	ExpectedNamespace string
	ExpectedPort      string
}{
	{
		moduleResolver{},
		"service-name.namespace:3000",
		"other-namespace",
		"service-name",
		"namespace",
		"3000",
	},
	{
		moduleResolver{UseAmbassadorNamespaceForServiceResolution: true},
		"service-name.namespace:3000",
		"other-namespace",
		"service-name",
		"namespace",
		"3000",
	},
	{
		moduleResolver{},
		"service-name:3000",
		"other-namespace",
		"service-name",
		"other-namespace",
		"3000",
	},
	{
		moduleResolver{UseAmbassadorNamespaceForServiceResolution: true},
		"service-name:3000",
		"other-namespace",
		"service-name",
		"default",
		"3000",
	},
	{
		moduleResolver{},
		"service-name",
		"other-namespace",
		"service-name",
		"other-namespace",
		"",
	},
	{
		moduleResolver{UseAmbassadorNamespaceForServiceResolution: true},
		"service-name",
		"other-namespace",
		"service-name",
		"default",
		"",
	},
	{
		moduleResolver{},
		"1.2.3.4",
		"other-namespace",
		"1.2.3.4",
		"other-namespace",
		"",
	},
	{
		moduleResolver{UseAmbassadorNamespaceForServiceResolution: true},
		"1.2.3.4:8080",
		"other-namespace",
		"1.2.3.4",
		"default",
		"8080",
	},
	{
		moduleResolver{},
		"1.2.3.4:blah",
		"other-namespace",
		"1",
		"2.3.4:blah",
		"",
	},
	{
		moduleResolver{UseAmbassadorNamespaceForServiceResolution: true},
		"1.2.3.4:blah",
		"other-namespace",
		"1",
		"2.3.4:blah",
		"",
	},
}

func TestParseService(t *testing.T) {
	for _, test := range serviceTests {
		// Make sure we ignore these also.
		for _, prefix := range []string{"", "http://", "https://"} {
			service := fmt.Sprintf("%s%s", prefix, test.Service)
			t.Run(fmt.Sprintf("%s,%s,%v", service, test.Namespace, test.Module.UseAmbassadorNamespaceForServiceResolution), func(t *testing.T) {
				name, namespace, port := test.Module.parseService(service, test.Namespace)
				assert.Equal(t, test.ExpectedService, name)
				assert.Equal(t, test.ExpectedNamespace, namespace)
				assert.Equal(t, test.ExpectedPort, port)
			})
		}
	}
}