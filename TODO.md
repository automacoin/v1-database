## TODO

* Flag Parser, Config Center.
* General logger package -> STDOUT output
* Travis CI
* Gitter Channel
* Version Release in Github
* Test Coverage Report
* Benchmarks
* Catch SIGINT
* Port to be set up as flag
* Test for API logger
* Test for API crash and recovery
* Add sentry for errors / exceptions
* GET /test-auth to be disabled in production

## 12 factor app manifest check

* I. Codebase: Flag parser, config center (TODO)
* II. Dependencies: Done with `go mod`.
* III. Config: See Flag parser
* IV. Backing services: See Flag parser
* V. Build, release, run: Done with Makefile
* VI. Processes: Complies, Backend Database to run in a stateful service.
* VII. Port binding: See Flag parser, all services to be accessed via REST API
* VIII. Concurrency: Complies. State lives in the backing database.
* IX. Disposability: Startup: Conn to Backing database, router prep. TODO Graceful shutdown
* X. Dev/prod parity: Complies: Variations are specified by command option flags.
* XI. Logs: TODO. Logs to STDOUT.
* XII. Admin processes: Complies. This service can be seen as a middleware to the backend database.
