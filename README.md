# 🧾 Billing Service Demo

Welcome to the **Billing Service Demo**. This project is designed to showcase technical expertise and system design.
We've integrated two powerful technologies:

- [Encore](https://github.com/encoredev/encore)
- [Temporal](https://temporal.io/) workflow engine.

## 🛠 Development Setup

To dive into this demo locally:

1. **Prerequisites**: Ensure you have `docker` and `docker-compose` set up on your machine.
  
2. We're utilizing **Temporalite** for a local Temporal server experience.

### 🌀 Launching Temporalite Server

1. 🔗 Grab the latest Temporalite release from [GitHub](https://github.com/temporalio/temporalite/releases/tag/v0.3.0).
  
2. 🚀 Start it up with the command:
   ```bash
   temporalite start --namespace default
   ```

3. 🌍 To access the Temporal dashboard, navigate to:
   ```
   http://127.0.0.1:8233
   ```

### 🎵 Working with Encore

For local development, ensure you've got [Encore](https://encore.dev/docs) ready.

1. First of all run test via Encore:
   ```bash
   encore test ./...
   ```

2. Run the following to kick off Encore:
   ```bash
   encore run
   ```

3. This spins up the Encore UI in your default browser at:
   ```
   http://localhost:9400
   ```

   Use the UI to monitor and analyze API calls seamlessly.
