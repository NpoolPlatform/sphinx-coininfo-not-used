package constant

import "time"

const (
	ServiceName = "sphinx-coininfo.npool.top"

	GrpcTimeout           = time.Second * 10
	DefaultPageSize int32 = 10
)
