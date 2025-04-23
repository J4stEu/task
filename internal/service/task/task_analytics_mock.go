package task

import (
	"context"
	"fmt"
	"time"
)

func (pm *TaskMock) WatchAnalytics(ctx context.Context) {
	interval := time.Second * 5

	for {
		select {
		case <-ctx.Done():
			return
		default:
			users, err := pm.repository.User.GetAll()
			if err != nil {
				fmt.Printf("Failed to watch analytics", err)
				SleepWithCTX(ctx, interval)
				continue
			}

			for i := range users {
				userAnalytics, err := pm.GetAnalyticsByUserID(users[i].ID)
				if err != nil {
					fmt.Println("Failed to watch analytics", err)
					continue
				}

				fmt.Printf("Analytics of user %v: %v\n", users[i].ID, userAnalytics)
			}

			SleepWithCTX(ctx, interval)
		}
	}
}
