open Shifts
open Substrings
open Nthelement
open Rutacorta
open Laberinto
open Pesos

let private grafo: Rutacorta.Graph = [
    ("i", ["a"; "b"], [6; 6]);
    ("a", ["i"; "c"; "d"], [5; 7; 9]);
    ("b", ["i"; "c"; "d"], [6; 10; 11]);
    ("c", ["a"; "b"; "x"], [7; 10; 21]);
    ("d", ["a"; "b"; "f"], [9; 11; 32]);
    ("f", ["d"], [32]);
    ("x", ["c"], [21])
]

let private labirynth: Laberinto.Graph = [
    (0, [2]);
    (2, [8; 3]);
    (8, [2; 9; 14]); // XXX: rompí la pared entre el 8 y el 14
    (9, [8; 3]);
    (3, [2; 9; 4]);
    (4, [3; 10]);
    (10, [4; 16]);
    (16, [10; 22]);
    (22, [16; 21; 28]); // XXX: rompí la pared entre 22 y 28
    (21, [22; 15]);
    (15, [21; 14]);
    (14, [15; 20; 13; 8]); // XXX: rompí la pared entre el 8 y 14
    (13, [14; 7; 19]); // XXX: rompí la pared entre el 13 y 19
    (7, [13; 1]);
    (20, [14; 26]);
    (26, [20; 27]);
    (27, [26; 28]);
    (28, [29; 34; 27; 22]); // XXX: rompí la pared entre el 22 y 28
    (29, [23; 28]);
    (23, [29; 17]);
    (17, [23; 11]);
    (11, [17; 5]);
    (5, [6; 11]);
    (34, [28; 35; 33]);
    (35, [36; 34]);
    (36, [30; 35]);
    (30, [24; 36]);
    (24, [18; 30]);
    (18, [12; 24]);
    (12, [18]);
    (33, [34; 32]);
    (32, [33; -1; 31]);
    (31, [32; 25]);
    (25, [31; 19]);
    (19, [25; 13]);
    (-1, [32]);
]

[<EntryPoint>]
let main (args: string array): int =
    shift "izq" 10 [1..20] |> printfn "%A"
    subStrings "la" ["la casa"; "el perro"; "pintando la cerca"; "la camisa"] |> printfn "%A"

    tryNthElement 4 ["hola"; "me"; "llamo"; "aaron"]
    |> function 
    | Some(result) -> printfn "%A" result 
    | None -> printfn "Element not found"

    prof2 "i" "x" grafo |> printfn "%A"

    Laberinto.shortestPath labirynth |> printfn "%A"
    Pesos.shortestPath "i" "x" grafo |> printfn "%A"

    0
