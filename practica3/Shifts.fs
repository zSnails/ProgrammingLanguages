module Shifts

let private zeroes _ = 0

let private leftShift (times: int) (datum: _ list): _ list =
    List.init times zeroes
    |> (List.removeManyAt 0 times datum |> List.append)

let private rightShift (times: int) (datum: _ list): _ list =
    List.take (List.length datum - times) datum
    |> (List.init times zeroes |> List.append)

let public shift dir times datum =
    function
    | "der" -> rightShift times datum
    | "izq" -> leftShift times datum
    | _ -> [] // NOTE: handle every other case, just so that the language server stops screaming at me
