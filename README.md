# Grafana & Prometheus Monitoring with Go Application

Este projeto visa demonstrar o uso do **Prometheus** para coletar métricas e o **Grafana** para visualização de dados, utilizando um aplicativo simples desenvolvido em **Go**. O aplicativo expõe métricas como o número de usuários online e a duração das requisições HTTP, que são coletadas e monitoradas pelo Prometheus. Essas métricas são então visualizadas no Grafana.

## 🚀 Funcionalidades do Projeto

- **Go Application**:
  - **Contador de Requisições HTTP**: Conta o número total de requisições HTTP feitas ao aplicativo.
  - **Métrica de Usuários Online**: Simula a quantidade de usuários online, utilizando uma métrica do tipo *Gauge*.
  - **Histograma de Duração de Requisições HTTP**: Mede o tempo de duração das requisições HTTP para análise de performance.
  
- **Prometheus**:
  - Exposição de métricas na rota `/metrics`, que pode ser monitorada pelo Prometheus.
  - Configuração do Prometheus para coleta de métricas da aplicação Go.

- **Grafana**:
  - Visualização das métricas do Prometheus.
  - Monitoramento de dados de requisições e usuários online.
  
## 🛠️ Tecnologias Utilizadas

- **Go**: Linguagem de programação para construir a aplicação.
- **Prometheus**: Sistema de monitoramento e coleta de métricas.
- **Grafana**: Plataforma de visualização de métricas.
- **Docker**: Containerização do aplicativo e serviços.
- **Docker Compose**: Gerenciamento de múltiplos contêineres Docker.

## 🧑‍💻 Como Rodar o Projeto

### 1. Clone este repositório
```bash
git clone https://github.com/Gardene10/Grafana_Prometheus.git
cd Grafana_Prometheus

2. Construa e rode o contêiner Docker
bash
Copy
Edit
docker-compose up --build
3. Acesse o aplicativo
O aplicativo estará disponível em:

Go Application: http://localhost:8181

Prometheus Metrics: http://localhost:8181/metrics

4. Acesse o Grafana
Grafana: http://localhost:3000

Usuário padrão: admin

Senha padrão: admin

📊 Métricas
goapp_online_users: Número de usuários online simulados no sistema.

goapp_http_requests_total: Contagem total de requisições HTTP feitas ao servidor.

goapp_http_request_duration: Tempo de duração das requisições HTTP (histograma).
