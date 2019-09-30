# NetworkSocketsGolang

El proyecto consta de un cliente GOLANG que envía su IP a un servidor GOLANG, este servidor responde la recepción del mensaje.

El proyecto es creado para analizar el uso de SOCKETS con GO.


# Basado en el proyecto

https://www.thepolyglotdeveloper.com/2017/05/network-sockets-with-the-go-programming-language/

A diferencia del proyecto de referencia:

  * El cliente y el servidor son archivos diferentes (client.go y server.go)
  * El Servidor creado con golang no envia un broadcast del mensaje recibido a los clientes conectados y responde al mismo cliente.
  * El clientes envia sus IP al servidor.
  * El cliente toma la dirección del servidor desde un archivo JSON.
  * El cliente se encuentra "escuchando" la conexión con el servidor mostrando en consola cuando servidor se encuentra o no disponible.
  
# Run Client

    go run client.go 

# Run Client

    go run server.go 





