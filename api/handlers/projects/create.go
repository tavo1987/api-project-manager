package projects

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Project is a type to create an project
type Project struct {
	ProjID string `json:"projId"`
	PK     string `json:"PK"`
	SK     string `json:"SK"`
	OrgID  string `json:"orgId"`
	Name   string `json:"name" validate:"required,min=3,max=32"`
	Type   string `json:"type" validate:"required"`
	Status string `json:"status" validate:"required"`
}

// Create this is a handler to create a new organization
func Create(c *fiber.Ctx) error {
	proj := new(Project)
	if err := c.BodyParser(proj); err != nil {
		c.JSON("something went wrong")
		return err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "codingec",
		Config:            aws.Config{Region: aws.String("us-east-1")},
	}))

	dynamoSvc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	proj.ProjID = uuid.New().String()
	proj.PK = "ORG#" + proj.OrgID
	proj.SK = "PROJ#" + proj.Type + "#" + proj.ProjID
	av, err := dynamodbattribute.MarshalMap(proj)
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

	return c.JSON(proj)
}
