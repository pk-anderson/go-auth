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