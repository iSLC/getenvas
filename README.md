# Description
Go lang utility to retrieve an environment variable as a certain value type or fall back to a default value if not set.

# Basic Example
```go
package main

import (
  "fmt"

  "github.com/iSLC/getenvas"
)

func main() {
  os.Setenv("GETENVAS_STRING_VARIABLE", "ABC")
  os.Setenv("GETENVAS_STRING_VARIABLE_EMPTY", "")

  // Strings don't require parsing and therefore they don't need to return an error
  fmt.Println("GETENVAS_STRING_VARIABLE:", getenvas.String("GETENVAS_STRING_VARIABLE", "default")) // "ABC"
  fmt.Println("GETENVAS_STRING_VARIABLE_EMPTY:", getenvas.String("GETENVAS_STRING_VARIABLE_EMPTY", "default")) // ""
  fmt.Println("GETENVAS_STRING_VARIABLE_NONEXISTENT:", getenvas.String("GETENVAS_STRING_VARIABLE_NONEXISTENT", "default")) // "default"

  os.Setenv("GETENVAS_BOOL_VARIABLE", "true")
  os.Setenv("GETENVAS_BOOL_VARIABLE_EMPTY", "")
  os.Setenv("GETENVAS_BOOL_VARIABLE_BAD", "???")

  if value, err := getenvas.Bool("GETENVAS_BOOL_VARIABLE", false); err != nil {
    fmt.Println("GETENVAS_BOOL_VARIABLE:", value, err)
  } else {
    fmt.Println("GETENVAS_BOOL_VARIABLE:", value) // true
  }

  if value, err := getenvas.Bool("GETENVAS_BOOL_VARIABLE_EMPTY", false); err != nil {
    fmt.Println("GETENVAS_BOOL_VARIABLE_EMPTY:", value, err) // false (invalid syntax)
  } else {
    fmt.Println("GETENVAS_BOOL_VARIABLE_EMPTY:", value)
  }

  if value, err := getenvas.Bool("GETENVAS_BOOL_VARIABLE_BAD", false); err != nil {
    fmt.Println("GETENVAS_BOOL_VARIABLE_BAD:", value, err) // false (invalid syntax)
  } else {
    fmt.Println("GETENVAS_BOOL_VARIABLE_BAD:", value)
  }

  if value, err := getenvas.Bool("GETENVAS_BOOL_VARIABLE_NONEXISTENT", false); err != nil {
    fmt.Println("GETENVAS_BOOL_VARIABLE_NONEXISTENT:", value, err)
  } else {
    fmt.Println("GETENVAS_BOOL_VARIABLE_NONEXISTENT:", value) // false
  }

  os.Setenv("GETENVAS_INT_VARIABLE", "667")
  os.Setenv("GETENVAS_INT_VARIABLE_EMPTY", "")
  os.Setenv("GETENVAS_INT_VARIABLE_BAD", "???")

  if value, err := getenvas.Int("GETENVAS_INT_VARIABLE", 236); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE:", value, err)
  } else {
    fmt.Println("GETENVAS_INT_VARIABLE:", value) // 236
  }

  if value, err := getenvas.Int("GETENVAS_INT_VARIABLE_EMPTY", 236); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE_EMPTY:", value, err) // 236 (invalid syntax)
  } else {
    fmt.Println("GETENVAS_INT_VARIABLE_EMPTY:", value)
  }

  if value, err := getenvas.Int("GETENVAS_INT_VARIABLE_BAD", 236); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE_BAD:", value, err) // 236 (invalid syntax)
  } else {
    fmt.Println("GETENVAS_INT_VARIABLE_BAD:", value)
  }

  if value, err := getenvas.Int("GETENVAS_INT_VARIABLE_NONEXISTENT", 236); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE_NONEXISTENT:", value, err)
  } else {
    fmt.Println("GETENVAS_INT_VARIABLE_NONEXISTENT:", value) // 236
  }

  // Repeat for:
  // Uint()
  // Int64()
  // Uint64()
  // Float32()
  // Float64()
  // Duration()
  // ...
}
```

# Into Variable Example
This exists mainly to mirror the functionality from the `flag` package.
```go
package main

import (
  "fmt"

  "github.com/iSLC/getenvas"
)

func main() {
  var str_var string
  // Again, strings don't need to return any errors
  getenvas.StringVar(&str_var, "GETENVAS_STRING_VARIABLE", "default")
  fmt.Println("GETENVAS_STRING_VARIABLE:", str_var)
  var int_var int

  if err := getenvas.IntVar(&int_var, "GETENVAS_INT_VARIABLE", 556); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE:", int_var, err) // int_var is filled with default value
  }
  fmt.Println("GETENVAS_INT_VARIABLE:", int_var)

  // Repeat for:
  // BoolVar()
  // UintVar()
  // Int64Var()
  // Uint64Var()
  // Float32Var()
  // Float64Var()
  // DurationVar()
  // ...
}
```

# Prefixed Names Example
A helper utility exists if you wish to "namespace" your variables.
```go
package main

import (
  "fmt"

  "github.com/iSLC/getenvas"
)

func main() {
  pfx_env := getenvas.Prefixed("GETENVAS_")

  fmt.Println("Prefix is: ", pfx_env.Get()) // pfx_env.Set("...")
  fmt.Println("MY_VAR will become: ", pfx_env.Do("MY_VAR")) // GETENVAS_MY_VAR

  // Both methods are available
  fmt.Println("GETENVAS_STRING_VARIABLE:", pfx_env.String("STRING_VARIABLE", "default"))
  var str_var string
  pfx_env.StringVar(&str_var, "STRING_VARIABLE", "default")
  fmt.Println("GETENVAS_STRING_VARIABLE:", str_var)

  // Or methods that may encounter parsing errors
  if value, err := pfx_env.Int("INT_VARIABLE", 476); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE:", value, err)
  } else {
    fmt.Println("GETENVAS_INT_VARIABLE:", value)
  }

  var int_var int
  if err := pfx_env.IntVar(&int_var, "INT_VARIABLE", 476); err != nil {
    fmt.Println("GETENVAS_INT_VARIABLE:", int_var, err) // int_var is filled with default value
  }
  fmt.Println("GETENVAS_INT_VARIABLE:", int_var)

  // All functions are aliased with the same name
}
```
