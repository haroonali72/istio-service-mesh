package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/types"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	"context"
	"errors"
	"github.com/google/uuid"
	"math"
	"net/http"
	"time"
)

const abortIndex int8 = math.MaxInt8 / 2

// Context is the most important part. It allows us to pass variables between middleware,
type Context struct {
	context.Context
	index       int8
	Keys        map[string]interface{}
	initialized bool
}

func (c *Context) reset() {

	c.index = -1
	c.Keys = nil
}

// Copy returns a copy of the current context that can be safely used outside the request's scope.
// This has to be used when the context has to be passed to a goroutine.
func (c *Context) Copy() *Context {
	var cp = *c
	cp.index = abortIndex
	return &cp
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}

/************************************/
/******** METADATA MANAGEMENT********/
/************************************/

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) Set(key string, value interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (c *Context) Exists(key string) bool {
	_, exists := c.Keys[key]
	return exists
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (c *Context) Get(key string) (value interface{}, exists bool) {
	value, exists = c.Keys[key]
	return
}
func (c *Context) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val, _ := c.Get(keyAsString)
		return val
	}
	return nil
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64(key string) (f64 float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTime(key string) (t time.Time) {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDuration(key string) (d time.Duration) {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Context) GetStringSlice(key string) (ss []string) {
	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) GetStringMap(key string) (sm map[string]interface{}) {
	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}

func (c *Context) ReadLoggingParameters(r *http.Request) (err error) {
	token := r.Header.Get("token")
	if len(token) <= 0 {
		return errors.New("invalid token")
	}
	companyId := r.Header.Get("company_id")
	user := r.Header.Get("user")
	if companyId == "" || user == "" {
		return errors.New("user or companyID must not be empty")
	}
	c.Set("company_id", companyId)
	c.Set("user", user)
	c.Set("user_id", user)
	c.Set("token", token)
	return nil
}
func (c *Context) InitializeLogger(requestURL, method, path, body string) {

	c.Set("service_name", constants.ServiceName)
	c.Set("http_request", types.LoggingHttpRequest{
		Url:       requestURL,
		Method:    method,
		Path:      path,
		Body:      body,
		RequestId: uuid.New().String(),
	})
	c.initialized = true
}
func (c *Context) AddProjectId(projectId string) {
	c.Set("project_id", projectId)
}

func (c *Context) AddUserId(projectId string) {
	c.Set("user", projectId)
}
func (c *Context) SendLog(message string, severity string, logType []string) {
	for i := 0; i < len(logType); i++ {
		switch constants.Logger(logType[i]) {
		case constants.Backend_logging:
			c.SendBackendLogs(message, severity)
		case constants.Frontend_logging:
			c.SendFrontendLogs(message, severity)

		}
	}
}

func (c *Context) SendBackendLogs(message interface{}, severity string) {
	if c.initialized {
		url := constants.LoggingURL + constants.BACKEND_LOGGING_ENDPOINT
		c.Set("severity", severity)
		c.Set("message", message)
		c.Set("resource_name", "solution")
		_, _, err := utils.Post(url, c.Keys, map[string]string{"Content-Type": "application/json"})
		if err != nil {
			utils.Error.Println(err)
		}
	}
}

func (c *Context) SendFrontendLogs(message interface{}, severity string) {
	url := constants.LoggingURL + constants.FRONTEND_LOGGING_ENDPOINT

	c.Set("severity", severity)
	c.Set("message", message)

	var data types.LoggingRequest
	data.Id = c.GetString("project_id")
	data.Service = constants.ServiceName
	data.Level = severity
	data.Message = message
	data.Type = "Project"
	data.CompanyId = c.GetString("company_id")

	_, _, err := utils.Post(url, data, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		utils.Error.Println(err)
	}
}
