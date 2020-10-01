package test

import (
	"fmt"
	"testing"

	"/modules/ipohelper"
)

var approvedRegions = []string{
	"uksouth",
	"ukwest",
}

func TestSub(*testing.T) {

	azureRegion := ipohelper.GetRandomRegion(t, approvedRegions, nil, "8206a2a5-073a-41d3-9bbe-c6c7a8ab67e6")
	fmt.Println(azureRegion)
	// sub := ipohelper.getTargetAzureSubscription("8206a2a5-073a-41d3-9bbe-c6c7a8ab67e6")
	// fmt.Println(sub)
}
