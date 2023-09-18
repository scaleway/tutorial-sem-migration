package main

import (
	"fmt"

	secret_manager "github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func main() {
	// Create a Scaleway client
	client, _ := scw.NewClient(scw.WithEnv())
	api := secret_manager.NewAPI(client)
	projectID, _ := client.GetDefaultProjectID()

	// IDs of the secrets to migrate (make sure you replace the IDs below with the IDs of the secrets you want to migrate)
	secretIds := []string{
		"ce8c000e-79df-4e95-9ac6-bb12ef3cd5fa",
		"60bb08e3-2f37-43b9-b932-75e109dbaa58",
        "042a7d50-511d-426f-bdfc-eaf66c81a5bd",
	}

	// We create a folder `new-folder` in order to migrate our secrets in it
	folderPath := "/new-folder"
	folder, _ := api.CreateFolder(&secret_manager.CreateFolderRequest{
		ProjectID: projectID,
		Name:      "new-folder",
		Path:      &folderPath,
	})
	fmt.Println("New folder : ")
	fmt.Println("Name: ", folder.Name, " Path: ", folder.Path)

	// We migrate our secrets into the folder `new-folder`
	destinationPath := "/new-folder"
	for _, secretId := range secretIds {
		updatedSecret, _ := api.UpdateSecret(&secret_manager.UpdateSecretRequest{
			SecretID: secretId,
			Path:     &destinationPath,
		})
		fmt.Println("updated secret : ")
		fmt.Println("ID: ", updatedSecret.ID, " Name: ", updatedSecret.Name, " Path: ", updatedSecret.Path)
	}
}
