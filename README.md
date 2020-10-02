# Teste de Software - Assignment 1

###### Otávio Celani

## Projeto

Projeto destinado à disciplina de Teste de Software, lencionada pelo professor Lesandro Ponciano, na PUC Minas.

### Diretórios

- **pkg:** contém o código do projeto.
- **bin:** arquivos binários compilados para sistemas unix.
- **tests-output:** arquivos com registros dos outputs dos testes.
- **img:** imagens e prints do experimento.
- - **grapviz:** grafos dos testes gerados em '.svg' com a ferramenta graphviz.
- - **prints:** prints tirados do experimento em execução, código e testes.
- - **draws:** imagens dos desenhos realizados sobre o grafo e análise de complexidade ciclomática.

### Execução do código

Opções:

- Em sistemas unix, é possível executar os binários na pasta 'bin', por exemplo:

```bash
$ ./bin/main
```

- Para executar o código diretamente, é preciso instalar a runtime da ferramaneta Golang. A seguir, acessar a pasta 'pkg' e executar o comando:

```bash
$ go run ./pkg/main.go
```

- Para executar os testes diretamente, é preciso instalar a runtime da ferrmaneta Golang. A seguir, acessar a pasta 'pkg' e executar o comando:

```bash
$ go test ./pkg
```

## Instruções do teste fornecidas pelo professor

### O que o aluno deve fazer:

    Em uma mesma classe, implementar um método que soluciona o problema apresentado;
    Derivar o Grafo de Fluxo de Controle do método;
    Determinar a complexidade ciclomática do grafo;
    Determinar um conjunto base de caminhos linearmente independentes a partir do grafo e do valor de complexidade ciclomática;
    Definir os casos de teste que vão forçar a execução de cada caminho do conjunto, indicando as entradas e respectivas saídas esperadas;
    Implementar e executar uma suíte de testes, que contemple os casos de teste identificados. É fundamental implementar os asserts que verificam se a saída gerada pelo método é igual à saída esperada.

### O que o aluno deve entregar:

    No texto da resposta, o aluno deve incluir: print da implementação, imagem do grafo, memória de cálculo da complexidade ciclomática, descrição do conjunto de caminhos; descrição do conjunto de casos de teste e print da implementação do conjunto de testes.

### Apresentação oral:

    Na aula do dia 02/10/2020, o aluno deve estar preparado para apresentar e responder perguntas sobre a entrega feita, incluindo a demonstração da implementação feita.

---

#### Desafio

É necessário calcular a média dos valores que estão na interseção entre dois conjuntos.

Tem-se dois vetores de tamanho m e n, respectivamente.
Dados os vetores, deve-se verificar os valores presentes nos dois vetores (a interseção) e calcular a média aritmética desses valores que estão na interseção.

##### Por exemplo:

- se vetor1=[19, 5, 2, 6] e vetor2=[5, 0, 9, 4, 18, 56], o valor resultante deve ser 5.

- se vetor1=[1, 3, 2, 6] e vetor2=[7, 0, 9, 4, 3, 1], o valor resultante deve ser 2.

Nesse caso, já está validado que cada vetor é um conjunto, ou seja, ele não possui valores repetidos.

##### O método deve:

- se chamar mediaIntersecao;
- receber como parâmetro m e n, vetor1 e vetor2;
- retornar o valor float que é da média dos valores na interseção entre os dois conjuntos;
- retornar o valor 0 se não houver interseção.
