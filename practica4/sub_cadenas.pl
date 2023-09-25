% vim:ft=prolog

% NOTE: esto depende de que subconj/2 exista, cargar el archivo `subconj.pl'
% antes de cargar este.

sub_cadenas(_, [], []).
sub_cadenas(Txt, [H|T], [H|Bs]) :-
    subconj(H, Txt),
    sub_cadenas(Txt, T, Bs),
    !.
sub_cadenas(Txt, [_|T], Bs) :-
    sub_cadenas(Txt, T, Bs),
    !.
