# TreeCLI

TreeCLI é uma ferramenta de linha de comando escrita em Go que gera uma representação em árvore da estrutura de diretórios. Suporta várias opções para personalizar a saída, incluindo exclusão de padrões, limitação de profundidade e filtragem por extensões de arquivo.

## **Instalação**

1. **Clone o repositório:**

```bash
   git clone https://github.com/seu-usuario/treecli.git
   cd treecli
```

2. **Compilar o programa:**

```bash
   go build -o treecli cmd/treecli/main.go
```

Ou usando o Makefile:

```bash
   make build
```

## **Uso**

```bash
    ./treecli [opções] [caminho]
```

- Se nenhum caminho for especificado, o diretório atual (.) será usado.
- As opções disponíveis estão listadas abaixo.

## **Opções**

- -h, --help: Exibe a mensagem de ajuda.
- -exclude: Padrões de exclusão (wildcards), separados por vírgula. Exemplo: -exclude='node_modules,vendor,*.go'
- -max-depth: Limite de profundidade da árvore (0 para ilimitado). Exemplo: -max-depth=3
- -ext: Lista de extensões de arquivo a serem incluídas, separadas por vírgula. Exemplo: -ext='.go,.md'
- -exclude-ext: Lista de extensões de arquivo a serem excluídas, separadas por vírgula. Exemplo: -exclude-ext='.txt,.log'

## **Exemplos**

1. **Exibir a estrutura completa do diretório atual:**

```bash
    ./treecli
```

2. **Excluir os diretórios node_modules, vendor e tudo que termina com ".go"**

```bash
    ./treecli -exclude='node_modules,vendor,*.go'
```

3. **Limitar a profundidade a 2 níveis:**

```bash
    ./treecli -max-depth=2
```

4. **Incluir apenas arquivos com extensões .go e .md:**

```bash
    ./treecli -ext='.go,.md'
```

5. **Combinar opções:**

```bash
    ./treecli -exclude -max-depth=2 --ext='.go,.md' ~/projetos
```

## **Makefile**

Você pode usar o Makefile para facilitar a compilação e limpeza do projeto.

- Compilar o programa:

```bash
    make build
```

- Limpar arquivos compilados:

```bash
    make clean
```