% vim:ft=prolog

subconj(_, []).
subconj(Y, [X|Tail]):-
  select(X, Y, Z),
  subconj(Z, Tail),
  !.
