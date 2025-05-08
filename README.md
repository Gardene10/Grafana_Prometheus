# 📊 Projeto de Monitoramento com Prometheus, Grafana e Go

Este projeto tem como objetivo demonstrar a criação e visualização de métricas personalizadas utilizando Go, Prometheus, Grafana e Docker. Ele foi desenvolvido com foco em aprendizado e prática de técnicas de observabilidade em aplicações modernas.

---

## 🚀 Tecnologias Utilizadas

- **Go (Golang)** – para criação da aplicação e exposição das métricas
- **Prometheus** – para coleta das métricas expostas pela aplicação
- **Grafana** – para visualização gráfica das métricas
- **Docker** e **Docker Compose** – para orquestração dos serviços
- **cAdvisor** – para monitoramento de métricas dos containers Docker

---

## 📦 Como Executar o Projeto

1. Clone o repositório:
   ```bash
   git clone https://github.com/Gardene10/Grafana_Prometheus.git
   cd Grafana_Prometheus

Suba os serviços com Docker Compose:

bash
Copy
Edit
docker-compose up -d
Acesse os serviços:

Grafana: http://localhost:3000 (login padrão: admin / admin)

Prometheus: http://localhost:9090

App Go (endpoint de métricas): http://localhost:8181/metrics

⚠️ Para iniciar o app Go, entre no contêiner com:

bash
Copy
Edit
docker exec -it app bash
go run main.go
📈 Métricas Implementadas
✅ Contador de Requisições HTTP
Conta o número total de requisições HTTP feitas à aplicação. Útil para acompanhar o volume de uso.

✅ Métrica de Usuários Online
Simula a quantidade de usuários online utilizando uma métrica do tipo Gauge, atualizada dinamicamente.

✅ Histograma de Duração das Requisições HTTP
Mede o tempo de duração de cada requisição HTTP, com buckets para análise de performance.

📊 Dashboards no Grafana
Painéis específicos foram criados para cada métrica, permitindo:

Monitoramento em tempo real

Análise histórica

Visualização clara com gráficos personalizados

📁 Estrutura do Projeto

├── docker-compose.yml
├── Dockerfile
├── prometheus.yml
├── main.go
└── README.md

🎯 Objetivo
Este projeto faz parte de uma trilha prática para aprimorar conhecimentos em observabilidade e monitoramento de aplicações.


---


# 📊 Monitoring Project with Prometheus, Grafana and Go

This project demonstrates how to create and visualize custom metrics using Go, Prometheus, Grafana, and Docker. It was developed as a way to practice observability techniques for modern applications.

---

## 🚀 Technologies Used

- **Go (Golang)** – to build the application and expose metrics
- **Prometheus** – to collect the metrics exposed by the app
- **Grafana** – to visually display collected metrics
- **Docker** and **Docker Compose** – for easy orchestration
- **cAdvisor** – for container-level monitoring

---

## 📦 How to Run the Project

1. Clone the repository:
   ```bash
   git clone https://github.com/Gardene10/Grafana_Prometheus.git
   cd Grafana_Prometheus

Start the services using Docker Compose:

bash
Copy
Edit
docker-compose up -d
Access the services:

Grafana: http://localhost:3000 (default login: admin / admin)

Prometheus: http://localhost:9090

Go App (metrics endpoint): http://localhost:8181/metrics

⚠️ To start the Go app, enter the container and run the command:

bash
Copy
Edit
docker exec -it app bash
go run main.go
📈 Implemented Metrics
✅ HTTP Requests Counter
Tracks the total number of HTTP requests made to the app. Useful for monitoring usage volume.

✅ Online Users Gauge
Simulates the number of online users using a Gauge type metric, updated dynamically at /metrics.

✅ HTTP Request Duration Histogram
Measures how long each HTTP request takes, providing performance insights with histogram buckets.

📊 Grafana Dashboards
A custom panel was created for each metric, allowing:

Real-time monitoring

Historical analysis

Clear visualization with tailored charts

📁 Project Structure

├── docker-compose.yml
├── Dockerfile
├── prometheus.yml
├── main.go
└── README.md

🎯 Objective
This project is part of a hands-on learning journey to strengthen observability and monitoring skills.

