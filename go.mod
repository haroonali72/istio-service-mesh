module istio-service-mesh

go 1.12

require (
	bitbucket.org/cloudplex-devs/kubernetes-services-deployment v0.0.0-20200406102106-7045f8f84475
	bitbucket.org/cloudplex-devs/microservices-mesh-engine v0.0.0-20200406093420-71af7fc3fe5c
	dmitri.shuralyov.com/gpu/mtl v0.0.0-20191203043605-d42048ed14fd // indirect
	github.com/astaxie/beego v1.12.1
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/creack/pty v1.1.9 // indirect
	github.com/elazarl/goproxy v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20191011121108-aa519ddbe484 // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/gogo/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/gophercloud/gophercloud v0.7.0 // indirect
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_model v0.0.0-20191202183732-d1d2010b5bee // indirect
	github.com/rogpeppe/go-charset v0.0.0-20190617161244-0dc95cdf6f31 // indirect
	github.com/rogpeppe/go-internal v1.5.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/urfave/cli/v2 v2.1.1
	go.opencensus.io v0.22.3
	golang.org/x/build v0.0.0-20200207163221-ab7b028cb90c
	golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340 // indirect
	golang.org/x/image v0.0.0-20191214001246-9130b4cfad52 // indirect
	golang.org/x/mobile v0.0.0-20191210151939-1a1fef82734d // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.2.8
	istio.io/api v0.0.0-20200208020912-9564cdd03c96
	istio.io/client-go v0.0.0-20200206191104-0c72ba04e5a1
	istio.io/gogo-genproto v0.0.0-20200207183027-a3495bac39f9 // indirect
	istio.io/istio v0.0.0-20191218042323-ae27ee6c4bf5 // indirect
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/cli-runtime v0.16.4 // indirect
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/gengo v0.0.0-20191120174120-e74f70b9b27e // indirect
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20191218082557-f07c713de883 // indirect
	sigs.k8s.io/yaml v1.2.0
)

replace (
	k8s.io/api => k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.4
	k8s.io/client-go => k8s.io/client-go v0.16.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.4
)
