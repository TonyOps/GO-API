
## Visão Geral
API RESTful desenvolvida em Go para gerenciamento de produtos e usuários, atualmente em produção. Implementa práticas recomendadas para projetos Go em ambientes de produção, priorizando clareza, manutenibilidade e escalabilidade. 

## Tecnologias Utilizadas
- **Linguagem**: Go (Golang)
- **Web Framework**: go-chi/chi (com JWT authentication)
- **Banco de Dados**: MySQL
- **Criptografia**: bcrypt para senhas
- **Geração de IDs**: UUID
- **Gerenciamento de Configuração**: Viper

## Estrutura do Projeto
```
cmd/
  └── server/          # Ponto de entrada da aplicação
configs/               # Configurações do sistema
internal/
  └── entity/          # Modelos de domínio e lógica de negócios
pkg/
  └── entity/          # Componentes compartilhados (IDs UUID)
```
Esta estrutura segue as melhores práticas para organização de código Go, facilitando a manutenção e escalabilidade do projeto. 

## Configuração

### Variáveis de Ambiente
Crie um arquivo `.env` na pasta `cmd/server` com as seguintes configurações:

```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=tony
WEB_SERVER_PORT=8080
JWT_SECRET=secret
JWT_EXPIRESIN=300
```

**Importante**: Em ambiente de produção, utilize valores seguros para `JWT_SECRET` e credenciais do banco de dados. 

### Dependências
O projeto utiliza Go Modules para gerenciamento de dependências. Para instalar:
```bash
go mod tidy
```
Mantenha as dependências atualizadas e realize auditorias regulares para garantir a segurança do projeto. 

## Executando a Aplicação

1. Certifique-se de ter o MySQL rodando com o banco de dados `tony` configurado
2. Instale as dependências:
   ```bash
   go mod download
   ```
3. Execute o servidor:
   ```bash
   cd cmd/server
   go run main.go
   ```
O servidor será iniciado na porta configurada (padrão: 8080). 

## Funcionalidades da API

### Autenticação
- Sistema de autenticação baseado em JWT (HS256)
- Validação de credenciais com hash bcrypt
- Token com expiração configurável via `.env` (padrão: 300 segundos)

### Gerenciamento de Produtos
- Validação rigorosa de dados:
  - Nome obrigatório
  - Preço deve ser maior que zero
  - ID gerado automaticamente com UUID
- Estrutura completa com descrição, preço e data de criação

### Gerenciamento de Usuários
- Cadastro com username, email e senha
- Armazenamento seguro de senhas com bcrypt
- Validação de credenciais através do método `ValidatePassword`

## Testes
O projeto inclui testes unitários abrangentes para garantir a integridade das entidades:

```bash
# Testar entidade Product
go test ./internal/entity/product_test.go

# Testar entidade User
go test ./internal/entity/user_test.go
```
Recomenda-se automatizar os testes e considerar o uso de mock servers para testes de integração. 

## Melhores Práticas Implementadas
- Validação rigorosa de dados nas entidades
- Separação clara de responsabilidades na estrutura do projeto
- Configuração centralizada com Viper
- Uso de UUIDs para identificadores únicos
- Documentação clara e legível da API 
