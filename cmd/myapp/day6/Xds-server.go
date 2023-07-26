package xds

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	// Create a new cache instance
	snapshotCache := cache.NewSnapshotCache(true, cache.IDHash{}, nil)

	// Create a new server instance
	srv := v3.NewServer(context.Background(), snapshotCache, nil)

	// Register a delta handler for dynamic configuration updates
	deltaCallbacks := delta.NewDeltaCallbacks(snapshotCache)
	srv.SetSnapshotHandler(delta.NewHandler(snapshotCache, deltaCallbacks))

	// Start a gRPC server on a given port
	grpcServer := grpc.NewServer()
	envoy.RegisterClusterDiscoveryServiceServer(grpcServer, srv)
	envoy.RegisterEndpointDiscoveryServiceServer(grpcServer, srv)
	envoy.RegisterRouteDiscoveryServiceServer(grpcServer, srv)
	envoy.RegisterSecretDiscoveryServiceServer(grpcServer, srv)

	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
