
---

# Payload Examples for Create and Update Queries

## Create Movie Payload

To create a new movie entry, use the following payload structure:

```json
{
  "isbn": "1234567890",
  "title": "Example Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

Replace the `isbn`, `title`, `firstname`, and `lastname` fields with the respective information for the movie and director you want to create.

---

## Update Movie Payload

To update an existing movie entry, use the following payload structure:

```json
{
  "isbn": "9876543210",
  "title": "Updated Example Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```

Replace the `isbn`, `title`, `firstname`, and `lastname` fields with the updated information for the movie and director.

---
