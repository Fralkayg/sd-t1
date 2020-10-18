# sd-t1

Integrantes:
	Camilo Farah, rol 201773534-1
	Christian Sepúlveda, rol 201791003-8

Para ejecutar:
	Para iniciar el servidor RabbitMQ, ejecutar el siguiente comando tanto en ssh root@dist54 y en ssh root@dist56:
		>> /sbin/service rabbitmq-server start

	Luego, cada maquina virtual (VM) tiene asignada una de las cuatro entidades. Por lo tanto, se debe ejecutar el siguiente comando para cada uno de ellas:
		- Para ssh root@dist53:  
		>> make clientes
		
		- Para ssh root@dist54:
		>> make logistica

		- Para ssh root@dist55:
		>> make camiones

		- Para ssh root@dist56:
		>> make finanza

Consideraciones:
	- El orden de ejecución debe ser así:
		- Logistica en dist54
		- Finanzas en dist56
		- Clientes en dist53
		- Camiones en dist55
	- Es necesario que se respete el orden de ejecución para asegurar que logistica y finanzas manejen la información correctamente, así como también para que tanto cliente y camión no tengan un comportamiento fuera de lo normal debido a que logistica no es capaz de recibir sus solicitudes.
	- Si desea utilizar sus propios datos para los pedidos de paquetes de clientes, modifique los archivos csv ("pymes.csv" y/o "retail.csv") como corresponda:
		- La primera línea del archivo se ignorará, pues equivaldrá a los headers. Luego:
			- Para retail: id,producto,valor,tienda,destino
			- Para pymes: id,producto,valor,tienda,destino,prioritario

	
