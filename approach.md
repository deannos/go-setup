**VERBALIZE THIS:**
"Before coding, let me clarify requirements and propose clean architecture approach:

1. Entity (Domain Model)
2. Repository Interface + Implementation
3. UseCase (Business Logic)
4. Handler (HTTP Layer)
5. Tests + Validation

This follows SOLID principles and enables independent testing/scaling.

Questions:

- Input validation requirements?
- Pagination/filtering needed?
- Error response format?
- Database constraints?"

**Talking Points**

**Architecture:** "5-layer clean architecture: Entity→Repo→UseCase→Handler→HTTP"
**Testability:** "Repository interfaces enable 100% unit test coverage without DB"
**Scalability:** "Connection pooling + context timeouts ready for horizontal scale"
**Production:** "Graceful shutdown + health checks = Kubernetes-ready"
**Tradeoffs:** "Direct entity binding = fast iteration vs DTOs = strict contracts"



## 🔍 **Tech Stack**

| Layer | Technology |
|-------|------------|
| Framework | Gin Gonic v1.9 |
| Database | PostgreSQL 15 |
| Testing | Testify + Mock |
| Config | godotenv |
| Logging | Structured Zap |

## 🎯 **For Interviews**

**90-minute coding round strategy:**
1. **0-5m**: Clarify requirements + architecture
2. **5-15m**: Copy boilerplate + `make run`
3. **15-70m**: Entity → Repository → UseCase → Handler
4. **70-85m**: Unit tests (green immediately)
5. **85-90m**: Live demo + tradeoffs discussion

## 📈 **Production Scale**

- **Connection Pooling**: 25 max connections
- **Graceful Shutdown**: Zero data loss
- **Health Checks**: Kubernetes/Load Balancer ready
- **Context Timeouts**: DDoS protection
- **API Versioning**: `/api/v1` future-proof

