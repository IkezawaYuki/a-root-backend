package external

import (
	"IkezawaYuki/a-root-backend/domain/entity"
)

type OpenaiDto struct {
	CustomerID      int
	DashboardStatus entity.DashboardStatus
	User            string
	System          string
}

type OpenaiResult struct {
	Content string
}
