# XBOT
Telegram xBot that receives messages and process by handlers.

## Structure
#### [.github](../.github) Contains the workflows for GitHub Action.
#### [build](../build) Packaging.
#### [deployments](../deployments) IaaS, PaaS, system and container orchestration deployment configurations and templates (helm charts for k8s deploy)
#### [cmd](../cmd) Main applications for this project.
#### [configs](../configs) Configuration file templates or default configs.
#### [docs](../docs) Contains all the documentation.
#### [internal](../internal) Private application and library code.

## How to add handler?
- add new handler (see example `hue.go`) in package `internal/handler/` implementing `Handler` interface
```go
type Handler interface {
	// Handle updates and return "true" if update processed to user
	Handle(bot *tgbotapi.BotAPI, update *tgbotapi.Update) (bool, error)
	// Setup handler with properties from config
	Setup(props map[string]interface{})
}
```
- add created handler in the list of all handlers
```go
var allHandlers = map[string]Handler{
	"hue": &HueHandler{},
}
```
- add configuration to `configs/default.yaml` (don't forget about helm values)
```yaml
...
handlers:
  - name: hue
    properties:
      min: 6
      max: 10
...
```

## Configuration
Path to config file will be read from env `APPLICATION_CONFIG` or default `configs/default.yaml`
```yaml
bot:
  debug: true #will print telegram api debug info about requests/responses
  username: hue_1338_bot
  token: #will be overriden by env BOT_TOKEN if this field is empty
handlers: #list of enabled handlers
  - name: hue #handler's name
    properties: #key-value properties that you can access in Setup method of handler
      min: 6
      max: 10
handle_policy: 0 # 0 -> ALL (all handlers will process update), 1 -> FIRST (only first handler that return "true" from Handle method)
update_config:
  timeout: 10 #seconds to request telegram api
log:
  level: DEBUG #logging level
```