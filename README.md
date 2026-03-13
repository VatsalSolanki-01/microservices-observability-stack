# Microservices Observability Stack

A hands-on DevOps project that demonstrates how monitoring, logging, dashboards, and alerting work together to provide visibility into distributed microservices systems.

---

# What This Project Does

This project builds a complete observability environment for a microservices based system using open source monitoring tools.

Application servers run microservices along with monitoring agents that collect metrics and logs.  
These signals are then sent to a centralized monitoring server where they are stored, visualized, and used for alerting.

This setup simulates how real world DevOps and SRE teams monitor production systems.

---

# How It Helps / Problems It Solves

Modern applications are often built using distributed microservices. When failures occur, it becomes difficult to understand what is happening across multiple services and infrastructure components.

My project improves visibility into the system by enabling:

• Infrastructure monitoring for CPU, memory, disk, and network usage  
• Application performance monitoring including request rate and latency  
• Centralized log aggregation for easier debugging  
• Real time service health monitoring  
• Automated alerting when services become unavailable  

This type of observability setup helps engineers quickly detect issues and respond before they impact users.

---

# Architecture Diagram

               +----------------------------+
               |     Application Server     |
               |----------------------------|
               |  Microservices             |
               |  Node Exporter             |
               |  Promtail                  |
               +-------------+--------------+
                             |
                             | Metrics & Logs
                             v
               +----------------------------+
               |      Monitoring Server     |
               |----------------------------|
               |  Prometheus                |
               |  Loki                      |
               |  Grafana                   |
               |  Alertmanager              |
               +------+-----------+---------+
                      |           |
                      |           |
                      v           v
             Grafana Dashboards   Alert Triggered
                                      |
                                      v
                               Email Notification


### Flow

1. **Node Exporter** collects system metrics from service nodes.  
2. **Prometheus** scrapes and stores metrics.  
3. **Promtail** collects logs from services.  
4. **Loki** aggregates logs centrally.  
5. **Grafana** visualizes metrics and logs through dashboards.  
6. **Alertmanager** sends alerts when failures occur.

---

# Tech Stack

| Tool | Logo |
|-----|-----|
| Docker | <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" width="40"/> |
| Prometheus | <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/prometheus/prometheus-original.svg" width="40"/> |
| Loki | <img src="https://raw.githubusercontent.com/grafana/loki/main/docs/sources/logo.png" width="40"/> |
| Node Exporter | <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/prometheus/prometheus-original.svg" width="40"/> |
| Grafana | <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/grafana/grafana-original.svg" width="40"/> |
| Alertmanager | <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/prometheus/prometheus-original.svg" width="40"/> |
| Promtail | <img src="https://raw.githubusercontent.com/grafana/loki/main/docs/sources/logo.png" width="40"/> |
| Docker Compose | <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" width="40"/> |

---

Turning distributed chaos into clean dashboards
