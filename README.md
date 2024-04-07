# features

Simple implementation of the flags feature.

### Usage

```golang
import "github.com/bavix/features"

const (
    Dishing features.Flag = iota
    Washing
    Cooking
    Driving
)
//...
toggles := features.New(Dishing, Washing, Cooking)

toggles.Has(Cooking) // true
toggles.Has(Driving) // false
```

---
Supported by

[![Supported by JetBrains](https://cdn.rawgit.com/bavix/development-through/46475b4b/jetbrains.svg)](https://www.jetbrains.com/)
