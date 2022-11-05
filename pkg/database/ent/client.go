// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	"github.com/MatthewBehnke/apis/pkg/database/ent/authorizationpolicy"
	"github.com/MatthewBehnke/apis/pkg/database/ent/migrate"
	"github.com/MatthewBehnke/apis/pkg/database/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AuthorizationPolicy is the client for interacting with the AuthorizationPolicy builders.
	AuthorizationPolicy *AuthorizationPolicyClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AuthorizationPolicy = NewAuthorizationPolicyClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		AuthorizationPolicy: NewAuthorizationPolicyClient(cfg),
		User:                NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		AuthorizationPolicy: NewAuthorizationPolicyClient(cfg),
		User:                NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AuthorizationPolicy.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the domain clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.AuthorizationPolicy.Use(hooks...)
	c.User.Use(hooks...)
}

// AuthorizationPolicyClient is a client for the AuthorizationPolicy schema.
type AuthorizationPolicyClient struct {
	config
}

// NewAuthorizationPolicyClient returns a client for the AuthorizationPolicy from the given config.
func NewAuthorizationPolicyClient(c config) *AuthorizationPolicyClient {
	return &AuthorizationPolicyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authorizationpolicy.Hooks(f(g(h())))`.
func (c *AuthorizationPolicyClient) Use(hooks ...Hook) {
	c.hooks.AuthorizationPolicy = append(c.hooks.AuthorizationPolicy, hooks...)
}

// Create returns a builder for creating a AuthorizationPolicy domain.
func (c *AuthorizationPolicyClient) Create() *AuthorizationPolicyCreate {
	mutation := newAuthorizationPolicyMutation(c.config, OpCreate)
	return &AuthorizationPolicyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthorizationPolicy entities.
func (c *AuthorizationPolicyClient) CreateBulk(builders ...*AuthorizationPolicyCreate) *AuthorizationPolicyCreateBulk {
	return &AuthorizationPolicyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthorizationPolicy.
func (c *AuthorizationPolicyClient) Update() *AuthorizationPolicyUpdate {
	mutation := newAuthorizationPolicyMutation(c.config, OpUpdate)
	return &AuthorizationPolicyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given domain.
func (c *AuthorizationPolicyClient) UpdateOne(ap *AuthorizationPolicy) *AuthorizationPolicyUpdateOne {
	mutation := newAuthorizationPolicyMutation(c.config, OpUpdateOne, withAuthorizationPolicy(ap))
	return &AuthorizationPolicyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthorizationPolicyClient) UpdateOneID(id int) *AuthorizationPolicyUpdateOne {
	mutation := newAuthorizationPolicyMutation(c.config, OpUpdateOne, withAuthorizationPolicyID(id))
	return &AuthorizationPolicyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthorizationPolicy.
func (c *AuthorizationPolicyClient) Delete() *AuthorizationPolicyDelete {
	mutation := newAuthorizationPolicyMutation(c.config, OpDelete)
	return &AuthorizationPolicyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given domain.
func (c *AuthorizationPolicyClient) DeleteOne(ap *AuthorizationPolicy) *AuthorizationPolicyDeleteOne {
	return c.DeleteOneID(ap.ID)
}

// DeleteOne returns a builder for deleting the given domain by its id.
func (c *AuthorizationPolicyClient) DeleteOneID(id int) *AuthorizationPolicyDeleteOne {
	builder := c.Delete().Where(authorizationpolicy.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthorizationPolicyDeleteOne{builder}
}

// Query returns a query builder for AuthorizationPolicy.
func (c *AuthorizationPolicyClient) Query() *AuthorizationPolicyQuery {
	return &AuthorizationPolicyQuery{
		config: c.config,
	}
}

// Get returns a AuthorizationPolicy domain by its id.
func (c *AuthorizationPolicyClient) Get(ctx context.Context, id int) (*AuthorizationPolicy, error) {
	return c.Query().Where(authorizationpolicy.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthorizationPolicyClient) GetX(ctx context.Context, id int) *AuthorizationPolicy {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AuthorizationPolicyClient) Hooks() []Hook {
	return c.hooks.AuthorizationPolicy
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User domain.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given domain.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given domain.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given domain by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User domain by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
