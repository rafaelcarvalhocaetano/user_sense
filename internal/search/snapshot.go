package search

import (
	"context"
	"log"
	"time"
)

type SnapshotImpl struct {
	sense *TypesenseClient
}

func NewSnapshot(client *TypesenseClient) *SnapshotImpl {
	return &SnapshotImpl{
		sense: client,
	}
}

func (snapshot *SnapshotImpl) CreateSnapshot() (map[string]interface{}, error) {
	snapshotPath := "/tmp/typesense-data"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := snapshot.sense.Client.Operations().Snapshot(ctx, snapshotPath)
	if err != nil {
		log.Printf("Error creating snapshot: %v", err)
		return nil, err
	}

	return map[string]interface{}{
		"snapshot_path": snapshotPath,
		"result":        result,
	}, nil
}
