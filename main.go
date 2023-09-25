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
		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
       		"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
	}

	// We now create a folder `my-folder` in order to migrate our secrets in it.
	folderPath := "/my-folder"
	folder, _ := api.CreateFolder(&secret_manager.CreateFolderRequest{
		ProjectID: projectID,
		Name:      "my-folder",
		Path:      &folderPath,
	})
	fmt.Println("Folder : ")
	fmt.Println("Name: ", folder.Name, " Path: ", folder.Path)

	// We migrate our secrets into the folder `my-folder`
	destinationPath := "/my-folder"
	for _, secretId := range secretIds {
		updatedSecret, _ := api.UpdateSecret(&secret_manager.UpdateSecretRequest{
			SecretID: secretId,
			Path:     &destinationPath,
		})
		fmt.Println("updated secret : ")
		fmt.Println("ID: ", updatedSecret.ID, " Name: ", updatedSecret.Name, " Path: ", updatedSecret.Path)
	}
}
