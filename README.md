# Empezar

## Desplegar lamdba en produccion
Debemos editar la instruccion del Makefile cambiando el `username` por nuestro usuario de aws

```sh
deploy_prod: build
	serverless deploy --stage prod --aws-profile username
```

El `username` debe coincidir con el que se encuentra en nuestro archivo de `credentials` de aws

```sh
cat ~/.aws/credentials
```

Correr el siguiente comando
```sh
make deploy_prod
```
