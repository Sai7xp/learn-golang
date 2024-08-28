There's no concept of inheritance in Go unlike Java or C++

The closest would be embedding a struct into another struct - syntax sugar then allows you to call methods of the embedded struct on the outer struct directly. But this is not inheritance, and is purely syntax sugar

## 1. Embedding

When you embed a struct in another struct, the inner struct's fields and methods are accessible directly through the outer struct, as if they were part of it. This provides a form of type reuse without the need for explicit inheritance

### When to Use Embedding

[Code Examples](basics/09Composition&Embedding/embedding.go)

**Code Reusability:** When you want to reuse fields and methods from another struct without having to explicitly reference the inner struct.

**Is-a Relationship:** When the outer struct can be considered a specialized form of the embedded struct.

## 2. Composition

Composition in Go refers to including one or more structs as fields of another struct. Unlike embedding, accessing the fields and methods of the included structs requires explicit referencing.

### When to Use Composition

[Code Examples](basics/09Composition&Embedding/composition.go)

**Has-a Relationship:** When the outer struct "has" a certain struct as part of its definition. For instance, a Person has an Address.

**Modularity:** When you want to clearly separate concerns by having distinct types for different functionalities.

**Flexibility:** When you want the freedom to change the inner struct without affecting the outer struct or when you want to include multiple types that might share similar fields or methods without implying a direct relationship.
