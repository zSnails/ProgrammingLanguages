module Nthelement


let nthElement<'K>(n: int) (datum: 'K list): Option<'K> =
    if n >= List.length datum then None
    else
        let mapper (idx: int) (elem: 'K) =
            match (idx, elem) with
            | (a, b) when a = n -> [b]
            | (a, b) -> []
        ([0..(List.length datum)-1], datum) 
        ||> List.map2 mapper
        |> List.collect id
        |> List.head |> Some

