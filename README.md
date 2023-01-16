# EJERCICIO RABBITMQ
======================================
EJERCICIO RABBITMQ es un repositorio que contiene la conexión básica y dos pequeños proyectos.
- Hello_World: Contiene el primer tutorial de crear un Producer y un Consumer en RabbitMQ de forma básica.
- Work_Queues: Contiene el segundo tutorial de crear un Producer y un Consumer en RabbitMQ con características adicionales.

## Instalación
Clonar el repositorio
```
$ git clone https://github.com/walter-arzapalo/ejercicio_rabbit.git
```
Instalar dependencias
```
$ go get -u github.com/walter-arzapalo/ejercicio_rabbit
$ go mod tidy
```

## Uso
### Archivo rabbit.yml
Crear un archivo .env para rellenar las variables de entorno.
Generar el archivo yml para la conexión con rabbitMQ.
```
$ go run generate.go
```
En el archivo generado, ubicar las credenciales que se piden para la conexión a su servidor rabbitMQ como el usuario y contraseña.

### Hello World!
Dentro de esta carpeta se encuentra un Producer (quien envia los mensajes al servidor) y un Consumer (quien recepciona el mensaje del servidor).
- Para ejecutar el Producer:
  ```
  $ go run ./Hello_World/send.go
  ```
- Para ejecutar el Consumer:
  ```
  $ go run ./Hello_World/receive.go
  ```

### Work Queues
Dentro de esta carpeta se encuentra un Producer (quien envia los mensajes al servidor) y un Consumer (quien recepciona el mensaje del servidor).

El Producer se encuentra modificado para aceptar inputs del usuario, el cual digitará el número de mensajes que desea enviar y la cantidad de tiempo (expresado en puntos) que desea que demore el procesamiento de cada mensaje.

Se pueden ejecutar cuantos Consumers se desee, ya que automáticamente se asignará 1 mensaje de la cola a cada Consumer y se dispondrá de manera equivalente a cada uno y 

- Para ejecutar el Producer:
  ```
  $ go run ./Work_Queues/new_task.go
  ```
- Para ejecutar el Consumer:
  ```
  $ go run ./Work_Queues/worker.go
  ```
