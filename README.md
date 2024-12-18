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
