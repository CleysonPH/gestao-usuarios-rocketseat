# Desafio - Criando uma API REST

## Introdução

Faaala Dev,

Nesse desafio você reforçará de forma prática os **conceitos** que aprendemos nesses módulos.

Como se trata de um desafio, ele necessita de alguns conhecimentos além dos abordados nesse módulo, então é importante ter autonomia para conseguir pesquisar essas coisas caso não saiba como resolver. Por isso, lembre-se, t**enha calma** e **acredite no seu processo.**

Além isso, este é um desafio prático e opcional, por tanto não há correção sobre ele.

O aprendizado daqui é muito importante e com certeza você conseguirá sair com muito conhecimento bacana 💜

## Sobre o desafio

Neste projeto você começará a construir uma API RESTful que irá realizar operações CRUD in-memory ***(cuidado com data-races!)***

## Instruções

Inicie um projeto (módulo) Go e implemente os seguintes endpoints:

| Método | URL            | Descrição                                                                                                   |
| ------ | -------------- | ----------------------------------------------------------------------------------------------------------- |
| POST   | /api/users     | Cria um usuário usando as informações enviadas no corpo da requisição.                                      |
| GET    | /api/users     | Retorna um array de usuários.                                                                               |
| GET    | /api/users/:id | Retorna o objeto do usuário com o id especificado.                                                          |
| DELETE | /api/users/:id | Remove o usuário com o id especificado e retorna o usuário deletado.                                        |
| PUT    | /api/users/:id | Atualiza o usuário com o id especificado usando dados do corpo da requisição. Retorna o usuário modificado. |

### Esquema do Usuário

Cada recurso de usuário deve seguir a seguinte estrutura (esquema):

```json
{
  "id": "",                     // UUID, required
  "first_name": "Jane Doe",     // String, required (len >= 2 && len <= 20)
  "last_name": "Jane Doe",      // String, required (len >= 2 && len <= 20)
  "biography": "Tendo diversão" // String, required (len >= 20 && len <= 450)
}
```

### “Banco de Dados” em memória

Como ainda não cobrimos como lidamos com banco de dados em Go iremos emular um banco de dados onde todo o conteúdo enviado por nossos usuários será feito num hash map, onde o id é a chave para o mesmo.

Exemplo:

```go
type id uuid.UUID

type user struct {
	FirstName string
	LastName  string
	biography string
}

type application struct {
	data map[id]user
}
```

Considerando isso, crie um pequeno pacote que implemente as seguintes funções para esse hash-map

- `FindAll` resolve para a lista de usuários (ou array vazio).
- `FindById` recebe um id e resolve para o usuário com esse id (ou null se o id não existir).
- `Insert` recebe um novo usuário {firstName, lastName, bio} e resolve para o usuário recém-criado {id, firstName, LastName, FullName, bio}.
- `Update` recebe um id e um usuário existente {firstName, lastName, bio} e resolve para o usuário atualizado {id, firstName, LastName, FullName, bio}. (Retorne um erro caso não exista.)
- `Delete` recebe um id e resolve para o usuário deletado {id, FirstName, LastName, FullName, bio}

### Especificações dos Endpoints

- Quando o cliente faz uma requisição POST para `/api/users`
    - Se o corpo da requisição estiver faltando a propriedade FullName, LastName ou bio:
        - Responda com o código de status HTTP 400 (Bad Request).
        - Retorne a seguinte resposta JSON: `{"message": "Please provide FirstName LastName and bio for the user"}`.
    - Se a informação sobre o usuário for válida:
        - Salve o novo usuário no banco de dados.
        - Responda com o código de status **HTTP 201 (Created)**.
        - Retorne o documento do usuário recém-criado, incluindo seu id.
    - Se houver um erro ao salvar o usuário:
        - Responda com o código de status **HTTP 500 (Server Error).**
        - Retorne o seguinte objeto JSON: `{"message": "There was an error while saving the user to the database"}`.
- Quando o cliente faz uma requisição GET para `/api/users`
    - Se houver um erro ao recuperar os usuários do banco de dados:
        - Responda com o código de status **HTTP 500**.
        - Retorne o seguinte objeto JSON: `{"message": "The users information could not be retrieved"}`.
- Quando o cliente faz uma requisição GET para `/api/users/:id`
    - Se o usuário com o id especificado não for encontrado:
        - Responda com o código de status **HTTP 404 (Not Found).**
        - Retorne o seguinte objeto JSON: `{"message": "The user with the specified ID does not exist"}`.
    - Se houver um erro ao recuperar o usuário do banco de dados:
        - Responda com o código de status **HTTP 500**.
        - Retorne o seguinte objeto JSON: `{"message": "The user information could not be retrieved"}`.
- Quando o cliente faz uma requisição DELETE para `/api/users/:id`
    - Se o usuário com o id especificado não for encontrado:
        - Responda com o código de status **HTTP 404 (Not Found).**
        - Retorne o seguinte objeto JSON: `{"message": "The user with the specified ID does not exist"}`.
    - Se houver um erro ao remover o usuário do banco de dados:
        - Responda com o código de status **HTTP 500.**
        - Retorne o seguinte objeto JSON: `{"message": "The user could not be removed"}`.
- Quando o cliente faz uma requisição PUT para `/api/users/:id`
    - Se o usuário com o id especificado não for encontrado:
        - Responda com o código de status **HTTP 404 (Not Found).**
        - Retorne o seguinte objeto JSON: `{"message": "The user with the specified ID does not exist"}`.
    - Se o corpo da requisição estiver faltando a propriedade name ou bio:
        - Responda com o código de status **HTTP 400 (Bad Request).**
        - Retorne a seguinte resposta JSON: `{"message": "Please provide name and bio for the user"}`.
    - Se houver um erro ao atualizar o usuário:
        - Responda com o código de status **HTTP 500.**
        - Retorne o seguinte objeto JSON: `{"message": "The user information could not be modified"}`.
    - Se o usuário for encontrado e a nova informação for válida:
        - Atualize o documento do usuário no banco de dados usando a nova informação enviada no corpo da requisição.
        - Responda com o código de status **HTTP 200 (OK).**
        - Retorne o documento do usuário recém-atualizado.

### Notas Importantes

- Teste seu trabalho manualmente usando Postman, Insomnia ou HTTPie.
- Em sua solução, é essencial que você siga as melhores práticas e produza resultados limpos e profissionais.

### Um passo além

Tome cuidado para não quebrar o MVP ao trabalhar nessas metas avançadas! Em caso de dúvida, crie um novo branch.

- Utilize de mutexes para ter certeza de que não haverá condições de corrida em nosso código