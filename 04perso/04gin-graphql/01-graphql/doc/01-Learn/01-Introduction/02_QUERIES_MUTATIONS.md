# Queries and Mutations

[*Source*](http://graphql.org/learn/queries/)

*On this page, you'll learn in detail about how to query a GraphQL server.*

# [Fields](http://graphql.org/learn/queries/#fields)
At its simplest, GraphQL is about asking for specific fields on objects. Let's start by looking at a very simple query and the result we get when we run it:
```graphql
{
  hero {
    name
  }
}
```
```json
{
  "data": {
    "hero": {
      "name": "R2-D2"
    }
  }
}
```

You can see immediately that **the query has exactly the same shape as the result**. This is essential to GraphQL, because you always get back what you expect, and the server knows exactly what fields the client is asking for.

In the previous example, we just asked for the name of our hero which returned a ``String``, but fields can also refer to Objects. In that case, you can make a sub-selection of fields for that object. GraphQL queries can traverse related objects and their fields, letting clients fetch lots of related data in one request, instead of making several roundtrips as one would need in a classic REST architecture.
```graphql
{
  hero {
    name
    # Queries can have comments!
    friends {
      name
    }
  }
}
```
```json
{
  "data": {
    "hero": {
      "name": "R2-D2",
      "friends": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Han Solo"
        },
        {
          "name": "Leia Organa"
        }
      ]
    }
  }
}
```
Note that in this example, the ``friends`` field returns an array of items. **GraphQL queries look the same for both single items or lists of items**, however we know which one to expect based on what is indicated in the schema.

## [Arguments](http://graphql.org/learn/queries/#arguments)
If the only thing we could do was traverse objects and their fields, GraphQL would already be a very useful language for data fetching. But when you add the ability to pass arguments to fields, things get much more interesting.
```graphql
{
  human(id: "1000") {
    name
    height
  }
}
```
```json
{
  "data": {
    "human": {
      "name": "Luke Skywalker",
      "height": 1.72
    }
  }
}
```

In a system like REST, you can only pass a single set of arguments - the query parameters and URL segments in your request. But in GraphQL, every field and nested object can get its own set of arguments, making GraphQL a complete replacement for making multiple API fetches. You can even pass arguments into scalar fields, to implement data transformations once on the server, instead of on every client separately.

```grahpql
{
  human(id: "1000") {
    name
    height(unit: FOOT)
  }
}
```
```json
{
  "data": {
    "human": {
      "name": "Luke Skywalker",
      "height": 5.6430448
    }
  }
}
```

Arguments can be of many different types. In the above example, we have used an Enumeration type, which represents one of a finite set of options (in this case, units of length, either METER or FOOT). GraphQL comes with a default set of types, but a GraphQL server can also declare its own custom types, as long as they can be serialized into your transport format.

[Read more about the GraphQL type system here.](http://graphql.org/learn/schema)

## [Aliases](http://graphql.org/learn/queries/#aliases)
If you have a sharp eye, you may have noticed that, since **the result object fields match the name of the field in the query but don't include arguments, you can't directly query for the same field with different arguments**. That's why you need aliases - they let you rename the result of a field to anything you want.

```graphql
{
  empireHero: hero(episode: EMPIRE) {
    name
  }
  jediHero: hero(episode: JEDI) {
    name
  }
}
```
```json
{
  "data": {
    "empireHero": {
      "name": "Luke Skywalker"
    },
    "jediHero": {
      "name": "R2-D2"
    }
  }
}
```

In the above example, **the two ``hero`` fields would have conflicted**, but since we can alias them to different names, we can get both results in one request.

## [Fragments](http://graphql.org/learn/queries/#fragments)

Let's say we had a relatively complicated page in our app, which let us look at two heroes side by side, along with their friends. You can imagine that such a query could quickly get complicated, because **we would need to repeat the fields at least once - one for each side of the comparison**.

That's why GraphQL includes **reusable units called fragments**. Fragments let you construct sets of fields, and then include them in queries where you need to. Here's an example of how you could solve the above situation using fragments:


```graphql
{
  leftComparison: hero(episode: EMPIRE) {
    ...comparisonFields
  }
  rightComparison: hero(episode: JEDI) {
    ...comparisonFields
  }
}

fragment comparisonFields on Character {
  name
  appearsIn
  friends {
    name
  }
}
```
```json
{
  "data": {
    "leftComparison": {
      "name": "Luke Skywalker",
      "appearsIn": [
        "NEWHOPE",
        "EMPIRE",
        "JEDI"
      ],
      "friends": [
        {
          "name": "Han Solo"
        },
        {
          "name": "Leia Organa"
        },
        {
          "name": "C-3PO"
        },
        {
          "name": "R2-D2"
        }
      ]
    },
    "rightComparison": {
      "name": "R2-D2",
      "appearsIn": [
        "NEWHOPE",
        "EMPIRE",
        "JEDI"
      ],
      "friends": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Han Solo"
        },
        {
          "name": "Leia Organa"
        }
      ]
    }
  }
}
```

You can see how the above query would be pretty repetitive if the fields were repeated. The concept of fragments is frequently used to split complicated application data requirements into smaller chunks, especially when you need to combine lots of UI components with different fragments into one initial data fetch.

## [Operation name](http://graphql.org/learn/queries/#operation-name)
Up until now, we have been using a shorthand syntax where we omit both the query keyword and the query name, but in production apps it's useful to use these to make our code less ambiguous. You'll need these optional parts to a GraphQL operation if you want to execute something other than a query or pass dynamic variables.

Here’s an example that includes the keyword ``query`` as operation type and ``HeroNameAndFriends`` as operation name :

```graphql
query HeroNameAndFriends {
  hero {
    name
    friends {
      name
    }
  }
}
```

```json
{
  "data": {
    "hero": {
      "name": "R2-D2",
      "friends": [
        {
          "name": "Luke Skywalker"
        },
        {
          "name": "Han Solo"
        },
        {
          "name": "Leia Organa"
        }
      ]
    }
  }
}
```



