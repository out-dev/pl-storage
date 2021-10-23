package main

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func main() {
	cred, err := aztables.NewSharedKeyCredential("myAccountName", "myAccountKey")
	if(err != nil){
		fmt.Print(err.Error())
	}
	serviceClient, err := aztables.NewServiceClient("https://<myAccountName>.table.core.windows.net/", cred, nil)
	if(err != nil){
		fmt.Print(err.Error())
	}
}