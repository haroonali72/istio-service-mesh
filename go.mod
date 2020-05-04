module bitbucket.org/cloudplex-devs/istio-service-mesh

go 1.13

require (
	bitbucket.org/cloudplex-devs/kubernetes-services-deployment v0.0.0-20200501120452-a31e2ef8b654
	bitbucket.org/cloudplex-devs/microservices-mesh-engine v0.0.0-20200501120105-a184b768a8da
	github.com/astaxie/beego v1.12.1
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/pkg/errors v0.9.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/urfave/cli/v2 v2.2.0
	go.opencensus.io v0.22.3
	golang.org/x/build v0.0.0-20200428202702-916311cec4e1
	google.golang.org/grpc v1.29.1
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.2.8
	istio.io/api v0.0.0-20200430220031-f818d6294944
	istio.io/client-go v0.0.0-20200430221616-6b954c6c31e4
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.0+incompatible
	sigs.k8s.io/yaml v1.2.0
)

replace (
	bitbucket.org/cloudplex-devs/microservices-mesh-engine => bitbucket.org/cloudplex-devs/microservices-mesh-engine v0.0.0-20200504074924-416086d04741
	istio.io/api => istio.io/api v0.0.0-20200208020912-9564cdd03c96
	istio.io/client-go => istio.io/client-go v0.0.0-20200206191104-0c72ba04e5a1
	istio.io/gogo-genproto => istio.io/gogo-genproto v0.0.0-20200207183027-a3495bac39f9 // indirect
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
)
