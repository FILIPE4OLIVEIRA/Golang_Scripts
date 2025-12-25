# Numerical Analysis (Golang)

Repositório com implementações simples de métodos numéricos em Go: busca de raízes, integração numérica, integrais múltiplas por Monte Carlo e resolução de EDOs (Euler e Runge-Kutta).

**Estrutura principal**

- `cmd/` — entrada `main.go` com exemplos de uso.
- `internal/rootfinding/` — Bisection, Secant, Newton-Raphson.
- `internal/integration/` — Trapezoidal, Simpson composite, Midpoint, double/triple integrals, área entre curvas.
- `internal/ode/` — Euler, Runge-Kutta 4 (1ª e 2ª ordem).
- `internal/suport/` — funções auxiliares (derivada numérica, média, etc.).

**Pré-requisitos**

- Go 1.18+ instalado

**Como executar (rápido)**

No terminal (na raiz do projeto):

```bash
go run cmd/main.go
```

Ou compilar e executar:

```bash
go build -o bin/numerical_analysis ./cmd
# Windows
.\bin\numerical_analysis.exe
# Linux / macOS
./bin/numerical_analysis
```

**O que o `cmd/main.go` faz**

- Executa exemplos de: Bisection, Secant, Newton-Raphson para uma função polinomial;
- Integração de `sin(x)` por Trapézio, Simpson e Midpoint;
- Cálculo da área entre `sin` e `cos` em `[0, π]`;
- Resolve uma EDO simples com Euler e Runge-Kutta 4; demonstração de RK4 para EDO de 2ª ordem.

> Observação: as funções de integral dupla/tripla usam Monte Carlo com grandes amostras — podem demorar.

**Editar/Executar exemplos**

- Abra `cmd/main.go` e adapte as funções/fhechos de exemplo conforme precisar.

**Atualizar repositório remoto (se renomeou localmente)**

Se você renomeou a pasta localmente mas também renomeou o repositório no GitHub, atualize o `origin`:

```bash
# ver URL atual
git remote get-url origin

# exemplo HTTPS
git remote set-url origin https://github.com/SEU_USUARIO/Numerical_Analysis.git

# exemplo SSH
git remote set-url origin git@github.com:SEU_USUARIO/Numerical_Analysis.git

# depois
git add README.md
git commit -m "Add README"
git push origin --all
```

**Contribuição**

- Faça fork, crie branch com feature/bugfix, abra PR com descrição.

**Licença**

Use a licença que preferir (ex.: MIT). Adicione um arquivo `LICENSE` se quiser publicar.

---
