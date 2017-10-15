package search

import "log"

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

	result, err := snapshot.sense.Client.Operations().Snapshot(snapshotPath)
	if err != nil {
		log.Printf("Error creating snapshot: %v", err)
		return nil, err
	}

	return map[string]interface{}{
		"snapshot_path": snapshotPath,
		"result":        result,
	}, nil
}
