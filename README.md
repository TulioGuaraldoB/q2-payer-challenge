# Desafio Back-End Pleno - Quero 2 Pay

### Mid-Level Back-End Challenge

<br />

## Settings
<p> Para rodar este projeto local, é sugerido que se tenha Docker instalado em sua máquina, juntamente com `docker-compose`. </p>

Ao rodar local com `docker` e `docker-compose` basta criar uma cópia do arquivo `.env.example` e mudar o nome do arquivo para apenas `.env`.

<p> Exemplo de `.env` para executar local: </p>

```bash
MYSQL_USER=q2pay
MYSQL_PASSWORD=q2pay
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_DATABASE=q2pay

PORT=8080

AUTHORIZER_URL=https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31
```

<p>No exemplo acima já possui a URL do autorizador utilizado e passado nas instruções desse challenge.</p>

## Technologies
<p> GoLang, MySQL, GORM, Fiber, Docker, Makefile, Swagger </p>

<br />

## Documentation / Swagger
<p> Todas as rotas do projeto poderão ser visualizadas ao executar o projeto e ir até a rota: </p>

<small>Exemplo:</small>
```bash
http://localhost:8080/swagger/
```