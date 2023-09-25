% vim:ft=prolog

sumlist([], 0).
sumlist([H|T], S) :-
    sumlist(T, K),
    S is H + K.
