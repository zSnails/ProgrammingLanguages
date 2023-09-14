module Nthelement


let tryNthElement<'K>(n: int) (datum: 'K list): Option<'K> =
    if n >= List.length datum then None
    else
        ([0..(List.length datum)-1], datum) 
        ||> List.map2 (fun a b -> if a = n then [b] else [])
        |> List.collect id
        |> List.head |> Some

