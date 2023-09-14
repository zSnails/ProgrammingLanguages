module Rutacorta

type Node = (string * string list * int list)
type Graph = Node list

// Función para generar vecinos
let rec vecinos (nodo: string) (grafo: Graph): string list =
    match grafo with
    | [] -> []
    | (head, neighbors, _) :: rest ->
        if head = nodo then neighbors
        else vecinos nodo rest

let miembro (elem) (lista): bool =
    List.exists (fun x -> x = elem) lista

// Función para extender una ruta
let extender ruta grafo = 
    List.filter
        (fun x -> x <> [])
        (List.map (fun x -> if (miembro x ruta) then [] else x::ruta) (vecinos (List.head ruta) grafo)) 

// Función principal de búsqueda en profundidad
let prof2 (ini: string) (fin: string) (grafo: Graph) =
    let rec prof_aux fin ruta grafo =
        if List.isEmpty ruta then []
        elif (List.head (List.head ruta)) = fin then
            List.rev (List.head ruta) //:: prof_aux fin ((extender (List.head ruta) grafo) @ (List.tail ruta)) grafo       
        else
            prof_aux fin ((List.tail ruta) @ (extender (List.head ruta) grafo)) grafo
    prof_aux fin [[ini]] grafo
