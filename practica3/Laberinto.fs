module Laberinto

type Node = (int * int list)
type Graph = Node list
type Route = int list

let rec private neighbors (nodo: int) (grafo: Graph): Route =
    match grafo with
    | [] -> []
    | (head, neigh) :: rest ->
        if head = nodo then neigh
        else neighbors nodo rest

let private isMember (elem) (lista): bool =
    List.exists (fun x -> x = elem) lista

let private extend (ruta: Route) (grafo: Graph): Route list = 
    List.filter
        (fun x -> x <> [])
        (List.map (fun x -> if (isMember x ruta) then [] else x::ruta) (neighbors (List.head ruta) grafo)) 

let private shortest (rutas: Route list): Route =
    List.minBy (fun l -> List.length l) rutas

let private fin = -1
let private ini = 0

let public shortestPath (grafo: Graph): Route =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) :: prof_aux fin ((extend (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extend (List.head ruta) grafo)) grafo
    prof_aux fin [[ini]] grafo |> shortest
