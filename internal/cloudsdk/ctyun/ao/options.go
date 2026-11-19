package ao

import (
	common "github.com/certimate-go/plugin-provider-ctcccloud-ao-deployer/internal/cloudsdk/ctyun/zz-shared-common"
)

func WithAkSk(ak, sk string) common.OptionsFunc {
	return common.WithAkSk(ak, sk)
}
