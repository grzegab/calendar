# Calendar

This app (REST API) is a calendar for booking lessons. For the frontend view a separate
 project is created in TypeScript with React (Next.js).
Calendar view is made in plain HTML + JS to show one previous month and 5 months in the future.

The main advantage is to show calendar and book hours.
Logged users can get notifications instantly when lessons are accepted or declined.
Additionally, this app is blocking IPs for brute-force attacks or scanning for vulnerabilities.

This is a work in progress, so it isn't functional yet. Things to be completed before release:
- [X] Write README
- [ ] Event storming images
- [ ] C4 diagrams
- [X] Scheduling module
- [X] Booking module
- [X] Users module
- [X] Graceful shutdown
- [ ] Logs for application
- [X] Add health check
- [ ] Recoverer for HTTP requests (middleware)
- [X] Create a config mechanism (read conf files)
- [ ] Login mechanism (JWT tokens)
- [X] Dockerize
- [ ] Makefile for simpler instructions
- [ ] PostgreSQL database setup
- [ ] WebSockets setup
- [ ] SMS, Email, Slack notifiers setup
- [ ] Event BUS
- [ ] Redis setup
- [ ] IP check (with ban users)
- [ ] PostgreSQL migrations
- [ ] Unit tests
- [ ] Integration tests (Database, other services)
- [ ] E2E tests (Happy path, other cases)
- [ ] Profiling with pprof
- [ ] Add GitHub pipeline

## Running app

First create own config file in `config/config.yaml` and fill it with your own settings. 
Each option not listed in config file will be taken from default file.
Then run `make run` to start the application.

### Main flow

Teachers are added by admin. Students can register themselfs but need to be activated by teacher.

Teacher add slots of time availble for booking. Those slots are set to be maxium 8 hours per day and max 2 hour per slot.
Students can book those slots. Student choose slot to book it. Multiple students can book same slot. Students are alowed to use max 30 hours of slots.
Teacher select student that is interesrd in using slot time. Teacher can also decilne all students from slot.

User will register with phone number and email. To login code will be sent by SMS or email.
After logging in, user can book time slots. If time slot is accepted user will receive notification by websocket.

##  REST API with DDD

This is a REST API application for booking lessons. 
HTTP requests are made to rest endpoints and then translated into DDD style.
This is a modular monolith system architecture. Each module has its own architecture (for the moment it is hexagonal architecture).
There is one DB (PostgreSQL) and one Redis instance.

The application follows Domain-Driven Design (DDD) principles, where each module encapsulates its own domain logic:

**Domain Layer** - Contains the core business logic, entities, value objects, and domain services. Each module (Users,
Scheduling, Bookings) has its own domain package with business rules and invariants. For example, the Users module
contains User entity with Email and PhoneNumber value objects that enforce validation rules.

**Application Layer** - Orchestrates domain objects to fulfill use cases. Application services (handlers) coordinate
domain operations without containing business logic themselves. They work with domain repositories and enforce
transaction boundaries.

**Infrastructure Layer** - Provides technical implementations for domain interfaces (repositories). For instance,
PostgreSQL repositories implement domain repository interfaces, allowing the domain to remain independent of
infrastructure concerns.

**Interfaces Layer** - Adapts external requests (HTTP) to application layer commands and queries. HTTP handlers
translate REST requests into domain operations and format responses.

Each module maintains its own bounded context with clearly defined domain models and repository interfaces. This modular
approach ensures loose coupling between modules while maintaining strong cohesion within each domain.

This application implements the CQRS (Command Query Responsibility Segregation) pattern, which separates read and write
operations:

**Commands** - Handle write operations (create, update, delete). Commands work with domain objects and enforce business
rules. They modify the state of the system and go through the domain layer to ensure consistency and validation.

**Queries** - Handle read operations. Queries bypass the domain layer and work directly with the database using
optimized views or read models. This approach provides better performance for read operations and allows for different
data models optimized for specific query needs.

This separation allows each side to be optimized independently - commands for consistency and business logic, queries
for performance and flexibility.

### Ports and Adapters (Hexagonal Architecture)

Each module follows the Hexagonal Architecture pattern (also known as Ports and Adapters), which separates the core
business logic from external concerns:

**Ports** - Interfaces that define how the application communicates with the outside world. There are two types:

- *Inbound Ports (Driving Ports)* - Define use cases that drive the application (e.g., application service handlers like
  `confirm_schedule.Handler`, `new_timeslot.Handler`). These are located in the `application/` directory and represent
  operations that can be performed on the domain.
- *Outbound Ports (Driven Ports)* - Define interfaces for external dependencies that the application needs (e.g.,
  `domain.Repository`, `domain.SchedulingRepository`). These are typically defined in the `domain/` package as
  interfaces.

**Adapters** - Concrete implementations that connect ports to external systems:

- *Inbound Adapters (Primary/Driving Adapters)* - Translate external requests into application use cases. HTTP handlers
  in `interfaces/http/` are examples - they receive REST requests, validate input, call appropriate application
  handlers, and format responses.
- *Outbound Adapters (Secondary/Driven Adapters)* - Implement outbound ports to connect with external systems. Database
  repositories in `infrastructure/postgres/` implement domain repository interfaces, translating domain operations into
  SQL queries.

For example, in the Users module:

- The `domain.Repository` interface is an outbound port defining what persistence operations the domain needs
- The `postgres.UserRepository` is an outbound adapter implementing this interface with PostgreSQL
- The `http.RegisterUserHandler` is an inbound adapter receiving HTTP requests
- The `application.RegisterUserService` is an inbound port (use case) orchestrating the registration flow

This architecture ensures that the domain layer remains independent of technical details, making the code more testable,
maintainable, and allowing easy swapping of infrastructure components without affecting business logic.

### Directory Structure

The project follows a modular monolith architecture with clear separation of concerns:
- cmd/api - entry point for the application
- internal/app - business logic
- internal/shared - shared code across modules (infrastructure, common utilities)
- internal/pkg - shared packages

Module source:
 - adapters/ - interfaces to external services (database,
 - ports/ - interfaces to modules (use cases)
 - application/ - services for application
 - domain/ – business logic
 - infrastructure/ - repositories for external services
 - interfaces/ - module entry points
 - module.go - module initialization main file

### Modules description

Application is divided into 3 modules. This segregation became clear after event storming. 
- User (settings, other)
- Scheduling (Teacher slots)
- Bookings (Student slots)

#### User module

Is responsible for managing users. User can register with phone number and email but stays inactive until teacher activate account.
After activation user can login with email or phone number. When logging with phone number code will be sent to email and vice versa.
after providing correct code user can login (get JWT token). Teacher account is created with admin role.

#### Scheduling module

Is responsible for managing slots of time for teachers. It manage:
- show all teacher slots (filters available)
- show slot info
- add slot
- confirm student slot
- cancel student slot

#### Booking module

Is responsible for managing slots of time for teachers. It manage:
- show all teacher slots (filters available)
- show slot info
- add slot
- confirm student slot
- cancel student slot

## Tech stack

 - Go
 - Chi Router (go-chi)
 - PostgreSQL
 - Redis
 - Gorilla WebSockets
 - Viper
 - JWT (go-jwt v5)

## Authors

 - Greg - Lord of PHP Solutions, Master of Golang Services, Warden of RESTful Gates, Keeper of gRPC Streams, Protector of SQL Queries, Overlord of RabbitMQ Channels
