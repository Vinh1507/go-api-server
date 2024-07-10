package controllers_test

import (
	"fmt"
	"go-api-server/initializers"
	"go-api-server/models"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_GetUserByID(t *testing.T) {
	// Setup
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_TEST"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

	// Create a user for testing
	user := models.User{Email: "test@example.com", Password: "1234"}
	result := db.Create(&user)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// Test case
	t.Run("Existing user", func(t *testing.T) {
		var fetchedUser models.User

		db.First(&fetchedUser, user.ID)

		assert.NoError(t, err, "should not return error")
		assert.NotNil(t, fetchedUser, "returned user should not be nil")
		assert.Equal(t, user.ID, fetchedUser.ID, "user IDs should match")
		assert.Equal(t, user.Password, fetchedUser.Password, "Password should match")
		assert.Equal(t, user.Email, fetchedUser.Email, "emails should match")
	})

	// t.Run("Non-existing user", func(t *testing.T) {
	// 	var fetchedUser models.User

	// 	db.First(&fetchedUser, user.ID+999999999)
	// 	assert.Error(t, err, "should return error")
	// })

}

type APITestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *APITestSuite) SetupSuite() {
	initializers.ConnectToTestDB()
	suite.db = initializers.TestDB
	// Run migrations here if necessary
	suite.db.AutoMigrate(&models.User{})

}

func (suite *APITestSuite) TearDownSuite() {
	db, err := suite.db.DB()
	if err != nil {
		suite.T().Fatal(err)
	}
	db.Close()
}

func (suite *APITestSuite) SetupTest() {
	// Clear tables or reset database state here
	suite.db.Exec("TRUNCATE TABLE your_table_name RESTART IDENTITY CASCADE;")
}

func (suite *APITestSuite) TearDownTest() {
	// Clear tables or reset database state here if needed
	suite.db.Exec("TRUNCATE TABLE your_table_name RESTART IDENTITY CASCADE;")
}

func (suite *APITestSuite) TestGetEndpoint() {
	// req, err := http.NewRequest("GET", "/api/v1/resource", nil)
	// if err != nil {
	// 	suite.T().Fatal(err)
	// }

	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(myGetEndpointHandler)
	// handler.ServeHTTP(rr, req)

	// assert.Equal(suite.T(), http.StatusOK, rr.Code)
	// // More assertions here
	user := models.User{Email: "test@example.com", Password: "1234"}
	result := suite.db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// Test case

	var err error
	t := suite.T()
	t.Run("Existing user", func(t *testing.T) {
		var fetchedUser models.User

		suite.db.First(&fetchedUser, user.ID)

		assert.NoError(t, err, "should not return error")
		assert.NotNil(t, fetchedUser, "returned user should not be nil")
		assert.Equal(t, user.ID, fetchedUser.ID, "user IDs should match")
		assert.Equal(t, user.Password, fetchedUser.Password, "Password should match")
		assert.Equal(t, user.Email, fetchedUser.Email, "emails should match")
	})
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
