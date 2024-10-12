# Go Auth Project

Este projeto é uma implementação de autenticação utilizando JSON Web Tokens (JWT) em Go, para ser utilizado em meus projetos pessoais.

## Funcionalidade

O projeto oferece uma API para autenticação de usuários, que inclui as seguintes funcionalidades:
- Geração de tokens JWT para usuários autenticados.
- Validação de tokens JWT para proteger rotas de acesso.

## Rotas da API

### 1. Login

- **Método**: `POST`
- **Endpoint**: `/login`
- **Descrição**: Gera um token JWT para o usuário, dado um ID, email e senha válidos.
- **Corpo da Requisição**:
    ```json
    {
        "id": "userId",
        "email": "testuser@gmail.com",
    }
    ```
- **Resposta**:
    ```json
    {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```

### 2. Validar Token

- **Método**: `GET`
- **Endpoint**: `/validate`
- **Descrição**: Valida o token JWT fornecido. Se o token for válido, retorna os dados do usuário contidos no token.
- **Cabeçalho da Requisição**:
    ```
    Authorization: Bearer <token>
    ```
- **Resposta**:
    ```json
    {
        "id": "userId",
        "email": "testuser@gmail.com",
    }
    ```

## Instruções para Executar

### Pré-requisitos

- Ter o [Go] instalado na sua máquina.
- Ter o arquivo config.yaml configurado para o gerenciamento de configuração.

### 1. Clonar o repositório

```
git clone <url_do_repositório>
cd <diretório_do_repositório>
```

### 2. Instalar as Dependências

Instale as dependências necessárias usando:

```
go mod tidy
```

### 3. Configurar o arquivo config.yaml

Crie um arquivo chamado config.yaml na raiz do projeto e adicione a seguinte configuração:

```
jwt:
  secret: "sua_chave_secreta_aqui"
```

### 4. Executar o Aplicativo

Inicie o servidor:

```
go run main.go
```

# Executando o Aplicativo com Docker

### 1. Construir a Imagem Docker

Para construir a imagem Docker do seu projeto, execute o seguinte comando na raiz do diretório do projeto (onde o Dockerfile está localizado):

```
docker build -t go-auth .
```

### 2. Executar o Contêiner Docker

Após a imagem ser construída com sucesso, você pode iniciar um contêiner a partir dela. Use o seguinte comando para rodar o contêiner e mapear a porta 8000 do contêiner para a porta 8000 da sua máquina local:

```
docker run -d -p 8000:8000 --name go-auth go-auth
```

### 3. Parar o Contêiner

Para parar o contêiner em execução, use o seguinte comando:

```
docker stop go-auth
```

### 4. Reiniciar o Contêiner

Para reiniciar o contêiner, utilize:

```
docker start go-auth
```

### 5. Remover o Contêiner

Se você precisar remover o contêiner (por exemplo, para recriá-lo), use:

```
docker rm go-auth
```

### 6. Acessar as Rotas da API

Depois de executar o contêiner, você pode acessar as rotas da API usando um cliente HTTP, como Postman ou Insomnia, ou até mesmo ferramentas como curl. A API estará disponível em:

Login: http://localhost:8000/login
Validar Token: http://localhost:8000/validate