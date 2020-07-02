module bitbucket.org/cloudplex-devs/istio-service-mesh

go 1.13

require (
	bitbucket.org/cloudplex-devs/kubernetes-services-deployment v0.0.0-20200507175720-6940251ff5c9
	bitbucket.org/cloudplex-devs/microservices-mesh-engine v1.4.4-0.20200622095151-d1d493546261
	github.com/asaskevich/govalidator v0.0.0-20200428143746-21a406dcc535 // indirect
	github.com/astaxie/beego v1.12.1
	github.com/go-openapi/spec v0.19.8 // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.1.1
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/urfave/cli/v2 v2.2.0
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	go.opencensus.io v0.22.3
	golang.org/x/build v0.0.0-20200508193432-bf27e2732389
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120 // indirect
	golang.org/x/sys v0.0.0-20200513112337-417ce2331b5c // indirect
	golang.org/x/tools v0.0.0-20200515010526-7d3b6ebf133d // indirect
	google.golang.org/genproto v0.0.0-20200514193133-8feb7f20f2a2 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.3.0
	istio.io/api v0.0.0-20200615162408-9b5293c30ef5
	istio.io/client-go v0.0.0-20200615164228-d77b0b53b6a0
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	istio.io/api => istio.io/api v0.0.0-20200615162408-9b5293c30ef5
	istio.io/client-go => istio.io/client-go v0.0.0-20200615164228-d77b0b53b6a0
	istio.io/gogo-genproto => istio.io/gogo-genproto v0.0.0-20200615162707-ab08778e3a4e // indirect
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
)
