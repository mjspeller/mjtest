package test

import (
	"fmt"
	"testing"

	// blah "./mjstest/modules/mjshelper"
	"github.com/mjspeller/mjtest/modules/mjshelper"
)

var approvedRegions = []string{
	"uksouth",
	"ukwest",
}

func TestSub(*testing.T) {

	// azureRegion := blah.GetRandomRegion(t, approvedRegions, nil, "8206a2a5-073a-41d3-9bbe-c6c7a8ab67e6")
	azureRegion2 := mjshelper.GetRandomRegion(t, approvedRegions, nil, "8206a2a5-073a-41d3-9bbe-c6c7a8ab67e6")
	// fmt.Println(azureRegion)
	fmt.Println(azureRegion2)
	// sub := ipohelper.getTargetAzureSubscription("8206a2a5-073a-41d3-9bbe-c6c7a8ab67e6")
	// fmt.Println(sub)
}
