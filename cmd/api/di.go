package api

import (
	"IAGE-test-task/services/task1"
	"IAGE-test-task/services/task2"
	"IAGE-test-task/services/task3"
	"go.uber.org/dig"

	_ "github.com/lib/pq"
)

func BindTask1Repository(redisRepo *task1.RedisRepository) task1.Repository {
	return redisRepo
}

func BindTask1Service(service *task1.PlainService) task1.Service {
	return service
}

func BindTask2Security(hmac *task2.HMAC) task2.Security {
	return hmac
}

func BindTask2Service(service *task2.PlainService) task2.Service {
	return service
}

func BindTask3Multiplier(mult *task3.Multiply) task3.TCPMultiplier {
	return mult
}

func BindTask3Service(service *task3.PlainService) task3.Service {
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
		task3.NewMultiply,
		task3.NewService,
		task3.NewHTTPProvider,
		BindTask1Repository,
		BindTask1Service,
		BindTask2Security,
		BindTask2Service,
		BindTask3Multiplier,
		BindTask3Service,
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