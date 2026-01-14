package auction

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCloseExpiredAuctions(t *testing.T) {
	// Set auction interval to 2 seconds for faster testing
	os.Setenv("AUCTION_INTERVAL", "2s")
	defer os.Unsetenv("AUCTION_INTERVAL")

	// Setup MongoDB test connection
	ctx := context.Background()
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:admin@localhost:27017/auctions?authSource=admin"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Skipf("MongoDB not available: %v", err)
		return
	}
	defer client.Disconnect(ctx)

	// Use a test database
	testDB := client.Database("auctions_test")
	defer testDB.Drop(ctx)

	// Create repository
	repo := NewAuctionRepository(testDB)

	// Create an expired auction (timestamp 5 seconds ago)
	expiredAuction := &auction_entity.Auction{
		Id:          "expired-auction-123",
		ProductName: "Test Product",
		Category:    "Electronics",
		Description: "Test description for expired auction",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now().Add(-5 * time.Second),
	}

	// Create a recent active auction (just created)
	activeAuction := &auction_entity.Auction{
		Id:          "active-auction-456",
		ProductName: "Active Product",
		Category:    "Books",
		Description: "Test description for active auction",
		Condition:   auction_entity.Used,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	// Insert both auctions
	if internalErr := repo.CreateAuction(ctx, expiredAuction); internalErr != nil {
		t.Fatalf("Failed to create expired auction: %v", internalErr)
	}

	if internalErr := repo.CreateAuction(ctx, activeAuction); internalErr != nil {
		t.Fatalf("Failed to create active auction: %v", internalErr)
	}

	// Manually call closeExpiredAuctions to test the logic
	repo.closeExpiredAuctions(ctx)

	// Verify that the expired auction was closed
	var expiredResult AuctionEntityMongo
	err = repo.Collection.FindOne(ctx, bson.M{"_id": expiredAuction.Id}).Decode(&expiredResult)
	if err != nil {
		t.Fatalf("Failed to find expired auction: %v", err)
	}

	if expiredResult.Status != auction_entity.Completed {
		t.Errorf("Expected expired auction status to be Completed, got %v", expiredResult.Status)
	}

	// Verify that the active auction is still active
	var activeResult AuctionEntityMongo
	err = repo.Collection.FindOne(ctx, bson.M{"_id": activeAuction.Id}).Decode(&activeResult)
	if err != nil {
		t.Fatalf("Failed to find active auction: %v", err)
	}

	if activeResult.Status != auction_entity.Active {
		t.Errorf("Expected active auction status to remain Active, got %v", activeResult.Status)
	}
}

func TestGetAuctionInterval(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		expected time.Duration
	}{
		{
			name:     "Valid duration from env",
			envValue: "10s",
			expected: 10 * time.Second,
		},
		{
			name:     "Invalid duration defaults to 5 minutes",
			envValue: "invalid",
			expected: 5 * time.Minute,
		},
		{
			name:     "Empty env defaults to 5 minutes",
			envValue: "",
			expected: 5 * time.Minute,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv("AUCTION_INTERVAL", tt.envValue)
			} else {
				os.Unsetenv("AUCTION_INTERVAL")
			}
			defer os.Unsetenv("AUCTION_INTERVAL")

			result := getAuctionInterval()
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAuctionAutoClosureWithGoroutine(t *testing.T) {
	// Set a very short interval for testing
	os.Setenv("AUCTION_INTERVAL", "1s")
	defer os.Unsetenv("AUCTION_INTERVAL")

	// Setup MongoDB test connection
	ctx := context.Background()
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:admin@localhost:27017/auctions?authSource=admin"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Skipf("MongoDB not available: %v", err)
		return
	}
	defer client.Disconnect(ctx)

	// Use a test database
	testDB := client.Database("auctions_test_goroutine")
	defer testDB.Drop(ctx)

	// Create repository (this will start the background goroutine)
	repo := NewAuctionRepository(testDB)

	// Create an auction with an old timestamp
	oldAuction := &auction_entity.Auction{
		Id:          "old-auction-789",
		ProductName: "Old Product",
		Category:    "Furniture",
		Description: "Test description for old auction",
		Condition:   auction_entity.Refurbished,
		Status:      auction_entity.Active,
		Timestamp:   time.Now().Add(-3 * time.Second),
	}

	if internalErr := repo.CreateAuction(ctx, oldAuction); internalErr != nil {
		t.Fatalf("Failed to create old auction: %v", internalErr)
	}

	// Wait for the goroutine to process (interval + some buffer)
	time.Sleep(2 * time.Second)

	// Verify the auction was automatically closed
	var result AuctionEntityMongo
	err = repo.Collection.FindOne(ctx, bson.M{"_id": oldAuction.Id}).Decode(&result)
	if err != nil {
		t.Fatalf("Failed to find old auction: %v", err)
	}

	if result.Status != auction_entity.Completed {
		t.Errorf("Expected auction to be automatically closed (Completed), got %v", result.Status)
	}
}

func TestConcurrentAuctionClosures(t *testing.T) {
	os.Setenv("AUCTION_INTERVAL", "2s")
	defer os.Unsetenv("AUCTION_INTERVAL")

	ctx := context.Background()
	mongoURI := os.Getenv("MONGODB_URL")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:admin@localhost:27017/auctions?authSource=admin"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Skipf("MongoDB not available: %v", err)
		return
	}
	defer client.Disconnect(ctx)

	testDB := client.Database("auctions_test_concurrent")
	defer testDB.Drop(ctx)

	repo := NewAuctionRepository(testDB)

	// Create multiple expired auctions
	numAuctions := 10
	for i := 0; i < numAuctions; i++ {
		auction := &auction_entity.Auction{
			Id:          "concurrent-auction-" + string(rune(i)),
			ProductName: "Concurrent Product",
			Category:    "Test",
			Description: "Test description for concurrent closure",
			Condition:   auction_entity.Used,
			Status:      auction_entity.Active,
			Timestamp:   time.Now().Add(-5 * time.Second),
		}

		if internalErr := repo.CreateAuction(ctx, auction); internalErr != nil {
			t.Fatalf("Failed to create auction %d: %v", i, internalErr)
		}
	}

	// Trigger closure
	repo.closeExpiredAuctions(ctx)

	// Count closed auctions
	count, err := repo.Collection.CountDocuments(ctx, bson.M{"status": auction_entity.Completed})
	if err != nil {
		t.Fatalf("Failed to count closed auctions: %v", err)
	}

	if count != int64(numAuctions) {
		t.Errorf("Expected %d auctions to be closed, got %d", numAuctions, count)
	}
}
