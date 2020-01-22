module istio-service-mesh

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/pkg/errors v0.8.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/swaggo/swag v1.6.5
	github.com/urfave/cli/v2 v2.0.0
	go.opencensus.io v0.22.2
	golang.org/x/build v0.0.0-20191220001908-17a7d8724fa7
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	golang.org/x/sys v0.0.0-20191218084908-4a24b4065292 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/genproto v0.0.0-20191216205247-b31c10ee225f // indirect
	google.golang.org/grpc v1.26.0
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.2.7 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
	istio.io/api v0.0.0-20191218031825-7bafbd24c11c
	istio.io/client-go v0.0.0-20191218043923-5fad2566daf6
	istio.io/istio v0.0.0-20191218042323-ae27ee6c4bf5
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.16.4
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/utils v0.0.0-20191218082557-f07c713de883 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.4
)
