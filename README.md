# On-Board-Unit-OBU-microservice
a vehicle-installed electronic device that collects and transmits travel, location, and vehicle data to central servers or roadside infrastructure. It is most commonly used in commercial trucking for automatic electronic toll collection and fleet telematics.

```mermaid
flowchart LR
    A[🚗 OBU Sender] -->|WebSocket| B[📡 Receiver] -->|Kafka| C[⚙️ Consumer] --> D[🧮 Calculator]
    C --> E[(🗄️ Database)]
```
```mermaid
sequenceDiagram
    autonumber
    participant V as 🚗 Vehicle (OBU)
    participant R as 📡 OBU Receiver<br/>(WebSocket)
    participant K as 📨 Kafka
    participant C as ⚙️ Consumer
    participant CAL as 🧮 Calculator Service
    participant DB as 🗄️ Database

    loop Continuously
        V->>R: WebSocket message<br/>(location, speed, etc.)
        R->>K: Produce to topic
    end

    loop Stream processing
        K->>C: Consume message
        C->>CAL: Send data for calculation
        C->>DB: Store raw data
        CAL-->>C: Result (optional)
    end
```