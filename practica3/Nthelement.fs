module Nthelement


let nthElement<'K>(n: int) (datum: 'K list): 'K =
    let mapper (idx: int) (elem: 'K) =
        match (idx, elem) with
        | (a, b) when a = 3 -> [a]
        | (a, b) -> []
    ([0..(List.length datum)-1], datum) 
    ||> List.map2 mapper
    |> List.collect id
    // TODO: see what the fuck is causing this
    |> List.head
