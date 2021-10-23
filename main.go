package main

import (
	SC "pl-storage/storage_configuration"

	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

func main() {
	var options *SC.StorageConfiguration = &SC.StorageConfiguration{}
	options.GetDefault()
	var azCfg, err = options.ToAzureTableClientOptions()
	serviceClient, err := aztables.NewServiceClientFromConnectionString("", &azCfg)
	if(err != nil){
		fmt.Print(err.Error())
	}
	tableOptions := aztables.CreateTableOptions{}
	serviceClient.CreateTable(context.Background(), "testtable", &tableOptions)
	
	var data map[string]interface{} = map[string]interface{}{
		"data": "simon",
	}

	myEntity := aztables.EDMEntity{
		Entity:     aztables.Entity{PartitionKey: "001234", RowKey: "RedMarker"},
		Properties: data,
	}

	marshalled, err := myEntity.MarshalJSON()
	
	client := serviceClient.NewClient("testtable")
	client.AddEntity(context.Background(), marshalled, nil)
}