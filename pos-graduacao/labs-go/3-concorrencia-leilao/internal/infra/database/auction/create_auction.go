package auction

import (
	"context"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection      *mongo.Collection
	auctionInterval time.Duration
	mu              sync.Mutex
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	auctionRepo := &AuctionRepository{
		Collection:      database.Collection("auctions"),
		auctionInterval: getAuctionInterval(),
	}

	auctionRepo.startAuctionMonitor(context.Background())

	return auctionRepo
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	return nil
}

// getAuctionInterval reads the AUCTION_INTERVAL environment variable
// and returns the duration for auction validity. Default is 5 minutes.
func getAuctionInterval() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		return time.Minute * 5
	}

	return duration
}

// startAuctionMonitor starts a goroutine that periodically checks for expired auctions
// and automatically closes them by updating their status to Completed.
func (ar *AuctionRepository) startAuctionMonitor(ctx context.Context) {
	ticker := time.NewTicker(ar.auctionInterval)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ar.closeExpiredAuctions(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}

// closeExpiredAuctions finds all active auctions that have exceeded their duration
// and updates their status to Completed.
func (ar *AuctionRepository) closeExpiredAuctions(ctx context.Context) {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	now := time.Now().Unix()
	expirationTime := now - int64(ar.auctionInterval.Seconds())

	filter := bson.M{
		"status":    auction_entity.Active,
		"timestamp": bson.M{"$lte": expirationTime},
	}

	update := bson.M{
		"$set": bson.M{
			"status": auction_entity.Completed,
		},
	}

	result, err := ar.Collection.UpdateMany(ctx, filter, update)
	if err != nil {
		logger.Error("Error trying to close expired auctions", err)
		return
	}

	if result.ModifiedCount > 0 {
		logger.Info("Closed expired auctions")
	}
}
