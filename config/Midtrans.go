package config

import (
	"github.com/midtrans/midtrans-go"
)

func SetupGlobalMidtransConfigApi() {
	midtrans.ServerKey = "SB-Mid-server-YyE7uWSDeo-SBo5lNU6XUA4l"
	// change value to `midtrans.Production`, if you want change the env to production
	midtrans.Environment = midtrans.Sandbox
}