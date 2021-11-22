package db

import (
	"context"
	"testing"
	"time"
)

func TestWithoutTimeout(t *testing.T) {
	_, err := Connection.Query("SELECT sleep(10)")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("success")
}

func TestWithTimout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := Connection.QueryContext(ctx, "SELECT SLEEP(10)")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("SUCCESS")
}