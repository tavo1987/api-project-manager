package organizations

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Organization is a type to create an organization
type Organization struct {
	OrgID uuid.UUID `json:"orgId"`
	PK    string    `json:"PK"`
	SK    string    `json:"SK"`
	Name  string    `json:"name" validate:"required,min=3,max=32"`
	Tier  string    `json:"tier" validate:"required"`
}

// Create this is a handler to create a new organization
func Create(c *fiber.Ctx) error {
	org := new(Organization)
	if err := c.BodyParser(org); err != nil {
		c.JSON("something went wrong")
		return err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "codingec",
		Config:            aws.Config{Region: aws.String("us-east-1")},
	}))

	dynamoSvc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	org.OrgID = uuid.New()
	org.PK = "ORG#" + org.OrgID.String()
	org.SK = "#METADATA#" + org.OrgID.String()
	av, err := dynamodbattribute.MarshalMap(org)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("project-manager"),
	}

	_, err = dynamoSvc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(org)
}
