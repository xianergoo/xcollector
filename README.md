project/
|-- cmd/
|   |-- api-service/
|       |-- main.go
|-- pkg/
|   |-- deviceapi/
|   |   |-- server.go
|   |   |-- handlers.go
|   |   |-- models.go
|   |   |-- ...
|   |-- tcpclient/
|   |   |-- client.go
|   |   |-- manager.go
|   |   |-- ...
|   |-- dataprocessing/
|   |   |-- processor.go
|   |   |-- storage.go
|   |   |-- ...
|-- internal/
|   |-- config/
|   |   |-- config.go
|   |-- log/
|   |   |-- logger.go
|-- go.mod
|-- go.sum
|-- README.md
|-- frontend/
|   |-- package.json
|   |-- public/
|   |   |-- index.html
|   |-- src/
|   |   |-- App.js
|   |   |-- components/
|   |   |   |-- DeviceList.js
|   |   |   |-- DeviceForm.js
|   |   |-- services/
|   |   |   |-- apiService.js
|   |   |-- ...