SOLID Principles - Set of 5 design principles for creating maintainable and scalable software. Adopting these practices can also contribute to avoiding code smells. Introduced by Robert C. Martin aka Uncle Bob.

- Single Responsibility Principle
- Open-Closed Principle
- Liskov Substitution Principle
- Interface Segregation Principle
- Dependency Inversion Principle

## Single Responsibility Principle

A class should have only one reason to change. Every Class should have only one responsibility.

## Open Closed Principle

Software entities should be open for extension, but closed for modification.

## Liskov Substitution Principle

Object of a superclass shall be replaceble with objects of its subclasses
without breaking the application.

## Interface Segregation Principle

Clients should not be forced to depend upon interfaces that they do not use.

Segregation means keeping things separated, and the Interface Segregation Principle is about separating the interfaces. Interfaces should be designed to be as small and specific as possible.

📌 Core idea: Split fat interfaces into smaller, more specific ones.

## Dependency Inversion Principle

High-level modules should not depent on low-level modules; Both should depend on abstractions.

HDFC Microservices Example: 
- We have a service where we have the main business logic and we have another layer called DAO or database layer where we have all db operations. If service layer directly depends on database layer it becomes tightly coupled. Incase if we want to change the db from aerospike to mysql we may have to make changes in service layer as well. And the Unit Testing becomes difficult
- So we should not directly depend on external dependencies, instead we should depend on abstractions.