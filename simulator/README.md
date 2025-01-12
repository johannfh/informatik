# Simulator

## WebSocket Communication Documentation

### Message Format

All messages are sent as **JSON**. If a client fails to comply, the server will immediately terminate the connection.

The type of message is determined by the `messageType` field.
Scoping of messages is done by seperating words in the `messageType` field with a "`.`".
Message types are all in lowerCamelCase by convention and case sensitive.

Example: **`topLevelNamespace.someScope.someAttribute.anActionOnTheAttribute`**

### `client` Scope

- **`client.game.water.add`**

  Tell the server to add new water to the game. May be a negative to decrease the water level.

  Minimum water level is `0` but no error will occur when it would go to a negative value because of a client.
  It would simply be held at `0`.

- **`client.game.entities.add`**
  
  Tell the server to spawn a new entity. The type of entity is determined by the `entity` field.

### `server` Scope

- **`server.game.water.updated`**

  Sent to all clients when the water level in the game changed.

- **`server.game.entities.updated`**

  Sent to all clients when entities in the game change.
  For example when a new entity is added, dies, or an entities property changes.

- **`server.game.event`**

  Used for general events in the game such as entity deaths or when a resource like `water` is depleted.

### `error` Scope

Errors can be either `Fatal` or `Non-Fatal`.

`Fatal` errors will cause the websocket connection to close while `Non-Fatal` errors will simply be a no-op.


#### Message Type

The `ErrorMessage` type looks as follows:

```go
type ErrorMessage struct {
	MessageType string `json:"messageType"`
	Message     string `json:"message"`
}
```

or in TypeScript:

```ts
type ErrorMessage = {
    messageType: string,
    message: string,
};
```

#### List of errors

- **`error.unknownMessageType`** (Error: Non-Fatal)

  Indicated an error with the `messageType` field of a prior message.

- **`error.invalidMessageFormat`** (Error: Non-Fatal)

  Indicates an error in a prior message format.

  For example when a message of a certain type is missing a field for that type or the fields type is incorrect.

  Example:

  ```js
  const ChangeWater = {
    messageType: "client.game.water.add",
    water: "5" // water is a string but should be a number
  }
  ```

  This would result in an error sent by the server.

- **`error.unknownEntityError`** (Error: Non-Fatal)
  
  Response to `client.game.entities.add` when the `entity` field contains an unknown entity.
