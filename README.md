# Billing service demo
This is a demo project for the billing to demonstrate tech and system design skills.
Service based on [Encore](https://github.com/encoredev/encore)
and [Temporal](https://temporal.io/) workflow engine.

## Development
To run the project locally, you need to have docker and docker-compose installed.
This demo uses Temporalite as locally hosted temporal server.

### Run temporalite server
You can find latest release of the temporalite server on the [GitHub](https://github.com/temporalio/temporalite/releases/tag/v0.3.0).
And then just run it:
```bash
$ temporalite start --namespace default
```
