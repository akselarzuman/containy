package main

import (
	"context"
	"log"
	"time"

	"github.com/akselarzuman/containy"
	"github.com/akselarzuman/containy/predefined"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	c := containy.New()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		if err := c.Cleanup(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	redis, err := c.CreateContainer(ctx, predefined.RedisConfig)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := c.CreateContainer(ctx, predefined.PostgresConfig(
		"postgres",
		"password",
		"testdb"))
	if err != nil {
		log.Fatal(err)
	}

	localstack, err := c.CreateContainer(ctx, predefined.LocalstackConfig(
		"dynamodb,s3",
		"us-east-1",
	))

	if err != nil {
		log.Fatal(err)
	}

	_ = redis
	_ = postgres
	_ = localstack

	time.Sleep(5 * time.Minute)
}
