package utils

import (
	"github.com/urfave/cli/v2"
	"istio-service-mesh/constants"
	"log"
	"os"
)

func InitFlags() error {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Usage:       "port for the application. default:8654",
			Destination: &constants.ServicePort,
			EnvVars:      []string{"PORT"},

		},
		&cli.StringFlag{
			Name:        "logging_engine_url",
			Usage:       "logging ip:port",
			Destination: &constants.LoggingURL,
			EnvVars:      []string{"LOGGING_ENGINE_URL"},
		},
		&cli.StringFlag{
			Name:        "kubernetes_engine_url",
			Usage:       "kubernetes ip:port ",
			Destination: &constants.KubernetesEngineURL,
			EnvVars:      []string{"KUBERNETES_ENGINE_URL"},
		},
		&cli.StringFlag{
			Name:        "redis_url",
			Usage:       "cluster ip:port ",
			Destination: &constants.NotificationURL,
			EnvVars:      []string{"REDIS_ENGINE_URL"},
		},
		&cli.StringFlag{
			Name:        "k8s_engine_url",
			Usage:       "cluster ip:port ",
			Destination: &constants.K8sEngineGRPCURL,
			EnvVars:      []string{"K8S_ENGINE_URL"},
		},
		&cli.StringFlag{
			Name:        "vault_url",
			Usage:       "cluster ip:port ",
			Destination: &constants.VaultURL,
			EnvVars:      []string{"VAULT_ENGINE_URL"},
		},
		&cli.StringFlag{
			Name:        "rbac_url",
			Usage:       "cluster ip:port ",
			Destination: &constants.RbacURL,
			EnvVars:      []string{"RBAC_URL"},
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
