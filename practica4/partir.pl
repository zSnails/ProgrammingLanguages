% vim:ft=prolog

partir([], _, [], []).
partir([H|Hs], U, [H|As], Bs) :-
    H < U,
    partir(Hs, U, As, Bs),
    !.
partir([H|Hs], U, As, [H|Bs]) :-
    H >= U,
    partir(Hs, U, As, Bs),
    !.
