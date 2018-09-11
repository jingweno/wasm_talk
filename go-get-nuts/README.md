# Go Get Nuts Wasm

A Go Webassembly port of https://github.com/jingweno/go-get-nuts. Demo: https://jingweno.github.io/wasm_talk/go-get-nuts.

`js.Callback` doesn't properly bind to JS `this`. Hack to use `eval` and `Function` to evaluate code.
