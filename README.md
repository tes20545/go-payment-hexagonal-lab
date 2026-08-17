## 💰 Go Payment Hexagonal Lab (Phase 1)
An advanced wallet system development project prioritizing data integrity and high performance, utilizing the Go language, Hexagonal Architecture, and infrastructure management via Terraform.

### 🎯 Project Goals (Phase 1)
- **Data Integrity**: verify right transactions 100% with ACID Transactions (Double-entry bookkeeping)

- **High Concurrency**: support concurrency transactions without Data Race condition

- **Performance**: design with consider about Memory Allocation and GC Overhead (Lead Implementation)

- **Infrastructure as Code**: manage development environment via Terraform and Docker

- **Reliability**: develop by not dependent on thrid-party libs as much as possible

### 🏗️ Architecture Stack
- **Language**: Go 1.2x (Focus on Pointer optimization & Generics)

- **Architecture**: Hexagonal (Ports & Adapters)

- **Database**: MySQL 8.0 (Inno DB Engine for ACID)

- **Infrastructure**: Terraform (Docker Provider)

- **Messaging** (Roadmap): Kafka for Event-driven processing

### 📅 Phase 1 Roadmap (Week 1 - 4)

#### **Week 1:** Foundations & Infra
[ ] Setup Infrastructure via Terraform (MySQL 8.0 Container)

[ ] Implement Hexagonal Layout (Domain, Port, Service, Adapter)

[ ] Database Schema Design (Wallets & Transactions tables)

[ ] Challenge: Implement Double-entry bookkeeping logic

#### **Week 2**: Core Logic & Concurrency
[ ] Develop Deposit & Transfer Use Cases

[ ] Handle DB Transactions with sql.Tx in Go

[ ] Challenge: Perform Load Test with 1,000+ Goroutines to verify Race-free logic

#### **Week 3**: Containerization & Optimization
[ ] Multi-stage Docker Build (Distroless image for security & size)

[ ] Setup Docker Compose for local orchestration

[ ] Challenge: Achieve Zero-Allocation in critical path logic (Benchmarking)

#### **Week 4**: Observability & IaC Advanced
[ ] Implement Terraform Network Isolation

[ ] Structured Logging & Custom Error Handling (Domain Errors)

[ ] Challenge: Automated Benchmark Report for GitHub Actions