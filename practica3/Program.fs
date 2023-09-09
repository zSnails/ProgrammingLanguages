open Shifts
open Substrings
open Nthelement

[<EntryPoint>]
let main (args: string array): int =
    shift "izq" 10 [1..20] |> printfn "%A"
    subStrings "la" ["la casa"; "el perro"; "pintando la cerca"; "la camisa"] |> printfn "%A"
    nthElement 2 ["hol", "me", "samuel", "adios"] |> printfn "%A"
    0
