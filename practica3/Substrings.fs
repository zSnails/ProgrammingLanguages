module Substrings

let subStrings (str: string) (strings: string list): string list =
    List.filter (fun (ss: string) -> ss.Contains(str)) strings
