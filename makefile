logistica:
	/sbin/service rabbitmq-server start
	go run logisticaService.go
finanzas:
	/sbin/service rabbitmq-server start
	go run finanzasClient.go
cliente:
	go run clientesClient.go
camion:
	go run camionClient.go
