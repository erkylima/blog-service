Microblog: Serviço RESTful em GoLang - Domínios: Pages, Posts, Comments, Categories
================================================================================

Este é um repositório para um serviço RESTful escrito em GoLang, implementando os domínios de Pages, Posts, Comments e Categories. A arquitetura de aplicação utilizada é a hexagonal (também conhecida como arquitetura hexagonal ou ports and adapters), que promove um design limpo e modularizado, facilitando a manutenção e evolução do sistema.

Pré-requisitos
--------------

*   [Docker](https://www.docker.com/) e Docker Compose instalados.

Configuração
------------

1.  Clone este repositório:

```git clone https://github.com/erkylima/blog-service.git```

1.  Na raiz do projeto, execute o seguinte comando para iniciar o MongoDB utilizando o Docker Compose:


```docker-compose up -d```

1.  Agora você pode compilar e executar o serviço GoLang.

Estrutura do Projeto
--------------------

*   `cmd/`: Contém o código de inicialização da aplicação.
*   `internal/`: Contém o código-fonte da aplicação organizado por módulos.
    *   `pages/`: Módulo para operações relacionadas às páginas.
    *   `posts/`: Módulo para operações relacionadas aos posts.
    *   `comments/`: Módulo para operações relacionadas aos comentários.
    *   `categories/`: Módulo para operações relacionadas às categorias.
    *   `shared/`: Módulo para operações compartilhadas entre outros módulos
*   `pkg/`: Contém pacotes compartilhados entre os módulos.
*   `docker-compose.yml`: Arquivo para iniciar o MongoDB.

Como Contribuir
---------------

1.  Fork este repositório.
2.  Crie uma branch com uma descrição significativa da mudança: `git checkout -b feature/nova-feature`.
3.  Desenvolva sua feature ou correção de bug.
4.  Certifique-se de que os testes passam: `go test ./...`.
5.  Faça o commit de suas mudanças: `git commit -m "Adicionando nova feature"`.
6.  Faça push para a sua branch: `git push origin feature/nova-feature`.
7.  Crie um novo Pull Request explicando suas mudanças.

Contato
-------

Para sugestões, dúvidas ou problemas, sinta-se à vontade para entrar em contato:

Érky Lima

Licença
-------

Este projeto está licenciado sob a Licença MIT.