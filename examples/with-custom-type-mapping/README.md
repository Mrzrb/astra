# With Custom Type Mapping Example
This is a basic example of how to use the Astra API with substitute typings. It utilises some basic features of the API, and some basic types for CRUD operations for a blog system. The results and data are faked, but the important thing to note here is that the types are carried through the entire system, and Astra can be used to generate a full CRUD system with a few lines of code.

We here setup the `database/sql.NullString` type (indicating the response from a database query may be null) to be mapped to a `string` type in the OpenAPI file. This is done by adding the following option to the `astra.New` function:

We also add our own custom `Comment` struct with a custom JSON marshaler to the OpenAPI file which returns a different type (string in this exmaple), by adding the following option to the `astra.New` function:

```go
astra.WithCustomTypeMapping(map[string]astra.TypeFormat{
    "database/sql.NullString": astra.TypeFormat{
        Type:   "string",
        Format: "",
    },
    "withcustomtypemapping/types.Comment": astra.TypeFormat{
        Type:   "string",
        Format: "",
    },
})
```

## Running the example

To run the example, you need to have a working Go installation. You can run the example by running:

```bash
go run .
```

## Important files

The important files in this example are:
* `main.go` - This is the main file that runs the example.
* `openapi.generated.yaml` - This is the OpenAPI file that is generated by Astra.