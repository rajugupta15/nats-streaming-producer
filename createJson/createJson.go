package CreateJson

import (
	"encoding/json"
	"fmt"
	"log"

	types "github.com/rajugupta15/nats-streaming-producer/types"
)

func CreateJson(timestamp int64, randstr string) string {
	j := types.J{
		Key: "3QFDf0Yb4fwh184rG8a+q/lDTl/UVBot8IUcDrtD7uo=",
		Events: []types.Events{
			types.Events{
				Event:     "hello:alsjdlkfadflas:aksdklfjlajsdlf:laksjkldfjald",
				Timestamp: timestamp,
				Properties: types.Properties{
					LocalOrderID:       "iwueyrtuweyowe",
					LocalPaymentID:     randstr,
					MarchantAppName:    "pwoeiurtpoo",
					MerchantAppVersion: "woeiruowiuer",
					MerchantAppBuild:   1,
					MerchantOptions: types.MerchantOptions{
						Name:        "wpoei[wpeotiweti",
						Description: "oietoiweotwetetuowei",
						Image:       "https://www.thegeeklinux.com/3g7nmJC.png",
						Currency:    "wopeiurtopiwur",
						Amount:      "100",
						Prefill: types.Prefill{
							Contact: "8888888888",
							Name:    "opweiurtopwut",
							Email:   "support@thegeeklinux.com",
						},
						External: types.External{
							Wallets: []types.Wallets{
								types.Wallets{
									"qoewo",
								},
							},
						},
						Key: "opqiwueorpquworuqpower",
					},
				},
			},
		},
		Context: types.Context{
			Mode: "test",
			Device: types.Device{
				ID:           "qowieurpoqiwuer",
				Manufacturer: "qowieuopquir",
				Model:        "oqiwueporquroie",
				Name:         "qoiewuopriquweor",
				Type:         "qoieuroquor",
				Version:      "oqiwuroquwr",
			},
			Sdk: types.Sdk{
				Version: "oqiwueropquew",
				Type:    "oqiewuropqiureo",
			},
			Network: types.Network{
				Bluetooth:           false,
				Carrier:             "oqiwuerpoqiuw",
				Cellular:            false,
				CellularNetworkType: "qpoiweruqoiweur",
				Wifi:                true,
			},
			Screen: types.Screen{
				Density: 2.625,
				Width:   1080,
				Height:  1920,
			},
			Locale:           "en-US",
			Timezone:         "Asia/Calcutta",
			UserAgent:        "oiuweopiuqowreupqioweurpoiqueopriwupoqiweurpqoirop",
			WebviewUserAgent: "qpwoieurpoqiuropeurqpowieuropqiwueroiquweorqweurqopwieurpqruoqiwureop",
		},
	}
	output, err := json.MarshalIndent(&j, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	str := fmt.Sprintf("%s", output)
	return str
}
