package setup

import (
	key "api/Key"
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func Setup() (*fiber.App, *firestore.Client) {

	app := fiber.New()

	var (
		ctx     context.Context
		appfire *firestore.Client
		userCol *firestore.CollectionRef
		err     error
	)

	jsonData := key.Key()

	// Create a new Firestore client using the Google Application Credentials path
	ctx = context.Background()

	opt := option.WithGRPCDialOption(grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time: 2 * time.Minute,
	}))

	appfire, err = firestore.NewClient(ctx, "minerva-95196", option.WithCredentialsJSON(jsonData), opt)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	//We will add login part later
	// app.Post("/login", func(c *fiber.Ctx) error {

	// })

	app.Post("/users", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Users")

		// Parse request body
		var newUser struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// Create new user
		user := map[string]interface{}{
			"name":     newUser.Name,
			"email":    newUser.Email,
			"password": newUser.Password,
		}

		// Add user to Firestore
		_, err := userCol.Doc(newUser.Email).Set(ctx, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add user to Firestore",
			})
		}

		resp := fiber.Map{
			"message": "User created successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	})

	///////////////////////////////TOPIC//////////////////////////////////////////////

	app.Post("/topic", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Topic")

		// Parse request body
		var newTopic struct {
			Topic string `json:"topic"`
		}

		if err := c.BodyParser(&newTopic); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// Create new Topic
		topic := map[string]interface{}{
			"topic": newTopic.Topic,
		}

		// Add Topic to Firestore
		_, err := userCol.Doc(newTopic.Topic).Set(ctx, topic)
		if err != nil {
			log.Printf("Failed to add Topic to Firestore: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add Topic to Firestore",
			})
		}

		resp := fiber.Map{
			"message": "Topic created successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	})

	//////////////////////////////RESEARCH//////////////////////////////////////////////////

	app.Post("/topic/research", func(c *fiber.Ctx) error {

		userCol = appfire.Collection("Research")

		// Parse request body
		var newResearch struct {
			ResearchHeader      string `json:"research_header"`
			ResearchContent     string `json:"research_content"`
			ResearchCreator     string `json:"research_creator"`
			ResearchContributor string `json:"research_contributor"`
		}

		if err := c.BodyParser(&newResearch); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// Create new user
		research := map[string]interface{}{
			"research_header":      newResearch.ResearchHeader,
			"research_content":     newResearch.ResearchContent,
			"research_creator":     newResearch.ResearchCreator,
			"research_contributor": newResearch.ResearchContributor,
		}

		// Add user to Firestore
		_, err := userCol.Doc(newResearch.ResearchHeader).Set(ctx, research)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add Research to Firestore",
			})
		}

		resp := fiber.Map{
			"message": "Research created successfully",
		}
		return c.Status(fiber.StatusCreated).JSON(resp)
	})

	return app, appfire
}