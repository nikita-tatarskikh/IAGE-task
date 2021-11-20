package api

import (
	"IAGE-test-task/services/task1"
	"IAGE-test-task/services/task2"
	"go.uber.org/dig"

	_ "github.com/lib/pq"
)

func BindTask1Repository(redisRepo *task1.RedisRepository) task1.Repository {
	return redisRepo
}

func BindTask1Service(service *task1.PlainService) task1.Service {
	return service
}

func BindTask2Security(redisRepo *task2.HMAC) task2.Security {
	return redisRepo
}

func BindTask2Service(service *task2.PlainService) task2.Service {
	return service
}


func BuildInRuntime() (*App, error) {
	c := dig.New()
	servicesConstructors := []interface{}{
		NewDB,
		task1.NewRepository,
		task1.NewService,
		task1.NewHTTPProvider,
		task2.NewHMAC,
		task2.NewService,
		task2.NewHTTPProvider,
		BindTask1Repository,
		BindTask1Service,
		BindTask2Security,
		BindTask2Service,
		NewRouter,
		NewListener,
		NewApp,
	}

	for _, service := range servicesConstructors {
		if err := c.Provide(service); err != nil {
			return nil, err
		}
	}

	var result *App
	err := c.Invoke(func(a *App) {
		result = a
	})
	return result, err
}