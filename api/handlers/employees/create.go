package employees

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
type Employee struct {
	EmpID string `json:"empId"`
	PK    string `json:"PK"`
	SK    string `json:"SK"`
	OrgID string `json:"orgId"`
	Name  string `json:"name" validate:"required,min=3,max=32"`
	Dob   string `json:"birthdate" validate:"required"`
	Email string `json:"email" validate:"required"`
}

// Create this is a handler to create a new organization
func Create(c *fiber.Ctx) error {
	emp := new(Employee)
	if err := c.BodyParser(emp); err != nil {
		c.JSON("something went wrong")
		return err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "codingec",
		Config:            aws.Config{Region: aws.String("us-east-1")},
	}))

	dynamoSvc := dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	emp.EmpID = uuid.New().String()
	emp.PK = "ORG#" + emp.OrgID
	emp.SK = "EMP#" + emp.EmpID
	av, err := dynamodbattribute.MarshalMap(emp)
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

	return c.JSON(emp)
}
