# API Documentation

This document provides detailed information about the GraphQL API endpoints, types, and operations available in the Ticketing System.

## Authentication

All API requests (except login) require authentication using JWT tokens.

### Login

```graphql
mutation Login($input: LoginInput!) {
  login(input: $input) {
    token
    user {
      id
      email
      name
      role
    }
  }
}
```

Input:
```json
{
  "input": {
    "email": "user@example.com",
    "password": "password123"
  }
}
```

## Queries

### Get Current User

```graphql
query Me {
  me {
    id
    email
    name
    role
    organization {
      id
      name
    }
  }
}
```

### List Tickets

```graphql
query Tickets($filter: TicketFilter) {
  tickets(filter: $filter) {
    id
    title
    description
    status
    priority
    createdAt
    updatedAt
    assignedTo {
      id
      name
    }
    createdBy {
      id
      name
    }
    organization {
      id
      name
    }
  }
}
```

Filter options:
```json
{
  "filter": {
    "status": "OPEN",
    "priority": "HIGH",
    "assignedTo": "user-id",
    "organization": "org-id",
    "search": "search term"
  }
}
```

### Get Single Ticket

```graphql
query Ticket($id: ID!) {
  ticket(id: $id) {
    id
    title
    description
    status
    priority
    createdAt
    updatedAt
    assignedTo {
      id
      name
    }
    createdBy {
      id
      name
    }
    organization {
      id
      name
    }
    comments {
      id
      content
      createdAt
      user {
        id
        name
      }
    }
  }
}
```

### List Users

```graphql
query Users($filter: UserFilter) {
  users(filter: $filter) {
    id
    email
    name
    role
    organization {
      id
      name
    }
    createdAt
    updatedAt
  }
}
```

Filter options:
```json
{
  "filter": {
    "role": "ADMIN",
    "organization": "org-id",
    "search": "search term"
  }
}
```

### List Organizations

```graphql
query Organizations($filter: OrganizationFilter) {
  organizations(filter: $filter) {
    id
    name
    description
    createdAt
    updatedAt
    users {
      id
      name
      role
    }
  }
}
```

Filter options:
```json
{
  "filter": {
    "search": "search term"
  }
}
```

## Mutations

### Create Ticket

```graphql
mutation CreateTicket($input: CreateTicketInput!) {
  createTicket(input: $input) {
    id
    title
    description
    status
    priority
    createdAt
    assignedTo {
      id
      name
    }
  }
}
```

Input:
```json
{
  "input": {
    "title": "New Ticket",
    "description": "Ticket description",
    "priority": "HIGH",
    "assignedTo": "user-id"
  }
}
```

### Update Ticket

```graphql
mutation UpdateTicket($id: ID!, $input: UpdateTicketInput!) {
  updateTicket(id: $id, input: $input) {
    id
    title
    description
    status
    priority
    updatedAt
    assignedTo {
      id
      name
    }
  }
}
```

Input:
```json
{
  "input": {
    "title": "Updated Title",
    "description": "Updated description",
    "status": "IN_PROGRESS",
    "priority": "MEDIUM",
    "assignedTo": "user-id"
  }
}
```

### Delete Ticket

```graphql
mutation DeleteTicket($id: ID!) {
  deleteTicket(id: $id)
}
```

### Create User

```graphql
mutation CreateUser($input: CreateUserInput!) {
  createUser(input: $input) {
    id
    email
    name
    role
    organization {
      id
      name
    }
  }
}
```

Input:
```json
{
  "input": {
    "email": "newuser@example.com",
    "password": "password123",
    "name": "New User",
    "role": "USER",
    "organization": "org-id"
  }
}
```

### Update User

```graphql
mutation UpdateUser($id: ID!, $input: UpdateUserInput!) {
  updateUser(id: $id, input: $input) {
    id
    email
    name
    role
    organization {
      id
      name
    }
  }
}
```

Input:
```json
{
  "input": {
    "name": "Updated Name",
    "role": "ADMIN",
    "organization": "org-id"
  }
}
```

### Delete User

```graphql
mutation DeleteUser($id: ID!) {
  deleteUser(id: $id)
}
```

### Create Organization

```graphql
mutation CreateOrganization($input: CreateOrganizationInput!) {
  createOrganization(input: $input) {
    id
    name
    description
  }
}
```

Input:
```json
{
  "input": {
    "name": "New Organization",
    "description": "Organization description"
  }
}
```

### Update Organization

```graphql
mutation UpdateOrganization($id: ID!, $input: UpdateOrganizationInput!) {
  updateOrganization(id: $id, input: $input) {
    id
    name
    description
  }
}
```

Input:
```json
{
  "input": {
    "name": "Updated Organization",
    "description": "Updated description"
  }
}
```

### Delete Organization

```graphql
mutation DeleteOrganization($id: ID!) {
  deleteOrganization(id: $id)
}
```

## Types

### User

```graphql
type User {
  id: ID!
  email: String!
  name: String!
  role: UserRole!
  organization: Organization
  createdAt: Time!
  updatedAt: Time!
}

enum UserRole {
  ADMIN
  USER
}
```

### Ticket

```graphql
type Ticket {
  id: ID!
  title: String!
  description: String!
  status: TicketStatus!
  priority: TicketPriority!
  assignedTo: User
  createdBy: User!
  organization: Organization!
  comments: [Comment!]
  createdAt: Time!
  updatedAt: Time!
}

enum TicketStatus {
  OPEN
  IN_PROGRESS
  RESOLVED
  CLOSED
}

enum TicketPriority {
  LOW
  MEDIUM
  HIGH
  URGENT
}
```

### Organization

```graphql
type Organization {
  id: ID!
  name: String!
  description: String
  users: [User!]
  createdAt: Time!
  updatedAt: Time!
}
```

### Comment

```graphql
type Comment {
  id: ID!
  content: String!
  ticket: Ticket!
  user: User!
  createdAt: Time!
  updatedAt: Time!
}
```

## Error Handling

The API uses standard GraphQL error handling. All errors will be returned in the following format:

```json
{
  "errors": [
    {
      "message": "Error message",
      "path": ["fieldName"],
      "extensions": {
        "code": "ERROR_CODE"
      }
    }
  ]
}
```

Common error codes:
- `UNAUTHENTICATED`: User is not authenticated
- `UNAUTHORIZED`: User does not have permission
- `NOT_FOUND`: Resource not found
- `VALIDATION_ERROR`: Input validation failed
- `INTERNAL_ERROR`: Server error 