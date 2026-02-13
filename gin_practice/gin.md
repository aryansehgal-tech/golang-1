## Create the Server

### `server := gin.Default()`

Creates:

* A router
* Logging middleware
* Recovery middleware (prevents crash on panic)

**Think of it as:**

> “Creates a web server with basic settings”

---

## Create a Route

### `server.GET("/test", handler)`

* `GET` → HTTP method
* `"/test"` → URL path
* `handler` → function that runs when route is hit

**Meaning:**

> When someone sends a GET request to `/test`, run this function.

---

## Handler Function

### `func(ctx *gin.Context)`

* `ctx` = request + response handler
* Contains:

  * Query parameters
  * Request body
  * Headers
  * Response helpers

**Think of it as:**

> The control center for handling HTTP requests.

---

## Send JSON Response

### `ctx.JSON(200, data)`

* `200` → HTTP status code (OK)
* `data` → converted automatically to JSON

---

## `gin.H`

### `gin.H{ "key": "value" }`

* Shortcut for `map[string]interface{}`
* Used to return JSON easily

**Example meaning:**

> Create JSON key-value response

---

## Start Server

### `server.Run(":8080")`

* Starts server on port 8080
* Blocks execution (keeps server running)

If no port is provided:

* Defaults to `:8080`

---

# Request Flow (Mental Model)

1. Client sends request
2. Router matches path
3. Handler runs
4. Response is sent
