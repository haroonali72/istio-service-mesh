module istio-service-mesh

go 1.12

require (
	bitbucket.org/cloudplex-devs/kubernetes-services-deployment v0.0.0-20200413141308-024c89b36e0b
	bitbucket.org/cloudplex-devs/microservices-mesh-engine v0.0.0-20200414054241-588ecb82a727
	github.com/astaxie/beego v1.12.1
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/urfave/cli/v2 v2.2.0
	go.opencensus.io v0.22.3
	golang.org/x/build v0.0.0-20200408230101-5bbd558901b3
	google.golang.org/grpc v1.28.1
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.2.8
	istio.io/api v0.0.0-20200407171655-fb462ece86fb
	istio.io/client-go v0.0.0-20200325170329-dc00bbff4229
	k8s.io/api v0.18.1
	k8s.io/apimachinery v0.18.1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/utils v0.0.0-20200327001022-6496210b90e8 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	istio.io/api => istio.io/api v0.0.0-20200208020912-9564cdd03c96
	istio.io/client-go => istio.io/client-go v0.0.0-20200206191104-0c72ba04e5a1
	istio.io/gogo-genproto => istio.io/gogo-genproto v0.0.0-20200207183027-a3495bac39f9 // indirect
	istio.io/istio => istio.io/istio v0.0.0-20191218042323-ae27ee6c4bf5 // indirect
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.4
)
