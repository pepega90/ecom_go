package config

import (
	"github.com/midtrans/midtrans-go"
)

func SetupMidtransKeyAccess() {
	midtrans.ServerKey = "SB-Mid-server-PIJxidJtnmbzXUEj6eIXg71_"
	midtrans.Environment = midtrans.Sandbox
	midtrans.ClientKey = "SB-Mid-client-NBdXSLxw0cOZivxF"
}
