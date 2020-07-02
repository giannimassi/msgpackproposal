# Variable tuple size proposal

## Introduction
I would like to propose a new feature for the code generator that would allow to generate code that implements a decoder/unmarshaller that is loose on array size, meaning the equivalent of [this code ](https://play.golang.org/p/68sN1H3Zh6w) can be written for msgpack.

As far as I can understand the feature I need is not supported yet (see https://github.com/tinylib/msgp/wiki/Type-Mapping-Rules#rule-3-tuple-to-struct ). I'm interested in understanding if it is feasible to add a flag to the generator command to support it or alternatively if there are better solution for this scope.

## Implementation

I have modified the generated code in my application to see what generated code should look like, as follows:

- I have removed the check on the length encoded in the array header
- For each field in the current version of the type I only try to read the value into the struct if the array header size is long enough to include the field, if not I exit early.
- After reading all expected fields, I look at the array header size and if the array contains more fields than expected, I simply skip the remaining fields.

I have put some example code in [this repository](https://github.com/giannimassi/msgpackproposal), the modified implementations are in main.go. `main_test.go` includes the tests documenting the requirement. The `*Variable` tests are the only one that pass since they use the modified version of the methods.

This is what the DecodeMsg method would look like:

```go
type StructV1 struct {
    A uint32
    B uint32
}

func (z *StructV1) DecodeMsg(dc *msgp.Reader) (err error) {
    var zb0001 uint32
    zb0001, err = dc.ReadArrayHeader()
    if err != nil {
        err = msgp.WrapError(err)
        return
    }
    if zb0001--; zb0001 == 0 {
        return
    }
    z.A, err = dc.ReadUint32()
    if err != nil {
        err = msgp.WrapError(err, "A")
        return
    }
    if zb0001--; zb0001 == 0 {
        return
    }
    z.B, err = dc.ReadUint32()
    if err != nil {
        err = msgp.WrapError(err, "B")
        return
    }
    for ; zb0001 > 0; zb0001-- {
        if err = dc.Skip(); err != nil {
            err = msgp.WrapError(err, zb0001)
            return
        }
    }
    return
}
```

The good thing about the implementation I am proposing is that it looks like it should be simple enough to add the code to the generator, since it does not depend on the type it is generating code for, but it simply adds the following after each fields read:

```go
    if zb0001--; zb0001 == 0 {
        return
    }
```

and this at the end:

```go
for ; zb0001 > 0; zb0001-- {
        if err = dc.Skip(); err != nil {
            err = msgp.WrapError(err, zb0001)
            return
        }
    }
```
where `zb0001` is the number of elements in the array, as read from the array header.

Do you think it is feasible to add a flag to generate this code? Something like `-var-tuple-size`?
Is there any other solution to my requirement? Any simpler solution?

[For reference: in my usecase I am interfacing with another program that already implements this feature and it uses this library: https://github.com/msgpack/msgpack-c]