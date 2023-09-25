% vim:ft=prolog

aplanar([], []) :- !.
aplanar([L|Ls], FlatL) :-
    !,
    aplanar(L, NewL),
    aplanar(Ls, NewLs),
    append(NewL, NewLs, FlatL).
aplanar(L, [L]).
