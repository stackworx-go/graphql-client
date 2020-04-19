// Code generated by github.com/stackworx-go/gqlgen-relay, DO NOT EDIT.
package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Client struct {
	http.Client
	Url string
}

type request struct {
	Query string `json:"query"`
	// OperationName string                 `json:"operationName"`
	Variables  interface{}            `json:"variables"`
	Extensions map[string]interface{} `json:"extensions"`
}

type GraphqlError struct {
	Errors []gqlerror.Error
}

func (e *GraphqlError) Error() string {
	return "Graphql Request Failed"
}

type CreateTodoInput struct {
	Text   string `json:"text"`
	UserId string `json:"userId"`
}

var CreateTodoMutation = `mutation CreateTodoMutation ($input: CreateTodoInput!) {
	createTodo(input: $input) {
		todo {
			id
			text
			done
			user {
				id
				name
			}
		}
	}
}
`

type CreateTodoMutationPayloadCreateTodoTodoUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CreateTodoMutationPayloadCreateTodoTodo struct {
	Id   string                                      `json:"id"`
	Text string                                      `json:"text"`
	Done bool                                        `json:"done"`
	User CreateTodoMutationPayloadCreateTodoTodoUser `json:"user"`
}
type CreateTodoMutationPayloadCreateTodo struct {
	Todo CreateTodoMutationPayloadCreateTodoTodo `json:"todo"`
}
type CreateTodoMutationPayload struct {
	CreateTodo CreateTodoMutationPayloadCreateTodo `json:"createTodo"`
}

type responseCreateTodoMutation struct {
	Data   *CreateTodoMutationPayload `json:"data"`
	Errors []gqlerror.Error           `json:errors`
}

func (c *Client) CreateTodoMutation(input CreateTodoInput) (*CreateTodoMutationPayload, error) {
	requestBody, err := json.Marshal(request{
		Query: CreateTodoMutation,
		Variables: map[string]interface{}{
			"input": input,
		},
	})

	if err != nil {
		return nil, err
	}

	resp, err := c.Post(c.Url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request Failed with status code: %d, body: %v", resp.StatusCode, body)
	}

	var payload responseCreateTodoMutation
	err = json.Unmarshal(body, &payload)

	if err != nil {
		return nil, err
	}

	if len(payload.Errors) > 0 {
		return nil, &GraphqlError{Errors: payload.Errors}
	}

	return payload.Data, nil
}

var TodosQuery = `query TodosQuery {
	todos {
		id
		text
		done
		user {
			id
			name
		}
	}
}
`

type TodosQueryPayloadTodosUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type TodosQueryPayloadTodos struct {
	Id   string                     `json:"id"`
	Text string                     `json:"text"`
	Done bool                       `json:"done"`
	User TodosQueryPayloadTodosUser `json:"user"`
}
type TodosQueryPayload struct {
	Todos []TodosQueryPayloadTodos `json:"todos"`
}

type responseTodosQuery struct {
	Data   *TodosQueryPayload `json:"data"`
	Errors []gqlerror.Error   `json:errors`
}

func (c *Client) TodosQuery() (*TodosQueryPayload, error) {
	requestBody, err := json.Marshal(request{
		Query: TodosQuery,
	})

	if err != nil {
		return nil, err
	}

	resp, err := c.Post(c.Url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request Failed with status code: %d, body: %v", resp.StatusCode, body)
	}

	var payload responseTodosQuery
	err = json.Unmarshal(body, &payload)

	if err != nil {
		return nil, err
	}

	if len(payload.Errors) > 0 {
		return nil, &GraphqlError{Errors: payload.Errors}
	}

	return payload.Data, nil
}

var TodosQueryWithVariables = `query TodosQueryWithVariables ($userId: ID!) {
	todos(userId: $userId) {
		id
		text
		done
		user {
			id
			name
		}
	}
}
`

type TodosQueryWithVariablesPayloadTodosUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type TodosQueryWithVariablesPayloadTodos struct {
	Id   string                                  `json:"id"`
	Text string                                  `json:"text"`
	Done bool                                    `json:"done"`
	User TodosQueryWithVariablesPayloadTodosUser `json:"user"`
}
type TodosQueryWithVariablesPayload struct {
	Todos []TodosQueryWithVariablesPayloadTodos `json:"todos"`
}

type responseTodosQueryWithVariables struct {
	Data   *TodosQueryWithVariablesPayload `json:"data"`
	Errors []gqlerror.Error                `json:errors`
}

func (c *Client) TodosQueryWithVariables(userId string) (*TodosQueryWithVariablesPayload, error) {
	requestBody, err := json.Marshal(request{
		Query: TodosQueryWithVariables,
		Variables: map[string]interface{}{
			"userId": userId,
		},
	})

	if err != nil {
		return nil, err
	}

	resp, err := c.Post(c.Url, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request Failed with status code: %d, body: %v", resp.StatusCode, body)
	}

	var payload responseTodosQueryWithVariables
	err = json.Unmarshal(body, &payload)

	if err != nil {
		return nil, err
	}

	if len(payload.Errors) > 0 {
		return nil, &GraphqlError{Errors: payload.Errors}
	}

	return payload.Data, nil
}
