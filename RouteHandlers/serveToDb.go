package requests

import (
	create "api/FireBase"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var (
	_, firebaseApp     = create.NewFireStore()
	fireStoreClient, _ = firebaseApp.Firestore(context.Background())
)

// PostTopicHandler handles the request and gives the proper response
func PostTopicHandler(c *fiber.Ctx) error {

	// Parse request body
	var newTopic struct {
		Topic string `json:"topic"`
	}
	//Body parser, Error Handler
	if err := c.BodyParser(&newTopic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Creates a reference to a collection to Topic path.
	userCol := fireStoreClient.Collection("Topic")
	//Creates unıq id for document
	docRefUID := userCol.NewDoc()

	// Create new Topic
	topic := map[string]interface{}{
		"topic":    newTopic.Topic,
		"topic_id": docRefUID.ID,
	}

	// Add Topic to Firestore
	_, err := docRefUID.Set(context.Background(), &topic)
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
}

// PostResearchHandler handles the request and gives the proper response
func PostResearchHandler(c *fiber.Ctx) error {

	// Parse request body
	var newResearch struct {
		ResearchHeader      string `json:"research_header"`
		ResearchContent     string `json:"research_content"`
		ResearchCreator     string `json:"research_creator"`
		ResearchContributor string `json:"research_contributor"`
		ResearchTopicId     string `json:"research_topic_id"`
	}

	//Body Parser,Error Handler
	if err := c.BodyParser(&newResearch); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Creates a reference to a collection group to Research path.
	colPath := fmt.Sprintf("Topic/%s/%s", newResearch.ResearchTopicId, newResearch.ResearchHeader)
	userCol := fireStoreClient.Collection(colPath)

	// Create new research
	research := map[string]interface{}{
		"research_header":      newResearch.ResearchHeader,
		"research_content":     newResearch.ResearchContent,
		"research_creator":     newResearch.ResearchCreator,
		"research_contributor": newResearch.ResearchContributor,
	}

	// Add Research to Firestore
	_, err := userCol.Doc(newResearch.ResearchHeader).Set(context.Background(), &research)
	if err != nil {
		log.Print(err, colPath)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add Research to Firestore",
		})
	}

	resp := fiber.Map{
		"message": "Research created successfully",
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}