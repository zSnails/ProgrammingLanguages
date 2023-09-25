% vim:ft=prolog

sumlist([], 0).
sumlist([H|T], S) :-
    sumlist(T, K),
    S is H + K.

% NOTE: Esta es una implementación bonus que se me ocurrió, suelta una
% expresión donde el resultado es la suma de todos los demás.
% sumlist([], 0).
% sumlist([H|T], H + K) :-
%     sumlist(T, K).
