module Pesos

type Node = (string * string list * int list)
type Graph = Node list

let rec neighbors (nodo: string) =
    function
    | [] -> []
    | (head, neigh, _) :: rest ->
        if head = nodo then neigh
        else neighbors nodo rest

let isMember (elem) (lista): bool =
    List.exists (fun x -> x = elem) lista

let extend ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map (fun x -> if (isMember x ruta) then [] else x::ruta) (neighbors (List.head ruta) grafo)) 

let rec getWeight from ``to`` = function
    | (head, items, weights)::rest when head = from ->
        List.tryFindIndex (fun x -> x = ``to``) items
        |> Option.bind (fun index -> if List.item index weights <> 0 then Some (List.item index weights) else None)
        |> Option.defaultValue 0
    | _::rest -> getWeight from ``to`` rest
    | [] -> 0


let weightClosure (graph: Graph) =
    fun ruta ->
        let rec accWeight acc = function
            | [] | [_] -> acc
            | x::y::rest -> accWeight (acc + getWeight x y graph) (y::rest)
        (ruta, accWeight 0 ruta)

let shortestPath (ini: string) (fin: string) (grafo: Graph): string list =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin ((extend (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extend (List.head ruta) grafo)) grafo
    prof_aux fin [[ini]] grafo |> List.map (weightClosure grafo) |> List.minBy (fun n -> snd n) |> fst
