package config

import (
	"github.com/midtrans/midtrans-go"
)

func SetupGlobalMidtransConfigApi() {
	midtrans.ServerKey = "SB-Mid-server-TvgWB_Y9s81-rbMBH7zZ8BHW"
	// change value to `midtrans.Production`, if you want change the env to production
	midtrans.Environment = midtrans.Sandbox
}