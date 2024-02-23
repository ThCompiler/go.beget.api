package main

import (
	"fmt"
	"log"

	"github.com/ThCompiler/go.beget.api/api/dns"
	"github.com/ThCompiler/go.beget.api/api/dns/build"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
)

func main() {
	client := core.Client{
		Login:    "vetan2o5",
		Password: "932usRC26Tp",
	}

	req, err := core.PrepareRequest[result.BoolResult](
		client,
		dns.CallChangeRecords("tmp.thecompiler.pw",
			build.NewBasicRecordsCreator().
				AddARecords(
					build.NewARecords().
						AddRecord(10, "8.8.8.8"),
				).Create(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := req.Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", *resp)
}