package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
	"github.com/zakisk/redhat-client/utils"
)

var (
	port string
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploys server to platform P",
	Long: `
examples of command. For example:
	
// upload two files to server
store deploy -p [docker | kubernetes]
`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Setenv("PORT", port)
		if err := deployOnDocker(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	deployCmd.Flags().StringVarP(&port, "port", "p", "9254", "in which platform server should be deployed (default: docker)")
}

func deployOnDocker() error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	spinner := utils.CreateDefaultSpinner("Deploying on docker", "Deployed successfully")
	spinner.Start()
	if err != nil {
		spinner.Stop()
		return fmt.Errorf("Unabel to create docker client, please make sure that docker is installed or running\n%s", err.Error())
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		spinner.Stop()
		return fmt.Errorf("Unabel to get images\n%s", err.Error())
	}

	isImagePulled := false
	for _, image := range images {
		repository := "<none>"
		if len(image.RepoTags) > 0 {
			splitted := strings.Split(image.RepoTags[0], ":")
			repository = splitted[0]
		} else if len(image.RepoDigests) > 0 {
			repository = strings.Split(image.RepoDigests[0], "@")[0]
		}

		isImagePulled = repository == "zakisk/redhat"
	}

	if !isImagePulled {
		resp, err := cli.ImagePull(context.Background(), "zakisk/redhat:latest", types.ImagePullOptions{})
		if err != nil {
			spinner.Stop()
			return err
		}
		defer resp.Close()

		// Decode the JSON response to read the status
		var response map[string]interface{}
		decoder := json.NewDecoder(resp)
		for decoder.More() {
			if err := decoder.Decode(&response); err != nil {
				spinner.Stop()
				log.Println(err.Error())
			}
			status, _ := response["status"].(string)
			fmt.Println(status)
		}
	}

	natPort := nat.Port(port + "/tcp")
	config := &container.Config{
		Image: "zakisk/redhat:latest",
		ExposedPorts: nat.PortSet{
			natPort: struct{}{},
		},
		Env: []string{"PORT=9254"},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			natPort: []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, nil, "redhat")
	if err != nil {
		spinner.Stop()
		return fmt.Errorf("Unabel to create container\n%s", err.Error())
	}

	if err := cli.ContainerStart(context.Background(), resp.ID, container.StartOptions{}); err != nil {
		spinner.Stop()
		return fmt.Errorf("Unabel to start container\n%s", err.Error())
	}

	return nil
}
